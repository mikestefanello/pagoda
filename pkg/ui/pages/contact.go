package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ContactUs(ctx echo.Context, form *forms.Contact) error {
	r := ui.NewRequest(ctx)
	r.Title = "Contact us"
	r.Metatags.Description = "Get in touch with us."

	g := make(Group, 0)

	if r.Htmx.Target != "contact" {
		g = append(g, components.Message(
			"is-link",
			"",
			Group{
				P(Text("This is an example of a form with inline, server-side validation and HTMX-powered AJAX submissions without writing a single line of JavaScript.")),
				P(Text("Only the form below will update async upon submission.")),
			}))
	}

	if form.IsDone() {
		g = append(g, components.Message(
			"is-large is-success",
			"Thank you!",
			Text("No email was actually sent but this entire operation was handled server-side and degrades without JavaScript enabled."),
		))
	} else {
		g = append(g, form.Render(r))
	}

	return r.Render(layouts.Primary, g)
}
