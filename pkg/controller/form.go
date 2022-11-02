package controller

import (
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

// FormSubmission represents the state of the submission of a form, not including the form itself
type FormSubmission struct {
	// IsSubmitted indicates if the form has been submitted
	IsSubmitted bool

	// Errors stores a slice of error message strings keyed by form struct field name
	Errors map[string][]string
}

// Process processes a submission for a form
func (f *FormSubmission) Process(ctx echo.Context, form interface{}) error {
	f.Errors = make(map[string][]string)
	f.IsSubmitted = true

	// Validate the form
	if err := ctx.Validate(form); err != nil {
		f.setErrorMessages(err)
	}

	return nil
}

// HasErrors indicates if the submission has any validation errors
func (f FormSubmission) HasErrors() bool {
	if f.Errors == nil {
		return false
	}
	return len(f.Errors) > 0
}

// FieldHasErrors indicates if a given field on the form has any validation errors
func (f FormSubmission) FieldHasErrors(fieldName string) bool {
	return len(f.GetFieldErrors(fieldName)) > 0
}

// SetFieldError sets an error message for a given field name
func (f *FormSubmission) SetFieldError(fieldName string, message string) {
	if f.Errors == nil {
		f.Errors = make(map[string][]string)
	}
	f.Errors[fieldName] = append(f.Errors[fieldName], message)
}

// GetFieldErrors gets the errors for a given field name
func (f FormSubmission) GetFieldErrors(fieldName string) []string {
	if f.Errors == nil {
		return []string{}
	}
	return f.Errors[fieldName]
}

// GetFieldStatusClass returns an HTML class based on the status of the field
func (f FormSubmission) GetFieldStatusClass(fieldName string) string {
	if f.IsSubmitted {
		if f.FieldHasErrors(fieldName) {
			return "is-danger"
		}
		return "is-success"
	}
	return ""
}

// IsDone indicates if the submission is considered done which is when it has been submitted
// and there are no errors.
func (f FormSubmission) IsDone() bool {
	return f.IsSubmitted && !f.HasErrors()
}

// setErrorMessages sets errors messages on the submission for all fields that failed validation
func (f *FormSubmission) setErrorMessages(err error) {
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
