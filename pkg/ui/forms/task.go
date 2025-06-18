package forms

import (
	"fmt"
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Task struct {
	Delay   int    `form:"delay" validate:"gte=0"`
	Message string `form:"message" validate:"required"`
	form.Submission
}

func (f *Task) Render(r *ui.Request) Node {
	return Form(
		ID("task"),
		Method(http.MethodPost),
		Attr("hx-post", r.Path(routenames.TaskSubmit)),
		FlashMessages(r),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Delay",
			Name:      "delay",
			InputType: "number",
			Label:     "Delay (in seconds)",
			Help:      "How long to wait until the task is executed",
			Value:     fmt.Sprint(f.Delay),
		}),
		TextareaField(TextareaFieldParams{
			Form:      f,
			FormField: "Message",
			Name:      "message",
			Label:     "Message",
			Value:     f.Message,
			Help:      "The message the task will output to the log",
		}),
		ControlGroup(
			FormButton(ColorPrimary, "Add task to queue"),
		),
		CSRF(r),
	)
}
