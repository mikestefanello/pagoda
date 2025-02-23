package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui"
)

type Pages struct{}

func init() {
	Register(new(Pages))
}

func (h *Pages) Init(c *services.Container) error {
	return nil
}

func (h *Pages) Routes(g *echo.Group) {
	g.GET("/", h.Home).Name = routenames.Home
	g.GET("/about", h.About).Name = routenames.About
}

func (h *Pages) Home(ctx echo.Context) error {
	pgr := page.NewPager(ctx, 4)

	return ui.Home(ctx, ui.Posts{
		Posts: h.fetchPosts(&pgr),
		Pager: pgr,
	})
}

// fetchPosts is a mock example of fetching posts to illustrate how paging works.
func (h *Pages) fetchPosts(pager *page.Pager) []ui.Post {
	pager.SetItems(20)
	posts := make([]ui.Post, 20)

	for k := range posts {
		posts[k] = ui.Post{
			Title: fmt.Sprintf("Post example #%d", k+1),
			Body:  fmt.Sprintf("Lorem ipsum example #%d ddolor sit amet, consectetur adipiscing elit. Nam elementum vulputate tristique.", k+1),
		}
	}
	return posts[pager.GetOffset() : pager.GetOffset()+pager.ItemsPerPage]
}

func (h *Pages) About(ctx echo.Context) error {
	return ui.About(ctx)
}
