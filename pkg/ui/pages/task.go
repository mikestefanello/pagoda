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

func AddTask(ctx echo.Context, form *forms.Task) error {
	r := ui.NewRequest(ctx)
	r.Title = "Create a task"
	r.Metatags.Description = "Test creating a task to see how it works."

	g := Group{
		Iff(r.Htmx.Target != "task", func() Node {
			return components.Message(
				"is-link",
				"",
				Group{
					P(Raw("Submitting this form will create an <i>ExampleTask</i> in the task queue. After the specified delay, the message will be logged by the queue processor.")),
					P(Raw("See <i>pkg/tasks</i> and the README for more information.")),
				})
		}),
		form.Render(r),
		Iff(r.Htmx.Target != "task", func() Node {
			return components.Message(
				"is-warning",
				"",
				Group{
					If(!r.IsAdmin, P(Text("Log in as an admin in order to access the task and queue monitoring UI."))),
					If(r.IsAdmin, P(Text("View all queued tasks by clicking on the Tasks link in the sidebar."))),
				})
		}),
	}

	return r.Render(layouts.Primary, g)
}
