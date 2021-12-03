package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"goweb/config"
	"goweb/container"
	"goweb/funcmap"

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
)

type Controller struct {
	Container *container.Container
}

func NewController(c *container.Container) Controller {
	return Controller{
		Container: c,
	}
}

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

	tmpl, ok := templates.Load(p.Name)
	if !ok {
		c.Logger().Error("Uncached page template requested")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	buf := new(bytes.Buffer)
	err := tmpl.(*template.Template).ExecuteTemplate(buf, p.Layout+TemplateExt, p)
	if err != nil {
		return err
	}

	return c.HTMLBlob(p.StatusCode, buf.Bytes())
}

func (t *Controller) parsePageTemplates(p Page) error {
	// Check if the template has not yet been parsed or if the app environment is local, so that templates reflect
	// changes without having the restart the server
	if _, ok := templates.Load(p.Name); !ok || t.Container.Config.App.Environment == config.EnvLocal {
		parsed, err :=
			template.New(p.Layout+TemplateExt).
				Funcs(funcMap).
				ParseFiles(
					fmt.Sprintf("%s/layouts/%s%s", TemplateDir, p.Layout, TemplateExt),
					fmt.Sprintf("%s/pages/%s%s", TemplateDir, p.Name, TemplateExt),
				)

		if err != nil {
			return err
		}

		parsed, err = parsed.ParseGlob(fmt.Sprintf("%s/components/*%s", TemplateDir, TemplateExt))

		if err != nil {
			return err
		}

		// Store the template so this process only happens once
		templates.Store(p.Name, parsed)
	}

	return nil
}

func (t *Controller) Redirect(c echo.Context, route string, routeParams ...interface{}) error {
	return c.Redirect(http.StatusFound, c.Echo().Reverse(route, routeParams))
}
