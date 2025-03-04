package layouts

import (
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/cache"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Primary(r *ui.Request, content Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Metatags(r),
				CSS(),
				JS(r),
			),
			Body(
				headerNavBar(r),
				Div(
					Class("container mt-5"),
					Div(
						Class("columns"),
						Div(
							Class("column is-2"),
							sidebarMenu(r),
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
		),
	)
}

func headerNavBar(r *ui.Request) Node {
	const cacheKey = "layout.headerNavBar"
	if n := cache.Get(cacheKey); n != nil {
		return n
	}

	n := Nav(
		Class("navbar is-dark"),
		Div(
			Class("container"),
			Div(
				Class("navbar-brand"),
				HxBoost(),
				A(
					Href(r.Path(routenames.Home)),
					Class("navbar-item"),
					Text("Pagoda"),
				),
			),
			Div(
				ID("navbarMenu"),
				Class("navbar-menu"),
				Div(
					Class("navbar-end"),
					search(r),
				),
			),
		),
	)
	cache.Set(cacheKey, n)
	return n
}

func search(r *ui.Request) Node {
	const cacheKey = "layout.search"
	if n := cache.Get("layout.search"); n != nil {
		return n
	}

	n := Div(
		Class("search mr-2 mt-1"),
		Attr("x-data", "{modal:false}"),
		Input(
			Class("input"),
			Type("search"),
			Placeholder("Search..."),
			Attr("@click", "modal = true; $nextTick(() => $refs.input.focus());"),
		),
		Div(
			Class("modal"),
			Attr(":class", "modal ? 'is-active' : ''"),
			Attr("x-show", "modal == true"),
			Div(
				Class("modal-background"),
			),
			Div(
				Class("modal-content"),
				Attr("@click.outside", "modal = false;"),
				Div(
					Class("box"),
					H2(
						Class("subtitle"),
						Text("Search"),
					),
					P(
						Class("control"),
						Input(
							Attr("hx-get", r.Path(routenames.Search)),
							Attr("hx-trigger", "keyup changed delay:500ms"),
							Attr("hx-target", "#results"),
							Name("query"),
							Class("input"),
							Type("search"),
							Placeholder("Search..."),
							Attr("x-ref", "input"),
						),
					),
					Div(
						Class("block"),
					),
					Div(
						ID("results"),
					),
				),
			),
			Button(
				Class("modal-close is-large"),
				Aria("label", "close"),
			),
		),
	)
	cache.Set(cacheKey, n)
	return n
}

func sidebarMenu(r *ui.Request) Node {
	return Aside(
		Class("menu"),
		HxBoost(),
		P(
			Class("menu-label"),
			Text("General"),
		),
		Ul(
			Class("menu-list"),
			MenuLink(r, "Dashboard", routenames.Home),
			MenuLink(r, "About", routenames.About),
			MenuLink(r, "Contact", routenames.Contact),
			MenuLink(r, "Cache", routenames.Cache),
			MenuLink(r, "Task", routenames.Task),
			MenuLink(r, "Files", routenames.Files),
		),
		P(
			Class("menu-label"),
			Text("Account"),
		),
		Ul(
			Class("menu-list"),
			If(r.IsAuth, MenuLink(r, "Logout", routenames.Logout)),
			If(!r.IsAuth, MenuLink(r, "Login", routenames.Login)),
			If(!r.IsAuth, MenuLink(r, "Register", routenames.Register)),
			If(!r.IsAuth, MenuLink(r, "Forgot password", routenames.ForgotPasswordSubmit)),
		),
	)
}
