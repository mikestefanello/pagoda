package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type ResetPassword struct {
	Password        string `form:"password" validate:"required"`
	ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
	form.Submission
}

func (f *ResetPassword) Render(r *ui.Request) Node {
	return Form(
		ID("reset-password"),
		Method(http.MethodPost),
		HxBoost(),
		Action(r.CurrentPath),
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
			FormField:   "PasswordConfirm",
			Name:        "password-confirm",
			InputType:   "password",
			Label:       "Confirm password",
			Placeholder: "******",
		}),
		ControlGroup(
			FormButton("btn-primary", "Update password"),
		),
		CSRF(r),
	)
}
