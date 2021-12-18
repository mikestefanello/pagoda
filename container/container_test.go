package container

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"goweb/config"
	"goweb/ent"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"
)

var (
	c   *Container
	ctx echo.Context
	usr *ent.User
	rec *httptest.ResponseRecorder
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = NewContainer()

	// Create a web context
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	rec = httptest.NewRecorder()
	ctx = c.Web.NewContext(req, rec)

	// Create a test user
	var err error
	usr, err = c.ORM.User.
		Create().
		SetEmail("test@test.dev").
		SetPassword("abc").
		SetName("Test User").
		Save(context.Background())

	if err != nil {
		panic(err)
	}

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
