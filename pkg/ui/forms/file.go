package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type File struct{}

func (f File) Render(r *ui.Request) Node {
	return Form(
		ID("files"),
		Method(http.MethodPost),
		Action(r.Path(routenames.FilesSubmit)),
		EncType("multipart/form-data"),
		FileField("file", "Choose a file.. "),
		ControlGroup(
			FormButton("is-link", "Upload"),
		),
		CSRF(r),
	)
}
