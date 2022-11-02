package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	type example struct {
		Value string `validate:"required"`
	}
	e := example{}
	err := c.Validator.Validate(e)
	assert.Error(t, err)
	e.Value = "a"
	err = c.Validator.Validate(e)
	assert.NoError(t, err)
}
