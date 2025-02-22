package ui

import (
	"github.com/mikestefanello/pagoda/pkg/routenames"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func layoutPrimary(r *request, content Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			head(r),
			Body(
				navBar(r),
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
							flashMessages(r),
							content,
						),
					),
				),
				footer(r),
			),
		),
	)
}

func layoutAuth(r *request, content Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			head(r),
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
										flashMessages(r),
										content,
										Div(
											Class("content is-small has-text-centered"),
											hxBoost(),
											A(Href(r.path(routenames.Login)), Text("Login")),
											Raw(" &#9676; "),
											A(Href(r.path("register")), Text("Create an account")),
											Raw(" &#9676; "),
											A(Href(r.path("forgot_password")), Text("Forgot password")),
										),
									),
								),
							),
						),
					),
				),
				footer(r),
			),
		),
	)
}
