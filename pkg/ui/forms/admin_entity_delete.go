package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AdminEntityDelete(r *ui.Request, entityTypeName string) Node {
	return Form(
		Method(http.MethodPost),
		P(
			Textf("Are you sure you want to delete this %s?", entityTypeName),
		),
		ControlGroup(
			FormButton("btn-error", "Delete"),
			ButtonLink(
				r.Path(routenames.AdminEntityList(entityTypeName)),
				"btn-link",
				"Cancel",
			),
		),
		CSRF(r),
	)
}
