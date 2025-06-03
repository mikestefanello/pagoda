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
			Data("theme", "dark"),
			Head(
				Metatags(r),
				CSS(),
				JS(),
			),
			Body(
				Div(
					Class("hero flex items-center justify-center min-h-screen"),
					Div(
						Class("flex-col hero-content"),
						Div(
							Class("card shadow-md bg-base-200 w-96"),
							Div(
								Class("card-body"),
								If(len(r.Title) > 0, H1(Class("text-2xl font-bold"), Text(r.Title))),
								FlashMessages(r),
								content,
							),
						),
					),
				),
				HtmxListeners(r),
			),
		),
	)
}

func authNavBar(r *ui.Request) Node {
	return cache.SetIfNotExists("authNavBar", func() Node {
		return Ul(
			Class("menu menu-horizontal bg-base-200"),
			Li(A(Href(r.Path(routenames.Login)), Text("Login"))),
			Li(A(Href(r.Path(routenames.Register)), Text("Create an account"))),
			Li(A(Href(r.Path(routenames.ForgotPassword)), Text("Forgot password"))),
		)
	})
}
