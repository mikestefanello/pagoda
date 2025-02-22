package ui

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type (
	ContactForm struct {
		Email      string `form:"email" validate:"required,email"`
		Department string `form:"department" validate:"required,oneof=sales marketing hr"`
		Message    string `form:"message" validate:"required"`
		form.Submission
	}

	LoginForm struct {
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required"`
		form.Submission
	}

	RegisterForm struct {
		Name            string `form:"name" validate:"required"`
		Email           string `form:"email" validate:"required,email"`
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		form.Submission
	}

	ForgotPasswordForm struct {
		Email string `form:"email" validate:"required,email"`
		form.Submission
	}
)

func (f *ContactForm) render(r *request) Node {
	return Form(
		ID("contact"),
		Method(http.MethodPost),
		Attr("hx-post", r.path(routenames.ContactSubmit)),
		formInput(input{
			form:      f,
			formField: "Email",
			name:      "email",
			inputType: "email",
			label:     "Email address",
			value:     f.Email,
		}),
		formRadios(radios{
			form:      f,
			formField: "Department",
			name:      "department",
			label:     "Department",
			value:     f.Department,
			options: []radio{
				{value: "sales", label: "Sales"},
				{value: "marketing", label: "Marketing"},
				{value: "hr", label: "HR"},
			},
		}),
		formTextarea(textarea{
			form:      f,
			formField: "Message",
			name:      "message",
			label:     "Message",
			value:     f.Message,
		}),
		formControlGroup(
			button("is-link", "Submit"),
		),
		csrf(r),
	)
}

func (f *LoginForm) render(r *request) Node {
	return Form(
		ID("login"),
		Method(http.MethodPost),
		Attr("hx-boost", "true"),
		Action(r.path(routenames.LoginSubmit)),
		flashMessages(r),
		formInput(input{
			form:      f,
			formField: "Email",
			name:      "email",
			inputType: "email",
			label:     "Email address",
			value:     f.Email,
		}),
		formInput(input{
			form:        f,
			formField:   "Password",
			name:        "password",
			inputType:   "password",
			label:       "Password",
			placeholder: "******",
		}),
		formControlGroup(
			button("is-link", "Login"),
			buttonLink(r.path(routenames.Home), "is-light", "Cancel"),
		),
		csrf(r),
	)
}

func (f *RegisterForm) render(r *request) Node {
	return Form(
		ID("register"),
		Method(http.MethodPost),
		Attr("hx-boost", "true"),
		Action(r.path(routenames.RegisterSubmit)),
		formInput(input{
			form:      f,
			formField: "Name",
			name:      "name",
			inputType: "text",
			label:     "Name",
			value:     f.Name,
		}),
		formInput(input{
			form:      f,
			formField: "Email",
			name:      "email",
			inputType: "email",
			label:     "Email address",
			value:     f.Email,
		}),
		formInput(input{
			form:        f,
			formField:   "Password",
			name:        "password",
			inputType:   "password",
			label:       "Password",
			placeholder: "******",
		}),
		formInput(input{
			form:        f,
			formField:   "PasswordConfirm",
			name:        "password-confirm",
			inputType:   "password",
			label:       "Confirm password",
			placeholder: "******",
		}),
		formControlGroup(
			button("is-primary", "Register"),
			buttonLink(r.path(routenames.Home), "is-light", "Cancel"),
		),
		csrf(r),
	)
}

func (f *ForgotPasswordForm) render(r *request) Node {
	return Form(
		ID("forgot-password"),
		Method(http.MethodPost),
		Attr("hx-boost", "true"),
		Action(r.path(routenames.ForgotPasswordSubmit)),
		formInput(input{
			form:      f,
			formField: "Email",
			name:      "email",
			inputType: "email",
			label:     "Email address",
			value:     f.Email,
		}),
		formControlGroup(
			button("is-primary", "Reset password"),
			buttonLink(r.path(routenames.Home), "is-light", "Cancel"),
		),
		csrf(r),
	)
}
