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

	// This pages helps to illustrate the different options you can take when using HTMX to introduce interactivity
	// to your web application. The following three options are available, but here, we're opting for the first one.
	// 1) Highly-optimized and progressive enhancement:
	//    This is highly-optimized because the server is doing the least amount of work possible, only rendering
	//    the least amount possible based on the incoming request. It's possible that even your route handler would
	//    want to check the HTMX request in order to limit what it does. With HTMX, it's possible to still return a
	//    normal, full page, but use hx-select to pluck out only the part you want to re-render. It requires some extra
	//    condition checks and code but performance is improved. Progressive enhancement refers to having a fully
	//    functional web app, even if JS was disabled, but providing the enhancement if JS is enabled. All of these
	//    examples should continue to work fine without JS.
	// 2) Not optimized and progressive enhancement:
	//    As mentioned previously, you can remove all of these conditions, re-render the entire page for every request,
	//    and rely on HTMX's hx-select to only replace what you want to (ie, the posts).
	// 3) Optimized and partial renderings:
	//    You could have a separate route that is only for fetching posts while paging, and that would render only
	//    that partial HTML, which HTMX would then use to inject in to this page.

	headerMsg := func() Node {
		return Group{
			Section(
				Class("hero is-info welcome is-small mb-5"),
				Div(
					Class("hero-body"),
					Div(
						Class("container"),
						H1(
							Class("title"),
							Iff(r.IsAuth, func() Node {
								return Text(fmt.Sprintf("Hello, %s", r.AuthUser.Name))
							}),
							If(!r.IsAuth, Text("Hello")),
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
		}
	}

	filesMsg := func() Node {
		return Message(
			"is-small is-warning mt-5",
			"Serving files",
			Group{
				Text("In the example posts above, check how the file URL contains a cache-buster query parameter which changes only when the app is restarted. "),
				Text("Static files also contain cache-control headers which are configured via middleware."),
			},
		)
	}

	g := Group{
		Iff(r.Htmx.Target != "posts", headerMsg),
		posts.Render(r.Path(routenames.Home)),
		Iff(r.Htmx.Target != "posts", filesMsg),
	}

	return r.Render(layouts.Primary, g)
}
