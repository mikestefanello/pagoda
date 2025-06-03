package components

import (
	"fmt"
	"math/rand"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Tab struct {
	Title, Body string
}

func Tabs(tabs []Tab) Node {
	g := make(Group, 0, len(tabs)*2)
	id := fmt.Sprintf("tabs-%d", rand.Int())

	for i, tab := range tabs {
		g = append(g,
			Input(
				Type("radio"),
				Name(id),
				Class("tab"),
				Aria("label", tab.Title),
				If(i == 0, Checked()),
			),
			Div(
				Class("tab-content bg-base-100 border-base-300 p-6"),
				Raw(tab.Body),
			))
	}

	return Div(
		Class("tabs tabs-lift"),
		g,
	)
}
