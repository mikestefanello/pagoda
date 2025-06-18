package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ContactUs(ctx echo.Context, form *forms.Contact) error {
	r := ui.NewRequest(ctx)
	r.Title = "Contact us"
	r.Metatags.Description = "Get in touch with us."

	g := Group{
		Iff(r.Htmx.Target != "contact", func() Node {
			return Card(CardParams{
				Title: "Card component",
				Body: Group{
					Span(Text("This is an example of a form with inline, server-side validation and HTMX-powered AJAX submissions without writing a single line of JavaScript.")),
					Span(Text("Only the form below will update async upon submission.")),
				},
				Color: ColorWarning,
				Size:  SizeMedium,
			})
		}),
		Iff(form.IsDone(), func() Node {
			return Card(CardParams{
				Title: "Thank you!",
				Body: Group{
					Span(Text("No email was actually sent but this entire operation was handled server-side and degrades without JavaScript enabled.")),
				},
				Color: ColorSuccess,
				Size:  SizeLarge,
			})
		}),
		Iff(!form.IsDone(), func() Node {
			return form.Render(r)
		}),
	}

	return r.Render(layouts.Primary, g)
}
