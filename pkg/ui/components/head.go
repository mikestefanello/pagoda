package components

import (
	"fmt"
	"strings"

	"github.com/mikestefanello/pagoda/pkg/ui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func JS(r *ui.Request) Node {
	return Group{
		Script(Src("https://unpkg.com/htmx.org@2.0.0/dist/htmx.min.js")),
		Script(Src("https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"), Defer()),
	}
}

func HtmxListeners(r *ui.Request) Node {
	const htmxErr = `
		document.body.addEventListener('htmx:beforeSwap', function(evt) {
			if (evt.detail.xhr.status >= 400){
				evt.detail.shouldSwap = true;
				evt.detail.target = htmx.find("body");
			}
		});
	`

	const htmxCSRF = `
		document.body.addEventListener('htmx:configRequest', function(evt)  {
			if (evt.detail.verb !== "get") {
				evt.detail.parameters['csrf'] = '%s';
			}
		})
	`

	return Group{
		Script(Raw(htmxErr)),
		Iff(len(r.CSRF) > 0, func() Node {
			return Script(Raw(fmt.Sprintf(htmxCSRF, r.CSRF)))
		}),
	}
}

func CSS() Node {
	return Link(
		Href("https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css"),
		Rel("stylesheet"),
	)
}

func Metatags(r *ui.Request) Node {
	return Group{
		Meta(Charset("utf-8")),
		Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
		Link(Rel("icon"), Href(ui.File("favicon.png"))),
		TitleEl(Text(r.Config.App.Name), If(r.Title != "", Text(" | "+r.Title))),
		If(r.Metatags.Description != "", Meta(Name("description"), Content(r.Metatags.Description))),
		If(len(r.Metatags.Keywords) > 0, Meta(Name("keywords"), Content(strings.Join(r.Metatags.Keywords, ", ")))),
	}
}
