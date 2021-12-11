package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"goweb/config"
	"goweb/container"
	"goweb/funcmap"
	"goweb/middleware"

	"github.com/eko/gocache/v2/marshaler"

	"github.com/eko/gocache/v2/store"

	"github.com/labstack/echo/v4"
)

var (
	// Cache of compiled page templates
	templates = sync.Map{}

	// Template function map
	funcMap = funcmap.GetFuncMap()

	templatePath = getTemplatesDirectoryPath()
)

type Controller struct {
	Container *container.Container
}

func NewController(c *container.Container) Controller {
	return Controller{
		Container: c,
	}
}

// TODO: Audit error handling (ie NewHTTPError)

func (t *Controller) RenderPage(c echo.Context, p Page) error {
	if p.Name == "" {
		c.Logger().Error("page render failed due to missing name")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	if p.AppName == "" {
		p.AppName = t.Container.Config.App.Name
	}

	if err := t.parsePageTemplates(p); err != nil {
		return err
	}

	buf, err := t.executeTemplates(c, p)
	if err != nil {
		return err
	}

	t.cachePage(c, p, buf)

	// Set any headers
	for k, v := range p.Headers {
		c.Response().Header().Set(k, v)
	}

	return c.HTMLBlob(p.StatusCode, buf.Bytes())
}

func (t *Controller) cachePage(c echo.Context, p Page, html *bytes.Buffer) {
	if !p.Cache.Enabled {
		return
	}

	if p.Cache.MaxAge == 0 {
		p.Cache.MaxAge = t.Container.Config.Cache.Expiration.Page
	}

	key := c.Request().URL.String()
	opts := &store.Options{
		Expiration: p.Cache.MaxAge,
		Tags:       p.Cache.Tags,
	}
	cp := middleware.CachedPage{
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

func (t *Controller) parsePageTemplates(p Page) error {
	// Check if the template has not yet been parsed or if the app environment is local, so that templates reflect
	// changes without having the restart the server
	if _, ok := templates.Load(p.Name); !ok || t.Container.Config.App.Environment == config.EnvLocal {
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

		parsed, err = parsed.ParseGlob(fmt.Sprintf("%s/components/*%s", templatePath, config.TemplateExt))

		if err != nil {
			return err
		}

		// Store the template so this process only happens once
		templates.Store(p.Name, parsed)
	}

	return nil
}

func (t *Controller) executeTemplates(c echo.Context, p Page) (*bytes.Buffer, error) {
	tmpl, ok := templates.Load(p.Name)
	if !ok {
		c.Logger().Error("uncached page template requested")
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	buf := new(bytes.Buffer)
	err := tmpl.(*template.Template).ExecuteTemplate(buf, p.Layout+config.TemplateExt, p)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (t *Controller) Redirect(c echo.Context, route string, routeParams ...interface{}) error {
	return c.Redirect(http.StatusFound, c.Echo().Reverse(route, routeParams))
}

// getTemplatesDirectoryPath gets the templates directory path
// This is needed incase this is called from a package outside of main,
// such as testing
func getTemplatesDirectoryPath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Join(filepath.Dir(d), config.TemplateDir)
}
