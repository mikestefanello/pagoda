package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	"github.com/mikestefanello/pagoda/pkg/ui/models"
	. "maragu.dev/gomponents"
)

func SearchResults(ctx echo.Context, results []*models.SearchResult) error {
	r := ui.NewRequest(ctx)

	g := make(Group, 0, len(results))
	for _, result := range results {
		g = append(g, result.Render())
	}

	return r.Render(layouts.Primary, g)
}
