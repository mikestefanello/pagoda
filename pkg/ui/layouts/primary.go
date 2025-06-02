package layouts

import (
	"github.com/mikestefanello/pagoda/ent/admin"
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
			Data("theme", "dark"),
			Head(
				Metatags(r),
				CSS(),
				JS(),
			),
			Body(
				Div(
					Class("drawer lg:drawer-open"),
					Input(
						ID("sidebar"),
						Type("checkbox"),
						Class("drawer-toggle"),
					),
					Div(
						Class("drawer-content flex flex-col p-7 prose-sm"),
						If(len(r.Title) > 0, H1(Text(r.Title))),
						FlashMessages(r),
						content,
						Label(
							For("sidebar"),
							Class("btn btn-primary drawer-button lg:hidden"),
							Text("Open drawer"),
						),
					),
					sidebarMenu(r),
				),
				//headerNavBar(r),
				//Div(
				//	Class("container mt-5"),
				//	Div(
				//		Class("columns"),
				//		Div(
				//			Class("column is-2"),
				//			sidebarMenu(r),
				//		),
				//		Div(
				//			Class("column is-10"),
				//			Div(
				//				Class("box"),
				//				If(len(r.Title) > 0, H1(Class("title"), Text(r.Title))),
				//				FlashMessages(r),
				//				content,
				//			),
				//		),
				//	),
				//),
				HtmxListeners(r),
			),
		),
	)
}

func headerNavBar(r *ui.Request) Node {
	return cache.SetIfNotExists("layout.headerNavBar", func() Node {
		return Nav(
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
	})
}

func search(r *ui.Request) Node {
	return cache.SetIfNotExists("layout.search", func() Node {
		return Div(
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
	})
}

func sidebarMenu(r *ui.Request) Node {
	header := func(text string) Node {
		return Li(
			Class("menu-title mt-3 uppercase"),
			Span(Text(text)),
		)
	}

	adminSubMenu := func() Node {
		entityTypeNames := admin.GetEntityTypeNames()
		entityTypeLinks := make(Group, len(entityTypeNames))
		for _, n := range entityTypeNames {
			entityTypeLinks = append(entityTypeLinks, MenuLink(r, n, routenames.AdminEntityList(n)))
		}

		return Group{
			header("Entities"),
			entityTypeLinks,
			header("Monitoring"),
			Li(
				A(
					Href(r.Path(routenames.AdminTasks)),
					Text("Tasks"),
					Target("_blank"),
				),
			),
		}
	}

	return Div(
		Class("drawer-side"),
		HxBoost(),
		Label(
			For("sidebar"),
			Aria("label", "close sidebar"),
			Class("drawer-overlay"),
		),
		Ul(
			Class("menu bg-base-200 text-base-content min-h-full w-80 p-4"),
			header("General"),
			MenuLink(r, "Dashboard", routenames.Home),
			MenuLink(r, "About", routenames.About),
			MenuLink(r, "Contact", routenames.Contact),
			MenuLink(r, "Cache", routenames.Cache),
			MenuLink(r, "Task", routenames.Task),
			MenuLink(r, "Files", routenames.Files),
			header("Account"),
			If(r.IsAuth, MenuLink(r, "Logout", routenames.Logout)),
			If(!r.IsAuth, MenuLink(r, "Login", routenames.Login)),
			If(!r.IsAuth, MenuLink(r, "Register", routenames.Register)),
			If(!r.IsAuth, MenuLink(r, "Forgot password", routenames.ForgotPasswordSubmit)),
			Iff(r.IsAdmin, adminSubMenu),
		),
	)
}
