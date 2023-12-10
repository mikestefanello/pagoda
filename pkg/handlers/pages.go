package handlers

import (
	"fmt"
	"html/template"

	"github.com/mikestefanello/pagoda/pkg/controller"

	"github.com/labstack/echo/v4"
)

const (
	routeNameAbout = "about"
	routeNameHome  = "home"
)

type (
	Pages struct {
		controller.Controller
	}

	post struct {
		Title string
		Body  string
	}

	aboutData struct {
		ShowCacheWarning bool
		FrontendTabs     []aboutTab
		BackendTabs      []aboutTab
	}

	aboutTab struct {
		Title string
		Body  template.HTML
	}
)

func (c *Pages) Routes(g *echo.Group) {
	g.GET("/", c.Home).Name = routeNameHome
	g.GET("/about", c.About).Name = routeNameAbout
}

func (c *Pages) Home(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "home"
	page.Metatags.Description = "Welcome to the homepage."
	page.Metatags.Keywords = []string{"Go", "MVC", "Web", "Software"}
	page.Pager = controller.NewPager(ctx, 4)
	page.Data = c.fetchPosts(&page.Pager)

	return c.RenderPage(ctx, page)
}

// fetchPosts is an mock example of fetching posts to illustrate how paging works
func (c *Pages) fetchPosts(pager *controller.Pager) []post {
	pager.SetItems(20)
	posts := make([]post, 20)

	for k := range posts {
		posts[k] = post{
			Title: fmt.Sprintf("Post example #%d", k+1),
			Body:  fmt.Sprintf("Lorem ipsum example #%d ddolor sit amet, consectetur adipiscing elit. Nam elementum vulputate tristique.", k+1),
		}
	}
	return posts[pager.GetOffset() : pager.GetOffset()+pager.ItemsPerPage]
}

func (c *Pages) About(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "about"
	page.Title = "About"

	// This page will be cached!
	page.Cache.Enabled = true
	page.Cache.Tags = []string{"page_about", "page:list"}

	// A simple example of how the Data field can contain anything you want to send to the templates
	// even though you wouldn't normally send markup like this
	page.Data = aboutData{
		ShowCacheWarning: true,
		FrontendTabs: []aboutTab{
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
		BackendTabs: []aboutTab{
			{
				Title: "Echo",
				Body:  template.HTML(`High performance, extensible, minimalist Go web framework. Visit <a href="https://echo.labstack.com/">echo.labstack.com</a> to learn more.`),
			},
			{
				Title: "Ent",
				Body:  template.HTML(`Simple, yet powerful ORM for modeling and querying data. Visit <a href="https://entgo.io/">entgo.io</a> to learn more.`),
			},
		},
	}

	return c.RenderPage(ctx, page)
}
