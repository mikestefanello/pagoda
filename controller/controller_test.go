package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"goweb/config"
	"goweb/msg"
	"goweb/services"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
	if err := c.Shutdown(); err != nil {
		panic(err)
	}
	os.Exit(exitVal)
}

func newContext(url string) echo.Context {
	req := httptest.NewRequest(http.MethodGet, url, strings.NewReader(""))
	return c.Web.NewContext(req, httptest.NewRecorder())
}

func initSesssion(t *testing.T, ctx echo.Context) {
	// Simulate an HTTP request through the session middleware to initiate the session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	handler := mw(echo.NotFoundHandler)
	assert.Error(t, handler(ctx))
}

func TestController_Redirect(t *testing.T) {
	ctx := newContext("/abc")
	ctr := NewController(c)
	err := ctr.Redirect(ctx, "home")
	require.NoError(t, err)
	assert.Equal(t, "", ctx.Response().Header().Get(echo.HeaderLocation))
	assert.Equal(t, http.StatusFound, ctx.Response().Status)
}

func TestController_SetValidationErrorMessages(t *testing.T) {
	type example struct {
		Name string `validate:"required" label:"Label test"`
	}
	e := example{}
	v := validator.New()
	err := v.Struct(e)
	require.Error(t, err)

	ctx := newContext("/")
	initSesssion(t, ctx)
	ctr := NewController(c)
	ctr.SetValidationErrorMessages(ctx, err, e)

	msgs := msg.Get(ctx, msg.TypeDanger)
	require.Len(t, msgs, 1)
	assert.Equal(t, "<strong>Label test</strong> is required.", msgs[0])
}
