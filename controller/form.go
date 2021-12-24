package controller

import (
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

type FormSubmission struct {
	IsSubmitted bool

	Errors map[string][]string
}

func (f *FormSubmission) Process(ctx echo.Context, form interface{}) error {
	f.Errors = make(map[string][]string)
	f.IsSubmitted = true

	// Validate the form
	if err := ctx.Validate(form); err != nil {
		f.setErrorMessages(form, err)
	}

	return nil
}

func (f FormSubmission) HasErrors() bool {
	if f.Errors == nil {
		return false
	}
	return len(f.Errors) > 0
}

func (f FormSubmission) FieldHasErrors(fieldName string) bool {
	return len(f.GetFieldErrors(fieldName)) > 0
}

func (f *FormSubmission) SetFieldError(fieldName string, message string) {
	if f.Errors == nil {
		f.Errors = make(map[string][]string)
	}
	f.Errors[fieldName] = append(f.Errors[fieldName], message)
}

func (f FormSubmission) GetFieldErrors(fieldName string) []string {
	if f.Errors == nil {
		return []string{}
	}

	errors, has := f.Errors[fieldName]
	if !has {
		return []string{}
	}

	return errors
}

func (f FormSubmission) GetFieldStatusClass(fieldName string) string {
	if f.IsSubmitted {
		if f.FieldHasErrors(fieldName) {
			return "is-danger"
		}
		return "is-success"
	}
	return ""
}

func (f FormSubmission) IsDone() bool {
	return f.IsSubmitted && !f.HasErrors()
}

func (f *FormSubmission) setErrorMessages(form interface{}, err error) {
	// Only this is supported right now
	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		return
	}

	for _, ve := range ves {
		var message string

		// Provide better error messages depending on the failed validation tag
		// This should be expanded as you use additional tags in your validation
		switch ve.Tag() {
		case "required":
			message = "This field is required."
		case "email":
			message = "Enter a valid email address."
		case "eqfield":
			message = "Does not match."
		default:
			message = "Invalid value."
		}

		// Add the error
		f.SetFieldError(ve.Field(), message)
	}
}
