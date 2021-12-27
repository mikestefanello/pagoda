package routes

import (
	"fmt"
	"math/rand"

	"goweb/controller"

	"github.com/labstack/echo/v4"
)

type (
	Search struct {
		controller.Controller
	}

	SearchResult struct {
		Title string
		URL   string
	}
)

func (c *Search) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "search"

	// Fake search results
	var results []SearchResult
	if search := ctx.QueryParam("query"); search != "" {
		for i := 0; i < 5; i++ {
			title := "Lorem ipsum example ddolor sit amet"
			index := rand.Intn(len(title))
			title = title[:index] + search + title[index:]
			results = append(results, SearchResult{
				Title: title,
				URL:   fmt.Sprintf("https://www.%s.com", search),
			})
		}
	}
	page.Data = results

	return c.RenderPage(ctx, page)
}
