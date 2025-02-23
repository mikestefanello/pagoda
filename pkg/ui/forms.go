package ui

import (
	"fmt"
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

	ResetPasswordForm struct {
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		form.Submission
	}

	TaskForm struct {
		Delay   int    `form:"delay" validate:"gte=0"`
		Message string `form:"message" validate:"required"`
		form.Submission
	}

	CacheForm struct {
		CurrentValue string
		Value        string `form:"value"`
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
		hxBoost(),
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
		hxBoost(),
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
		hxBoost(),
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

func (f *ResetPasswordForm) render(r *request) Node {
	return Form(
		ID("reset-password"),
		Method(http.MethodPost),
		hxBoost(),
		Action(r.Path),
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
			button("is-primary", "Update password"),
		),
		csrf(r),
	)
}

func (f *TaskForm) render(r *request) Node {
	return Form(
		ID("task"),
		Method(http.MethodPost),
		Attr("hx-post", r.path(routenames.TaskSubmit)),
		flashMessages(r),
		formInput(input{
			form:      f,
			formField: "Delay",
			name:      "delay",
			inputType: "number",
			label:     "Delay (in seconds)",
			help:      "How long to wait until the task is executed",
			value:     fmt.Sprint(f.Delay),
		}),
		formTextarea(textarea{
			form:      f,
			formField: "Message",
			name:      "message",
			label:     "Message",
			value:     f.Message,
			help:      "The message the task will output to the log",
		}),
		formControlGroup(
			button("is-link", "Add task to queue"),
		),
		csrf(r),
	)
}

func (f *CacheForm) render(r *request) Node {
	return Form(
		ID("cache"),
		Method(http.MethodPost),
		Attr("hx-post", r.path(routenames.CacheSubmit)),
		message(
			"is-info",
			"Test the cache",
			Group{
				P(Text("This route handler shows how the default in-memory cache works. Try updating the value using the form below and see how it persists after you reload the page.")),
				P(Text("HTMX makes it easy to re-render the cached value after the form is submitted.")),
			},
		),
		Label(
			For("value"),
			Class("value"),
			Text("Value in cache: "),
		),
		If(f.CurrentValue != "", Span(Class("tag is-success"), Text(f.CurrentValue))),
		If(f.CurrentValue == "", I(Text("(empty)"))),
		formInput(input{
			form:      f,
			formField: "Value",
			name:      "value",
			inputType: "text",
			label:     "Value",
			value:     f.Value,
		}),
		formControlGroup(
			button("is-link", "Update cache"),
		),
		csrf(r),
	)
}
