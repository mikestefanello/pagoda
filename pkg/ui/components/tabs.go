package components

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Tab struct {
	Title, Body string
}

func Tabs(heading, description string, items []Tab) Node {
	renderTitles := func() Node {
		g := make(Group, len(items))
		for i, item := range items {
			g[i] = Li(
				Attr(":class", fmt.Sprintf("{'is-active': tab === %d}", i)),
				Attr("@click", fmt.Sprintf("tab = %d", i)),
				A(Text(item.Title)),
			)
		}
		return g
	}

	renderBodies := func() Node {
		g := make(Group, len(items))
		for i, item := range items {
			g[i] = Div(
				Attr("x-show", fmt.Sprintf("tab == %d", i)),
				P(Raw(" "+item.Body)),
			)
		}
		return g
	}

	return Div(
		P(
			Class("subtitle mt-5"),
			Text(heading),
		),
		P(
			Class("mb-4"),
			Text(description),
		),
		Div(
			Attr("x-data", "{tab: 0}"),
			Div(
				Class("tabs"),
				Ul(renderTitles()),
			),
			renderBodies(),
		),
	)
}
