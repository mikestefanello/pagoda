package ui

import (
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
										authNavBar(r),
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
