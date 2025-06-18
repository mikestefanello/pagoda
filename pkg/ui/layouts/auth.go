package layouts

import (
	"github.com/mikestefanello/pagoda/pkg/ui"
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
