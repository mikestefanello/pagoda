package routes

import (
	"fmt"

	"goweb/controller"

	"github.com/labstack/echo/v4"
)

type (
	Home struct {
		controller.Controller
	}

	Post struct {
		Title string
		Body  string
	}
)

func (c *Home) Get(ctx echo.Context) error {
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
func (c *Home) fetchPosts(pager *controller.Pager) []Post {
	pager.SetItems(20)
	posts := make([]Post, 20)

	for k := range posts {
		posts[k] = Post{
			Title: fmt.Sprintf("Post example #%d", k+1),
			Body:  fmt.Sprintf("Lorem ipsum example #%d ddolor sit amet, consectetur adipiscing elit. Nam elementum vulputate tristique.", k+1),
		}
	}
	return posts[pager.GetOffset() : pager.GetOffset()+pager.ItemsPerPage]
}
