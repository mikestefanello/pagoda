package layouts

import (
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Admin(r *ui.Request, content Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Metatags(r),
				CSS(),
				JS(r),
			),
			Body(
				Div(
					Class("box"),
					Div(
						Class("columns"),
						Div(
							Class("column is-2"),
							adminMenu(r),
						),
						Div(
							Class("column is-10"),
							If(len(r.Title) > 0, H1(Class("title"), Text(r.Title))),
							FlashMessages(r),
							content,
						),
					),
				),
			),
			HtmxListeners(r),
		),
	)
}

func adminMenu(r *ui.Request) Node {
	return Aside(
		Class("menu"),
		HxBoost(),
		P(
			Class("menu-label"),
			Text("Content"),
		),
		Ul(
			Class("menu-list"),
			MenuLink(r, "Users", "admin:user_list"),
			MenuLink(r, "Tokens", "admin:passwordtoken_list"),
		),
		P(
			Class("menu-label"),
			Text("Account"),
		),
		Ul(
			Class("menu-list"),
			If(r.IsAuth, MenuLink(r, "Logout", routenames.Logout)),
		),
	)
}
