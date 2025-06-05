package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Register struct {
	Name            string `form:"name" validate:"required"`
	Email           string `form:"email" validate:"required,email"`
	Password        string `form:"password" validate:"required"`
	ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
	form.Submission
}

func (f *Register) Render(r *ui.Request) Node {
	return Form(
		ID("register"),
		Method(http.MethodPost),
		HxBoost(),
		Action(r.Path(routenames.RegisterSubmit)),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Name",
			Name:      "name",
			InputType: "text",
			Label:     "Name",
			Value:     f.Name,
		}),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Email",
			Name:      "email",
			InputType: "email",
			Label:     "Email address",
			Value:     f.Email,
		}),
		InputField(InputFieldParams{
			Form:        f,
			FormField:   "Password",
			Name:        "password",
			InputType:   "password",
			Label:       "Password",
			Placeholder: "******",
		}),
		InputField(InputFieldParams{
			Form:        f,
			FormField:   "ConfirmPassword",
			Name:        "password-confirm",
			InputType:   "password",
			Label:       "Confirm password",
			Placeholder: "******",
		}),
		ControlGroup(
			FormButton("btn-primary", "Register"),
			ButtonLink(r.Path(routenames.Home), "btn-link", "Cancel"),
		),
		CSRF(r),
		Div(
			Class("text-center text-base-content/50 mt-4"),
			Text("Already have an account? "),
			A(
				Href(r.Path(routenames.Login)),
				Text("Login"),
			),
		),
	)
}
