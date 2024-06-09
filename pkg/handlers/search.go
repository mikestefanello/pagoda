package handlers

import (
	"fmt"
	"math/rand"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/templates"
)

const routeNameSearch = "search"

type (
	Search struct {
		controller.Controller
	}

	searchResult struct {
		Title string
		URL   string
	}
)

func init() {
	Register(new(Search))
}

func (c *Search) Init(ct *services.Container) error {
	c.Controller = controller.NewController(ct)
	return nil
}

func (c *Search) Routes(g *echo.Group) {
	g.GET("/search", c.Page).Name = routeNameSearch
}

func (c *Search) Page(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutMain
	page.Name = templates.PageSearch

	// Fake search results
	var results []searchResult
	if search := ctx.QueryParam("query"); search != "" {
		for i := 0; i < 5; i++ {
			title := "Lorem ipsum example ddolor sit amet"
			index := rand.Intn(len(title))
			title = title[:index] + search + title[index:]
			results = append(results, searchResult{
				Title: title,
				URL:   fmt.Sprintf("https://www.%s.com", search),
			})
		}
	}
	page.Data = results

	return c.RenderPage(ctx, page)
}
