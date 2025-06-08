package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type CardParams struct {
	Title  string
	Body   Group
	Footer Group
	Color  Color
	Size   Size
}

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

func Divider(text string) Node {
	return Div(
		Class("divider"),
		Text(text),
	)
}

func Card(params CardParams) Node {
	var colorClass, sizeClass string

	switch params.Color {
	case ColorSuccess:
		colorClass = "bg-success text-success-content"
	case ColorPrimary:
		colorClass = "bg-primary text-primary-content"
	case ColorAccent:
		colorClass = "bg-accent text-accent-content"
	case ColorNeutral:
		colorClass = "bg-neutral text-neutral-content"
	case ColorWarning:
		colorClass = "bg-warning text-warning-content"
	case ColorInfo:
		colorClass = "bg-info text-info-content"
	}

	switch params.Size {
	case SizeMedium:
		sizeClass = "card-md"
	case SizeLarge:
		sizeClass = "card-lg"
	}

	return Div(
		Class("cards mb-2 "+colorClass+" "+sizeClass),
		Div(
			Class("card-body"),
			If(len(params.Title) > 0, Span(
				Class("card-title"),
				Text(params.Title),
			)),
			params.Body,
			If(params.Footer != nil, Div(
				Class("card-actions justify-end"),
				params.Footer,
			)),
		),
	)
}
