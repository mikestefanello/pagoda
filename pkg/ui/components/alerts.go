package components

import (
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func FlashMessages(r *ui.Request) Node {
	var g Group
	for _, typ := range []msg.Type{
		msg.TypeSuccess,
		msg.TypeInfo,
		msg.TypeWarning,
		msg.TypeDanger,
	} {
		for _, str := range msg.Get(r.Context, typ) {
			g = append(g, Notification(typ, str))
		}
	}

	return g
}

func Notification(typ msg.Type, text string) Node {
	var class string

	switch typ {
	case msg.TypeSuccess:
		class = "success"
	case msg.TypeInfo:
		class = "info"
	case msg.TypeWarning:
		class = "warning"
	case msg.TypeDanger:
		class = "danger"
	}

	return Div(
		Class("notification is-"+class),
		Attr("x-data", "{show: true}"),
		Attr("x-show", "show"),
		Button(
			Class("delete"),
			Attr("@click", "show = false"),
		),
		Text(text),
	)
}

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
