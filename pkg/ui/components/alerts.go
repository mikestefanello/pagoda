package components

import (
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/icons"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func FlashMessages(r *ui.Request) Node {
	var g Group
	var color Color

	for _, typ := range []msg.Type{
		msg.TypeSuccess,
		msg.TypeInfo,
		msg.TypeWarning,
		msg.TypeError,
	} {
		for _, str := range msg.Get(r.Context, typ) {
			switch typ {
			case msg.TypeSuccess:
				color = ColorSuccess
			case msg.TypeInfo:
				color = ColorInfo
			case msg.TypeWarning:
				color = ColorWarning
			case msg.TypeError:
				color = ColorError
			}

			g = append(g, Alert(color, str))
		}
	}

	return g
}

func Alert(color Color, text string) Node {
	var class string

	switch color {
	case ColorSuccess:
		class = "alert-success"
	case ColorInfo:
		class = "alert-info"
	case ColorWarning:
		class = "alert-warning"
	case ColorError:
		class = "alert-error"
	}

	return Div(
		Role("alert"),
		Class("alert mb-2 "+class),
		Attr("x-data", "{show: true}"),
		Attr("x-show", "show"),
		Span(
			Attr("@click", "show = false"),
			Class("cursor-pointer"),
			icons.XCircle(),
		),
		Span(Text(text)),
	)
}
