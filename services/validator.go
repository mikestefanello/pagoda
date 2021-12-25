package services

import (
	"github.com/go-playground/validator/v10"
)

// Validator provides validation mainly validating structs within the web context
type Validator struct {
	// validator stores the underlying validator
	validator *validator.Validate
}

// NewValidator creats a new Validator
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate validates a struct
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
