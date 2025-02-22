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
		formSubmit("Submit"),
		csrf(r),
	)
}
