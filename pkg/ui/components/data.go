package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Badge(color Color, text string) Node {
	var class string

	switch color {
	case ColorSuccess:
		class = "badge-success"
	case ColorWarning:
		class = "badge-warning"
	}

	return Div(
		Class("badge "+class),
		Text(text),
	)
}
