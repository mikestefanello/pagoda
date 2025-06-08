package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/icons"
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
			Stats(
				Stat{
					Title: "User name",
					Value: func() string {
						if r.IsAuth {
							return r.AuthUser.Name
						}
						return "(not logged in)"
					}(),
					Description: "The logged in user's name",
					Icon:        icons.UserCircle(),
				},
				Stat{
					Title: "Admin status",
					Value: func() string {
						if r.IsAdmin {
							return "Administrator"
						}
						return "Non-administrator"
					}(),
					Description: "Use `make admin` to create an admin account",
					Icon:        icons.LockClosed(),
				},
				Stat{
					Title:       "GitHub Stars",
					Value:       "2,500+",
					Description: "Star if you like Pagoda",
					Icon:        icons.Star(),
				},
			),
			H2(Text("Recent posts")),
			Span(Text("Below is an example of both paging and AJAX fetching using HTMX")),
		}
	}

	cards := func() Node {
		return Div(
			Class("flex w-full gap-2 mt-5"),
			Card(CardParams{
				Title: "Serving files",
				Body: Group{
					Text("In the example posts above, check how the file URL contains a cache-buster query parameter which changes only when the app is restarted. "),
					Text("Static files also contain cache-control headers which are configured via middleware."),
				},
				Color: ColorWarning,
				Size:  SizeSmall,
			}),
			Card(CardParams{
				Title: "Documentation",
				Body: Group{
					Text("Have you read through the entire documentation? If not, you may be missing functionality or have questions. "),
				},
				Footer: Group{
					ButtonLink(ColorNeutral, "https://github.com/mikestefanello/pagoda?tab=readme-ov-file#table-of-contents", "Learn more"),
				},
				Color: ColorNeutral,
				Size:  SizeSmall,
			}),
		)
	}

	g := Group{
		Iff(r.Htmx.Target != "posts", headerMsg),
		posts.Render(r.Path(routenames.Home)),
		Iff(r.Htmx.Target != "posts", cards),
	}

	return r.Render(layouts.Primary, g)
}
