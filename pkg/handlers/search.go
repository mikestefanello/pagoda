package handlers

import (
	"fmt"
	"math/rand"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/models"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

type Search struct{}

func init() {
	Register(new(Search))
}

func (h *Search) Init(c *services.Container) error {
	return nil
}

func (h *Search) Routes(g *echo.Group) {
	g.GET("/search", h.Page).Name = routenames.Search
}

func (h *Search) Page(ctx echo.Context) error {
	// Fake search results.
	results := make([]*models.SearchResult, 0, 5)
	if search := ctx.QueryParam("query"); search != "" {
		for i := 0; i < 5; i++ {
			title := "Lorem ipsum example ddolor sit amet"
			index := rand.Intn(len(title))
			title = title[:index] + search + title[index:]
			results = append(results, &models.SearchResult{
				Title: title,
				URL:   fmt.Sprintf("https://www.%s.com", search),
			})
		}
	}

	return pages.SearchResults(ctx, results)
}
