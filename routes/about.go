package routes

import (
	"html/template"

	"goweb/controller"

	"github.com/labstack/echo/v4"
)

type (
	About struct {
		controller.Controller
	}

	AboutData struct {
		ShowCacheWarning bool
		Tabs             []AboutTab
	}

	AboutTab struct {
		Title string
		Body  template.HTML
	}
)

func (c *About) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "about"
	page.Title = "About"

	// This page will be cached!
	page.Cache.Enabled = true
	page.Cache.Tags = []string{"page_about", "page:list"}

	// A simple example of how the Data field can contain anything you want to send to the templates
	page.Data = AboutData{
		ShowCacheWarning: true,
		Tabs: []AboutTab{
			{
				Title: "HTMX",
				Body:  template.HTML(`Completes HTML as a hypertext by providing attributes to AJAXify anything and much more. Visit <a href="https://htmx.org/">htmx.org</a> to learn more.`),
			},
			{
				Title: "Alpine.js",
				Body:  template.HTML(`Drop-in, Vue-like functionality written directly in your markup. Visit <a href="https://alpinejs.dev/">alpinejs.dev</a> to learn more.`),
			},
			{
				Title: "Bulma",
				Body:  template.HTML(`Ready-to-use frontend components that you can easily combine to build responsive web interfaces with no JavaScript requirements. Visit <a href="https://bulma.io/">bulma.io</a> to learn more.`),
			},
		},
	}

	return c.RenderPage(ctx, page)
}
