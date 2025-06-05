package components

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func MenuLink(r *ui.Request, icon Node, title, routeName string, routeParams ...any) Node {
	href := r.Path(routeName, routeParams...)

	return Li(
		Class("ml-2"),
		A(
			Href(href),
			icon,
			Text(title),
			Classes{
				"menu-active": href == r.CurrentPath,
				"p-2":         true,
			},
		),
	)
}

func Pager(page int, path string, hasNext bool) Node {
	href := func(page int) string {
		return fmt.Sprintf("%s?%s=%d",
			path,
			pager.QueryKey,
			page,
		)
	}

	return Div(
		Class("join"),
		A(
			Class("join-item btn"),
			Text("«"),
			If(page <= 1, Disabled()),
			Href(href(page-1)),
		),
		Button(
			Class("join-item btn"),
			Textf("Page %d", page),
		),
		A(
			Class("join-item btn"),
			Text("»"),
			If(!hasNext, Disabled()),
			Href(href(page+1)),
		),
	)
}
