package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"goweb/config"
	"goweb/services"

	"github.com/labstack/echo/v4"
)

var (
	c *services.Container
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = services.NewContainer()

	// Run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}

func newContext(url string) echo.Context {
	req := httptest.NewRequest(http.MethodGet, url, strings.NewReader(""))
	return c.Web.NewContext(req, httptest.NewRecorder())
}
