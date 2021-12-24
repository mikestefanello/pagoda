package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"

	"goweb/htmx"
	"goweb/middleware"
	"goweb/msg"
	"goweb/services"

	"github.com/go-playground/validator/v10"

	"github.com/eko/gocache/v2/marshaler"

	"github.com/eko/gocache/v2/store"

	"github.com/labstack/echo/v4"
)

// Controller provides base functionality and dependencies to routes.
// The proposed pattern is to embed a Controller in each individual route struct and to use
// the router to inject the container so your routes have access to the services within the container
type Controller struct {
	// Container stores a services container which contains dependencies
	Container *services.Container
}

// NewController creates a new Controller
func NewController(c *services.Container) Controller {
	return Controller{
		Container: c,
	}
}

// RenderPage renders a Page as an HTTP response
func (c *Controller) RenderPage(ctx echo.Context, page Page) error {
	var buf *bytes.Buffer
	var err error

	// Page name is required
	if page.Name == "" {
		ctx.Logger().Error("page render failed due to missing name")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Use the app name in configuration if a value was not set
	if page.AppName == "" {
		page.AppName = c.Container.Config.App.Name
	}

	// Check if this is an HTMX request
	if page.HTMX.Request.Enabled && !page.HTMX.Request.Boosted {
		// Disable caching
		page.Cache.Enabled = false

		// Parse and execute
		buf, err = c.Container.TemplateRenderer.ParseAndExecute(
			"page:htmx",
			page.Name,
			"htmx",
			[]string{
				"htmx",
				fmt.Sprintf("pages/%s", page.Name),
			},
			[]string{"components"},
			page,
		)
	} else {
		// Parse and execute the templates for the Page
		// As mentioned in the documentation for the Page struct, the templates used for the page will be:
		// 1. The layout/base template specified in Page.Layout
		// 2. The content template specified in Page.Name
		// 3. All templates within the components directory
		// Also included is the function map provided by the funcmap package
		buf, err = c.Container.TemplateRenderer.ParseAndExecute(
			"page",
			page.Name,
			page.Layout,
			[]string{
				fmt.Sprintf("layouts/%s", page.Layout),
				fmt.Sprintf("pages/%s", page.Name),
			},
			[]string{"components"},
			page,
		)
	}

	if err != nil {
		ctx.Logger().Errorf("failed to parse and execute templates: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
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
func (c *Controller) cachePage(ctx echo.Context, page Page, html *bytes.Buffer) {
	if !page.Cache.Enabled {
		return
	}

	// If no expiration time was provided, default to the configuration value
	if page.Cache.Expiration == 0 {
		page.Cache.Expiration = c.Container.Config.Cache.Expiration.Page
	}

	// Extract the headers
	headers := make(map[string]string)
	for k, v := range ctx.Response().Header() {
		headers[k] = v[0]
	}

	// The request URL is used as the cache key so the middleware can serve the
	// cached page on matching requests
	key := ctx.Request().URL.String()
	opts := &store.Options{
		Expiration: page.Cache.Expiration,
		Tags:       page.Cache.Tags,
	}
	cp := middleware.CachedPage{
		URL:        key,
		HTML:       html.Bytes(),
		Headers:    headers,
		StatusCode: ctx.Response().Status,
	}

	err := marshaler.New(c.Container.Cache).Set(ctx.Request().Context(), key, cp, opts)
	if err != nil {
		ctx.Logger().Errorf("failed to cache page: %v", err)
		return
	}

	ctx.Logger().Infof("cached page")
}

// Redirect redirects to a given route name with optional route parameters
func (c *Controller) Redirect(ctx echo.Context, route string, routeParams ...interface{}) error {
	url := ctx.Echo().Reverse(route, routeParams)
	h := htmx.Response{}
	h.Redirect = url
	h.Apply(ctx)
	return ctx.Redirect(http.StatusFound, url)
}

func (c *Controller) Fail(ctx echo.Context, err error, log string) error {
	ctx.Logger().Errorf("%s: %v", log, err)
	return echo.NewHTTPError(500)
}

// SetValidationErrorMessages sets error flash messages for validation failures of a given struct
// and attempts to provide more user-friendly wording.
// The error should result from the validator module and the data should be the struct that failed
// validation.
// This method supports including a struct tag of "labeL" on each field which will be the name
// of the field used in the error messages, for example:
//  - FirstName string `form:"first-name" validate:"required" label:"First name"`
// Only a few validator tags are supported below. Expand them as needed.
func (c *Controller) SetValidationErrorMessages(ctx echo.Context, err error, data interface{}) {
	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		return
	}

	for _, ve := range ves {
		var message string

		// Default the field label to the name of the struct field
		label := ve.StructField()

		// Attempt to get a label from the field's struct tag
		if field, ok := reflect.TypeOf(data).FieldByName(ve.Field()); ok {
			if labelTag := field.Tag.Get("label"); labelTag != "" {
				label = labelTag
			}
		}

		// Provide better error messages depending on the failed validation tag
		// This should be expanded as you use additional tags in your validation
		switch ve.Tag() {
		case "required":
			message = "%s is required."
		case "email":
			message = "%s must be a valid email address."
		case "eqfield":
			message = "%s must match."
		default:
			message = "%s is not a valid value."
		}

		msg.Danger(ctx, fmt.Sprintf(message, "<strong>"+label+"</strong>"))
	}
}
