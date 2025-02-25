package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	"github.com/mikestefanello/pagoda/pkg/ui/models"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func UploadFile(ctx echo.Context, files []*models.File) error {
	r := ui.NewRequest(ctx)
	r.Title = "Upload a file"

	fileList := make(Group, 0, len(files))
	for _, file := range files {
		fileList = append(fileList, file.Render())
	}

	n := Group{
		components.Message(
			"is-link",
			"",
			P(Text("This is a very basic example of how to handle file uploads. Files uploaded will be saved to the directory specified in your configuration.")),
		),
		Hr(),
		forms.File{}.Render(r),
		Hr(),
		H3(
			Class("title"),
			Text("Uploaded files"),
		),
		components.Message("is-warning", "", P(Text("Below are all files in the configured upload directory."))),
		Table(
			Class("table"),
			THead(
				Tr(
					Th(Text("Filename")),
					Th(Text("Size")),
					Th(Text("Modified on")),
				),
			),
			TBody(
				fileList,
			),
		),
	}

	return r.Render(layouts.Primary, n)
}
