package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	assert.NotNil(t, c.Web)
	assert.NotNil(t, c.Config)
	assert.NotNil(t, c.Validator)
	assert.NotNil(t, c.Cache)
	assert.NotNil(t, c.Database)
	assert.NotNil(t, c.ORM)
	assert.NotNil(t, c.Mail)
	assert.NotNil(t, c.Auth)
	assert.NotNil(t, c.TemplateRenderer)
	assert.NotNil(t, c.Tasks)
}
