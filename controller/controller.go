package controller

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"sync"

	"goweb/config"
	"goweb/funcmap"
	"goweb/middleware"
	"goweb/msg"
	"goweb/services"

	"github.com/go-playground/validator/v10"

	"github.com/eko/gocache/v2/marshaler"

	"github.com/eko/gocache/v2/store"

	"github.com/labstack/echo/v4"
)

var (
	// templates stores a cache of parsed page templates
	templates = sync.Map{}

	// funcMap stores the Template function map
	funcMap = funcmap.GetFuncMap()

	// templatePath stores the complete path to the templates directory
	templatePath = getTemplatesDirectoryPath()
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
func (t *Controller) RenderPage(c echo.Context, p Page) error {
	// Page name is required
	if p.Name == "" {
		c.Logger().Error("page render failed due to missing name")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// Use the app name in configuration if a value was not set
	if p.AppName == "" {
		p.AppName = t.Container.Config.App.Name
	}

	// Parse the templates in the page and store them in a cache, if not yet done
	if err := t.parsePageTemplates(p); err != nil {
		c.Logger().Errorf("failed to parse templates: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// Execute the parsed templates to render the page
	buf, err := t.executeTemplates(c, p)
	if err != nil {
		c.Logger().Errorf("failed to execute templates: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// Cache this page, if caching was enabled
	t.cachePage(c, p, buf)

	// Set any headers
	for k, v := range p.Headers {
		c.Response().Header().Set(k, v)
	}

	return c.HTMLBlob(p.StatusCode, buf.Bytes())
}

// cachePage caches the HTML for a given Page if the Page has caching enabled
func (t *Controller) cachePage(c echo.Context, p Page, html *bytes.Buffer) {
	if !p.Cache.Enabled {
		return
	}

	// If no expiration time was provided, default to the configuration value
	if p.Cache.Expiration == 0 {
		p.Cache.Expiration = t.Container.Config.Cache.Expiration.Page
	}

	// The request URL is used as the cache key so the middleware can serve the
	// cached page on matching requests
	key := c.Request().URL.String()
	opts := &store.Options{
		Expiration: p.Cache.Expiration,
		Tags:       p.Cache.Tags,
	}
	cp := middleware.CachedPage{
		URL:        key,
		HTML:       html.Bytes(),
		Headers:    p.Headers,
		StatusCode: p.StatusCode,
	}

	err := marshaler.New(t.Container.Cache).Set(c.Request().Context(), key, cp, opts)
	if err != nil {
		c.Logger().Errorf("failed to cache page: %v", err)
		return
	}

	c.Logger().Infof("cached page")
}

// parsePageTemplates parses the templates for the given Page and caches them to avoid duplicate operations
// If the configuration indicates that the environment is local, the cache is bypassed for template changes
// can be seen without having to restart the application.
// As mentioned in the documentation for the Page struct, the templates used for the page will be:
// 1. The layout/based template specified in Page.Layout
// 2. The content template specified in Page.Name
// 3. All templates within the components directory
func (t *Controller) parsePageTemplates(p Page) error {
	// Check if the template has not yet been parsed or if the app environment is local, so that templates reflect
	// changes without having the restart the server
	if _, ok := templates.Load(p.Name); !ok || t.Container.Config.App.Environment == config.EnvLocal {
		// Parse the Layout and Name templates
		parsed, err :=
			template.New(p.Layout+config.TemplateExt).
				Funcs(funcMap).
				ParseFiles(
					fmt.Sprintf("%s/layouts/%s%s", templatePath, p.Layout, config.TemplateExt),
					fmt.Sprintf("%s/pages/%s%s", templatePath, p.Name, config.TemplateExt),
				)

		if err != nil {
			return err
		}

		// Parse all templates within the components directory
		parsed, err = parsed.ParseGlob(fmt.Sprintf("%s/components/*%s", templatePath, config.TemplateExt))

		if err != nil {
			return err
		}

		// Store the template so this process only happens once
		templates.Store(p.Name, parsed)
	}

	return nil
}

// executeTemplates executes the cached templates belonging to Page and renders the Page within them
func (t *Controller) executeTemplates(c echo.Context, p Page) (*bytes.Buffer, error) {
	tmpl, ok := templates.Load(p.Name)
	if !ok {
		return nil, errors.New("uncached page template requested")
	}

	buf := new(bytes.Buffer)
	err := tmpl.(*template.Template).ExecuteTemplate(buf, p.Layout+config.TemplateExt, p)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// Redirect redirects to a given route name with optional route parameters
func (t *Controller) Redirect(c echo.Context, route string, routeParams ...interface{}) error {
	return c.Redirect(http.StatusFound, c.Echo().Reverse(route, routeParams))
}

// SetValidationErrorMessages sets error flash messages for validation failures of a given struct
// and attempts to provide more user-friendly wording.
// The error should result from the validator module and the data should be the struct that failed
// validation.
// This method supports including a struct tag of "labeL" on each field which will be the name
// of the field used in the error messages, for example:
//  - FirstName string `form:"first-name" validate:"required" label:"First name"`
// Only a few validator tags are supported below. Expand them as needed.
func (t *Controller) SetValidationErrorMessages(c echo.Context, err error, data interface{}) {
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

		msg.Danger(c, fmt.Sprintf(message, "<strong>"+label+"</strong>"))
	}
}

// getTemplatesDirectoryPath gets the templates directory path
// This is needed incase this is called from a package outside of main,
// such as within tests
func getTemplatesDirectoryPath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Join(filepath.Dir(d), config.TemplateDir)
}
