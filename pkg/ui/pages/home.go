package pages

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	"github.com/mikestefanello/pagoda/pkg/ui/models"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(ctx echo.Context, posts *models.Posts) error {
	r := ui.NewRequest(ctx)
	r.Metatags.Description = "This is the home page."
	r.Metatags.Keywords = []string{"Software", "Coding", "Go"}

	g := make(Group, 0)

	if r.Htmx.Target != "posts" {
		var hello string
		if r.IsAuth {
			hello = fmt.Sprintf("Hello, %s", r.AuthUser.Name)
		} else {
			hello = "Hello"
		}

		g = append(g,
			Section(
				Class("hero is-info welcome is-small mb-5"),
				Div(
					Class("hero-body"),
					Div(
						Class("container"),
						H1(
							Class("title"),
							Text(hello),
						),
						H2(
							Class("subtitle"),
							If(!r.IsAuth, Text("Please login in to your account.")),
							If(r.IsAuth, Text("Welcome back!")),
						),
					),
				),
			),
			H2(Class("title"), Text("Recent posts")),
			H3(Class("subtitle"), Text("Below is an example of both paging and AJAX fetching using HTMX")),
		)
	}

	g = append(g, posts.Render(r.Path(routenames.Home)))

	if r.Htmx.Target != "posts" {
		g = append(g, Message(
			"is-small is-warning mt-5",
			"Serving files",
			Group{
				Text("In the example posts above, check how the file URL contains a cache-buster query parameter which changes only when the app is restarted. "),
				Text("Static files also contain cache-control headers which are configured via middleware."),
			},
		))
	}

	return r.Render(layouts.Primary, g)
}
