package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/templates"
)

// cachedPageGroup stores the cache group for cached pages
const cachedPageGroup = "page"

type (
	// TemplateRenderer provides a flexible and easy to use method of rendering simple templates or complex sets of
	// templates while also providing caching and/or hot-reloading depending on your current environment
	TemplateRenderer struct {
		// templateCache stores a cache of parsed page templates
		templateCache sync.Map

		// funcMap stores the template function map
		funcMap template.FuncMap

		// config stores application configuration
		config *config.Config

		cache *CacheClient
	}

	// TemplateParsed is a wrapper around parsed templates which are stored in the TemplateRenderer cache
	TemplateParsed struct {
		// Template is the parsed template
		Template *template.Template

		// build stores the build data used to parse the template
		build *templateBuild
	}

	// templateBuild stores the build data used to parse a template
	templateBuild struct {
		group       string
		key         string
		base        string
		files       []string
		directories []string
	}

	// templateBuilder handles chaining a template parse operation
	templateBuilder struct {
		build    *templateBuild
		renderer *TemplateRenderer
	}

	// CachedPage is what is used to store a rendered Page in the cache
	CachedPage struct {
		// URL stores the URL of the requested page
		URL string

		// HTML stores the complete HTML of the rendered Page
		HTML []byte

		// StatusCode stores the HTTP status code
		StatusCode int

		// Headers stores the HTTP headers
		Headers map[string]string
	}
)

// NewTemplateRenderer creates a new TemplateRenderer
func NewTemplateRenderer(cfg *config.Config, cache *CacheClient, fm template.FuncMap) *TemplateRenderer {
	return &TemplateRenderer{
		templateCache: sync.Map{},
		funcMap:       fm,
		config:        cfg,
		cache:         cache,
	}
}

// Parse creates a template build operation
func (t *TemplateRenderer) Parse() *templateBuilder {
	return &templateBuilder{
		renderer: t,
		build:    &templateBuild{},
	}
}

// RenderPage renders a Page as an HTTP response
func (t *TemplateRenderer) RenderPage(ctx echo.Context, page page.Page) error {
	var buf *bytes.Buffer
	var err error
	templateGroup := "page"

	// Page name is required
	if page.Name == "" {
		return echo.NewHTTPError(http.StatusInternalServerError, "page render failed due to missing name")
	}

	// Use the app name in configuration if a value was not set
	if page.AppName == "" {
		page.AppName = t.config.App.Name
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
	buf, err = t.
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
	t.cachePage(ctx, page, buf)

	return ctx.HTMLBlob(ctx.Response().Status, buf.Bytes())
}

// cachePage caches the HTML for a given Page if the Page has caching enabled
func (t *TemplateRenderer) cachePage(ctx echo.Context, page page.Page, html *bytes.Buffer) {
	if !page.Cache.Enabled || page.IsAuth {
		return
	}

	// If no expiration time was provided, default to the configuration value
	if page.Cache.Expiration == 0 {
		page.Cache.Expiration = t.config.Cache.Expiration.Page
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

	err := t.cache.
		Set().
		Group(cachedPageGroup).
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

// GetCachedPage attempts to fetch a cached page for a given URL
func (t *TemplateRenderer) GetCachedPage(ctx echo.Context, url string) (*CachedPage, error) {
	p, err := t.cache.
		Get().
		Group(cachedPageGroup).
		Key(url).
		Type(new(CachedPage)).
		Fetch(ctx.Request().Context())

	if err != nil {
		return nil, err
	}

	return p.(*CachedPage), nil
}

// getCacheKey gets a cache key for a given group and ID
func (t *TemplateRenderer) getCacheKey(group, key string) string {
	if group != "" {
		return fmt.Sprintf("%s:%s", group, key)
	}
	return key
}

// parse parses a set of templates and caches them for quick execution
// If the application environment is set to local, the cache will be bypassed and templates will be
// parsed upon each request so hot-reloading is possible without restarts.
// Also included will be the function map provided by the funcmap package.
func (t *TemplateRenderer) parse(build *templateBuild) (*TemplateParsed, error) {
	var tp *TemplateParsed
	var err error

	switch {
	case build.key == "":
		return nil, errors.New("cannot parse template without key")
	case len(build.files) == 0 && len(build.directories) == 0:
		return nil, errors.New("cannot parse template without files or directories")
	case build.base == "":
		return nil, errors.New("cannot parse template without base")
	}

	// Generate the cache key
	cacheKey := t.getCacheKey(build.group, build.key)

	// Check if the template has not yet been parsed or if the app environment is local, so that
	// templates reflect changes without having the restart the server
	if tp, err = t.Load(build.group, build.key); err != nil || t.config.App.Environment == config.EnvLocal {
		// Initialize the parsed template with the function map
		parsed := template.New(build.base + config.TemplateExt).
			Funcs(t.funcMap)

		// Format the requested files
		for k, v := range build.files {
			build.files[k] = fmt.Sprintf("%s%s", v, config.TemplateExt)
		}

		// Include all files within the requested directories
		for k, v := range build.directories {
			build.directories[k] = fmt.Sprintf("%s/*%s", v, config.TemplateExt)
		}

		// Get the templates
		var tpl fs.FS
		if t.config.App.Environment == config.EnvLocal {
			tpl = templates.GetOS()
		} else {
			tpl = templates.Get()
		}

		// Parse the templates
		parsed, err = parsed.ParseFS(tpl, append(build.files, build.directories...)...)
		if err != nil {
			return nil, err
		}

		// Store the template so this process only happens once
		tp = &TemplateParsed{
			Template: parsed,
			build:    build,
		}
		t.templateCache.Store(cacheKey, tp)
	}

	return tp, nil
}

// Load loads a template from the cache
func (t *TemplateRenderer) Load(group, key string) (*TemplateParsed, error) {
	load, ok := t.templateCache.Load(t.getCacheKey(group, key))
	if !ok {
		return nil, errors.New("uncached page template requested")
	}

	tmpl, ok := load.(*TemplateParsed)
	if !ok {
		return nil, errors.New("unable to cast cached template")
	}

	return tmpl, nil
}

// Execute executes a template with the given data and provides the output
func (t *TemplateParsed) Execute(data any) (*bytes.Buffer, error) {
	if t.Template == nil {
		return nil, errors.New("cannot execute template: template not initialized")
	}

	buf := new(bytes.Buffer)
	err := t.Template.ExecuteTemplate(buf, t.build.base+config.TemplateExt, data)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// Group sets the cache group for the template being built
func (t *templateBuilder) Group(group string) *templateBuilder {
	t.build.group = group
	return t
}

// Key sets the cache key for the template being built
func (t *templateBuilder) Key(key string) *templateBuilder {
	t.build.key = key
	return t
}

// Base sets the name of the base template to be used during template parsing and execution.
// This should be only the file name without a directory or extension.
func (t *templateBuilder) Base(base string) *templateBuilder {
	t.build.base = base
	return t
}

// Files sets a list of template files to include in the parse.
// This should not include the file extension and the paths should be relative to the templates directory.
func (t *templateBuilder) Files(files ...string) *templateBuilder {
	t.build.files = files
	return t
}

// Directories sets a list of directories that all template files within will be parsed.
// The paths should be relative to the templates directory.
func (t *templateBuilder) Directories(directories ...string) *templateBuilder {
	t.build.directories = directories
	return t
}

// Store parsed the templates and stores them in the cache
func (t *templateBuilder) Store() (*TemplateParsed, error) {
	return t.renderer.parse(t.build)
}

// Execute executes the template with the given data.
// If the template has not already been cached, this will parse and cache the template
func (t *templateBuilder) Execute(data any) (*bytes.Buffer, error) {
	tp, err := t.Store()
	if err != nil {
		return nil, err
	}

	return tp.Execute(data)
}
