package layouts

import (
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/cache"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/icons"
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
						Class("drawer-content flex flex-col p-7 prose-base"),
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
				searchModal(r),
				HtmxListeners(r),
			),
		),
	)
}

func search() Node {
	return cache.SetIfNotExists("layout.search", func() Node {
		return Div(
			Class("ml-2"),
			Attr("x-data", ""),
			Label(
				Class("input"),
				icons.MagnifyingGlass(),
				Input(
					Type("search"),
					Class("grow"),
					Placeholder("Search"),
					Attr("@click", "search_modal.showModal();"),
				),
			),

		)
	})
}

func searchModal(r *ui.Request) Node {
	return cache.SetIfNotExists("layout.searchModal", func() Node {
		return Dialog(
			ID("search_modal"),
			Class("modal"),
			Div(
				Class("modal-box"),
				Form(
					Method("dialog"),
					Button(
						Class("btn btn-sm btn-circle btn-ghost absolute right-2 top-2"),
						Text("✕"),
					),
				),
				H3(
					Class("text-lg font-bold mb-2"),
					Text("Search"),
				),
				Input(
					Attr("hx-get", r.Path(routenames.Search)),
					Attr("hx-trigger", "keyup changed delay:500ms"),
					Attr("hx-target", "#results"),
					Name("query"),
					Class("input w-full"),
					Type("search"),
					Placeholder("Search..."),
				),
				Ul(
					ID("results"),
					Class("list"),
				),
			),
			Form(
				Method("dialog"),
				Class("modal-backdrop"),
				Button(
					Text("close"),
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
		entityTypeLinks := make(Group, len(admin.GetSchema()))
		for _, n := range admin.GetSchema() {
			entityTypeLinks = append(
				entityTypeLinks,
				MenuLink(r, icons.PencilSquare(), n.Name, routenames.AdminEntityList(n.Name)),
			)
		}

		return Group{
			header("Entities"),
			entityTypeLinks,
			header("Monitoring"),
			Li(
				A(
					icons.CircleStack(),
					Href(r.Path(routenames.AdminTasks)),
					Text("Tasks"),
					Target("_blank"),
				),
			),
		}
	}

	return Div(
		Class("drawer-side"),
		Label(
			For("sidebar"),
			Aria("label", "close sidebar"),
			Class("drawer-overlay"),
		),
		Div(
			Class("menu bg-base-200 text-base-content min-h-full w-80 p-4"),
			Div(
				Class("w-2/3 mx-auto mt-3 mb-10"),
				Img(
					Src(ui.StaticFile("logo.png")),
				),
			),
			search(),
			Ul(
				HxBoost(),
				header("General"),
				MenuLink(r, icons.Home(), "Dashboard", routenames.Home),
				MenuLink(r, icons.Info(), "About", routenames.About),
				MenuLink(r, icons.Mail(), "Contact", routenames.Contact),
				MenuLink(r, icons.Archive(), "Cache", routenames.Cache),
				MenuLink(r, icons.CircleStack(), "Task", routenames.Task),
				MenuLink(r, icons.Document(), "Files", routenames.Files),
				header("Account"),
				If(r.IsAuth, MenuLink(r, icons.Exit(), "Logout", routenames.Logout)),
				If(!r.IsAuth, MenuLink(r, icons.Enter(), "Login", routenames.Login)),
				If(!r.IsAuth, MenuLink(r, icons.UserPlus(), "Register", routenames.Register)),
				If(!r.IsAuth, MenuLink(r, icons.QuestionCircle(), "Forgot password", routenames.ForgotPasswordSubmit)),
				Iff(r.IsAdmin, adminSubMenu),
			),
		),
	)
}
