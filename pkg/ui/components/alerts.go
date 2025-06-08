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
	for _, typ := range []msg.Type{
		msg.TypeSuccess,
		msg.TypeInfo,
		msg.TypeWarning,
		msg.TypeError,
	} {
		for _, str := range msg.Get(r.Context, typ) {
			g = append(g, Alert(typ, str))
		}
	}

	return g
}

func Alert(typ msg.Type, text string) Node {
	var class string

	switch typ {
	case msg.TypeSuccess:
		class = "alert-success"
	case msg.TypeInfo:
		class = "alert-info"
	case msg.TypeWarning:
		class = "alert-warning"
	case msg.TypeError:
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

// Deprecated
func Message(class, header string, body Node) Node {
	return Article(
		Class("message "+class),
		If(header != "", Div(
			Class("message-header"),
			P(Text(header)),
		)),
		Div(
			Class("message-body"),
			body,
		),
	)
}
