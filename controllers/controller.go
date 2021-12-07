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

	"github.com/eko/gocache/v2/store"

	"github.com/labstack/echo/v4"
)

const (
	TemplateDir = "views"
	TemplateExt = ".gohtml"
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
		c.Logger().Error("Page render failed due to missing name")
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

	t.cachePage(c, p)

	return c.HTMLBlob(p.StatusCode, buf.Bytes())
}

func (t *Controller) cachePage(c echo.Context, p Page) {
	if !p.Cache.Enabled {
		return
	}

	if p.Cache.MaxAge == 0 {
		p.Cache.MaxAge = t.Container.Config.Cache.MaxAge.Page
	}

	key := c.Request().URL.String()
	opts := &store.Options{
		Expiration: p.Cache.MaxAge,
		Tags:       p.Cache.Tags,
	}
	err := t.Container.Cache.Set(c.Request().Context(), key, "my-value", opts)
	if err != nil {
		c.Logger().Errorf("Failed to cache page: %s", key)
		c.Logger().Error(err)
		return
	}
}

func (t *Controller) parsePageTemplates(p Page) error {
	// Check if the template has not yet been parsed or if the app environment is local, so that templates reflect
	// changes without having the restart the server
	if _, ok := templates.Load(p.Name); !ok || t.Container.Config.App.Environment == config.EnvLocal {
		parsed, err :=
			template.New(p.Layout+TemplateExt).
				Funcs(funcMap).
				ParseFiles(
					fmt.Sprintf("%s/layouts/%s%s", templatePath, p.Layout, TemplateExt),
					fmt.Sprintf("%s/pages/%s%s", templatePath, p.Name, TemplateExt),
				)

		if err != nil {
			return err
		}

		parsed, err = parsed.ParseGlob(fmt.Sprintf("%s/components/*%s", templatePath, TemplateExt))

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
		c.Logger().Error("Uncached page template requested")
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	buf := new(bytes.Buffer)
	err := tmpl.(*template.Template).ExecuteTemplate(buf, p.Layout+TemplateExt, p)
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
	return filepath.Join(filepath.Dir(d), TemplateDir)
}
