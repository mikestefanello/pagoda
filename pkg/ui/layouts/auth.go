package layouts

import (
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/cache"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Auth(r *ui.Request, content Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Metatags(r),
				CSS(),
				JS(r),
			),
			Body(
				Section(
					Class("hero is-fullheight"),
					Div(
						Class("hero-body"),
						Div(
							Class("container"),
							Div(
								Class("columns is-centered"),
								Div(
									Class("column is-half"),
									If(len(r.Title) > 0, H1(Class("title"), Text(r.Title))),
									Div(
										Class("notification"),
										FlashMessages(r),
										content,
										authNavBar(r),
									),
								),
							),
						),
					),
				),
			),
		),
	)
}

func authNavBar(r *ui.Request) Node {
	return cache.SetIfNotExists("authNavBar", func() Node {
		return Nav(
			Class("navbar"),
			Div(
				Class("navbar-menu"),
				Div(
					Class("navbar-start"),
					A(Class("navbar-item"), Href(r.Path(routenames.Login)), Text("Login")),
					A(Class("navbar-item"), Href(r.Path(routenames.Register)), Text("Create an account")),
					A(Class("navbar-item"), Href(r.Path(routenames.ForgotPassword)), Text("Forgot password")),
				),
			),
		)
	})
}
