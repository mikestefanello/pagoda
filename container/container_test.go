package container

import (
	"os"
	"testing"

	"goweb/config"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestNewContainer(t *testing.T) {
	c := NewContainer()
	assert.NotNil(t, c.Web)
	assert.NotNil(t, c.Config)
	assert.NotNil(t, c.Cache)
	assert.NotNil(t, c.Database)
	assert.NotNil(t, c.ORM)
	assert.NotNil(t, c.Mail)
	assert.NotNil(t, c.Auth)
}
