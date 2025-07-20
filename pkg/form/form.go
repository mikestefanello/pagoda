package form

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
)

// Form represents a form that can be submitted and validated.
type Form interface {
	// Submit marks the form as submitted, stores a pointer to it in the context, binds the request
	// values to the struct fields, and validates the input based on the struct tags.
	// Returns a validator.ValidationErrors, if the form values were not valid, or an echo.HTTPError,
	// if the request failed to process.
	Submit(c echo.Context, form any) error

	// IsSubmitted returns true if the form was submitted.
	IsSubmitted() bool

	// IsValid returns true if the form has no validation errors.
	IsValid() bool

	// IsDone returns true if the form was submitted and has no validation errors.
	IsDone() bool

	// FieldHasErrors returns true if a given struct field has validation errors.
	FieldHasErrors(fieldName string) bool

	// SetFieldError sets a validation error message for a given struct field.
	SetFieldError(fieldName string, message string)

	// GetFieldErrors returns the validation errors for a given struct field.
	GetFieldErrors(fieldName string) []string
}

// Get gets a form from the context or initializes a new copy if one is not set.
func Get[T any](ctx echo.Context) *T {
	if v := ctx.Get(context.FormKey); v != nil {
		if form, ok := v.(*T); ok {
			return form
		}
	}
	var v T
	return &v
}

// Clear removes the form set in the context.
func Clear(ctx echo.Context) {
	ctx.Set(context.FormKey, nil)
}

// Submit submits a form.
// See Form.Submit().
func Submit(ctx echo.Context, form Form) error {
	return form.Submit(ctx, form)
}
