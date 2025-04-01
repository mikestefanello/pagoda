package components

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/ui"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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

func HxBoost() Node {
	return Attr("hx-boost", "true")
}
