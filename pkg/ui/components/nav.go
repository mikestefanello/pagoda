package components

import (
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
