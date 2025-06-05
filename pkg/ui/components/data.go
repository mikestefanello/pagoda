package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Badge(class, text string) Node {
	return Div(
		Class("badge "+class),
		Text(text),
	)
}
