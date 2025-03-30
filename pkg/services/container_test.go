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
	assert.NotNil(t, c.Files)
	assert.NotNil(t, c.ORM)
	assert.NotNil(t, c.Mail)
	assert.NotNil(t, c.Auth)
	assert.NotNil(t, c.Tasks)

	g := c.Graph
	if g == nil {
		//c.ORM.User.Create().
	}
}
