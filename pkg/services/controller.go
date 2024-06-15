package services

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/templates"

	"github.com/labstack/echo/v4"
)

// Controller TODO
type Controller struct {
	config   *config.Config
	cache    *CacheClient
	renderer *TemplateRenderer
}

// TODO where does the cache stuff belong?

// CachedPageGroup stores the cache group for cached pages
const CachedPageGroup = "page"

// CachedPage is what is used to store a rendered Page in the cache
type CachedPage struct {
	// URL stores the URL of the requested page
	URL string

	// HTML stores the complete HTML of the rendered Page
	HTML []byte

	// StatusCode stores the HTTP status code
	StatusCode int

	// Headers stores the HTTP headers
	Headers map[string]string
}

// NewController creates a new Controller
func NewController(cfg *config.Config, cache *CacheClient, renderer *TemplateRenderer) *Controller {
	return &Controller{
		config:   cfg,
		cache:    cache,
		renderer: renderer,
	}
}

// RenderPage renders a Page as an HTTP response
func (c *Controller) RenderPage(ctx echo.Context, page page.Page) error {
	var buf *bytes.Buffer
	var err error
	templateGroup := "page"

	// Page name is required
	if page.Name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "page render failed due to missing name")
	}

	// Use the app name in configuration if a value was not set
	if page.AppName == "" {
		page.AppName = c.config.App.Name
	}

	// Check if this is an HTMX non-boosted request which indicates that only partial
	// content should be rendered
	if page.HTMX.Request.Enabled && !page.HTMX.Request.Boosted {
		// Switch the layout which will only render the page content
		page.Layout = templates.LayoutHTMX

		// Alter the template group so this is cached separately
		templateGroup = "page:htmx"
	}

	// Parse and execute the templates for the Page
	// As mentioned in the documentation for the Page struct, the templates used for the page will be:
	// 1. The layout/base template specified in Page.Layout
	// 2. The content template specified in Page.Name
	// 3. All templates within the components directory
	// Also included is the function map provided by the funcmap package
	buf, err = c.renderer.
		Parse().
		Group(templateGroup).
		Key(string(page.Name)).
		Base(string(page.Layout)).
		Files(
			fmt.Sprintf("layouts/%s", page.Layout),
			fmt.Sprintf("pages/%s", page.Name),
		).
		Directories("components").
		Execute(page)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("failed to parse and execute templates: %s", err),
		)
	}

	// Set the status code
	ctx.Response().Status = page.StatusCode

	// Set any headers
	for k, v := range page.Headers {
		ctx.Response().Header().Set(k, v)
	}

	// Apply the HTMX response, if one
	if page.HTMX.Response != nil {
		page.HTMX.Response.Apply(ctx)
	}

	// Cache this page, if caching was enabled
	c.cachePage(ctx, page, buf)

	return ctx.HTMLBlob(ctx.Response().Status, buf.Bytes())
}

// cachePage caches the HTML for a given Page if the Page has caching enabled
func (c *Controller) cachePage(ctx echo.Context, page page.Page, html *bytes.Buffer) {
	if !page.Cache.Enabled || page.IsAuth {
		return
	}

	// If no expiration time was provided, default to the configuration value
	if page.Cache.Expiration == 0 {
		page.Cache.Expiration = c.config.Cache.Expiration.Page
	}

	// Extract the headers
	headers := make(map[string]string)
	for k, v := range ctx.Response().Header() {
		headers[k] = v[0]
	}

	// The request URL is used as the cache key so the middleware can serve the
	// cached page on matching requests
	key := ctx.Request().URL.String()
	cp := CachedPage{
		URL:        key,
		HTML:       html.Bytes(),
		Headers:    headers,
		StatusCode: ctx.Response().Status,
	}

	err := c.cache.
		Set().
		Group(CachedPageGroup).
		Key(key).
		Tags(page.Cache.Tags...).
		Expiration(page.Cache.Expiration).
		Data(cp).
		Save(ctx.Request().Context())

	switch {
	case err == nil:
		log.Ctx(ctx).Debug("cached page")
	case !context.IsCanceledError(err):
		log.Ctx(ctx).Error("failed to cache page",
			"error", err,
		)
	}
}
