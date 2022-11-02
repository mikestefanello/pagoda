package routes

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/controller"

	"github.com/labstack/echo/v4"
)

type (
	home struct {
		controller.Controller
	}

	post struct {
		Title string
		Body  string
	}
)

func (c *home) Get(ctx echo.Context) error {
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
func (c *home) fetchPosts(pager *controller.Pager) []post {
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
