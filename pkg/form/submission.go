package form

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mikestefanello/pagoda/pkg/context"

	"github.com/labstack/echo/v4"
)

// Submission represents the state of the submission of a form, not including the form itself.
// This satisfies the Form interface.
type Submission struct {
	// isSubmitted indicates if the form has been submitted
	isSubmitted bool

	// errors stores a slice of error message strings keyed by form struct field name
	errors map[string][]string
}

func (f *Submission) Submit(ctx echo.Context, form any) error {
	f.isSubmitted = true

	// Set in context so the form can later be retrieved
	ctx.Set(context.FormKey, form)

	// Bind the values from the incoming request to the form struct
	if err := ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to bind form: %v", err))
	}

	// Validate the form
	if err := ctx.Validate(form); err != nil {
		f.setErrorMessages(err)
		return err
	}

	return nil
}

func (f *Submission) IsSubmitted() bool {
	return f.isSubmitted
}

func (f *Submission) IsValid() bool {
	if f.errors == nil {
		return true
	}
	return len(f.errors) == 0
}

func (f *Submission) IsDone() bool {
	return f.IsSubmitted() && f.IsValid()
}

func (f *Submission) FieldHasErrors(fieldName string) bool {
	return len(f.GetFieldErrors(fieldName)) > 0
}

func (f *Submission) SetFieldError(fieldName string, message string) {
	if f.errors == nil {
		f.errors = make(map[string][]string)
	}
	f.errors[fieldName] = append(f.errors[fieldName], message)
}

func (f *Submission) GetFieldErrors(fieldName string) []string {
	if f.errors == nil {
		return []string{}
	}
	return f.errors[fieldName]
}

func (f *Submission) GetFieldStatusClass(fieldName string) string {
	if f.isSubmitted {
		if f.FieldHasErrors(fieldName) {
			return "is-danger"
		}
		return "is-success"
	}
	return ""
}

// setErrorMessages sets errors messages on the submission for all fields that failed validation
func (f *Submission) setErrorMessages(err error) {
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
		case "gte":
			message = fmt.Sprintf("Must be greater than or equal to %v.", ve.Param())
		default:
			message = "Invalid value."
		}

		// Add the error
		f.SetFieldError(ve.Field(), message)
	}
}
