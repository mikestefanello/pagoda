package controller

import (
	"reflect"

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

func (f FormSubmission) FieldHasError(fieldName string) bool {
	return len(f.GetFieldErrors(fieldName)) > 0
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

func (f *FormSubmission) setErrorMessages(form interface{}, err error) {
	// Only this is supported right now
	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		return
	}

	formType := reflect.TypeOf(form)

	for _, ve := range ves {
		var message string

		// Default the field form name to the name of the struct field
		fieldName := ve.StructField()

		// Attempt to get the form field name from the field's struct tag
		if field, ok := formType.FieldByName(ve.Field()); ok {
			if fieldNameTag := field.Tag.Get("form"); fieldNameTag != "" {
				fieldName = fieldNameTag
			}
		}

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
		f.Errors[fieldName] = append(f.Errors[fieldName], message)
	}
}
