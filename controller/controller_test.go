package controller

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"goweb/config"
	"goweb/middleware"
	"goweb/msg"
	"goweb/services"

	"github.com/eko/gocache/v2/store"

	"github.com/eko/gocache/v2/marshaler"

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
	defer func() {
		if err := c.Shutdown(); err != nil {
			c.Web.Logger.Fatal(err)
		}
	}()

	// Run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}

func newContext(url string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, url, strings.NewReader(""))
	rec := httptest.NewRecorder()
	return c.Web.NewContext(req, rec), rec
}

func initSesssion(t *testing.T, ctx echo.Context) {
	// Simulate an HTTP request through the session middleware to initiate the session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	handler := mw(echo.NotFoundHandler)
	assert.Error(t, handler(ctx))
}

func TestController_Redirect(t *testing.T) {
	ctx, _ := newContext("/abc")
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

	ctx, _ := newContext("/")
	initSesssion(t, ctx)
	ctr := NewController(c)
	ctr.SetValidationErrorMessages(ctx, err, e)

	msgs := msg.Get(ctx, msg.TypeDanger)
	require.Len(t, msgs, 1)
	assert.Equal(t, "<strong>Label test</strong> is required.", msgs[0])
}

func TestController_RenderPage(t *testing.T) {
	setup := func() (echo.Context, *httptest.ResponseRecorder, Controller, Page) {
		ctx, rec := newContext("/test/TestController_RenderPage")
		initSesssion(t, ctx)
		ctr := NewController(c)

		p := NewPage(ctx)
		p.Name = "home"
		p.Layout = "main"
		p.Cache.Enabled = false
		p.Headers["a"] = "b"
		p.Headers["c"] = "d"
		p.StatusCode = http.StatusCreated
		return ctx, rec, ctr, p
	}

	t.Run("missing name", func(t *testing.T) {
		// Rendering should fail if the Page has no name
		ctx, _, ctr, p := setup()
		p.Name = ""
		err := ctr.RenderPage(ctx, p)
		assert.Error(t, err)
	})

	t.Run("no page cache", func(t *testing.T) {
		ctx, _, ctr, p := setup()
		err := ctr.RenderPage(ctx, p)
		require.NoError(t, err)

		// Check status code and headers
		assert.Equal(t, http.StatusCreated, ctx.Response().Status)
		for k, v := range p.Headers {
			assert.Equal(t, v, ctx.Response().Header().Get(k))
		}

		// Check the template cache
		parsed, err := c.TemplateRenderer.Load("controller", p.Name)
		assert.NoError(t, err)

		// Check that all expected templates were parsed.
		// This includes the name, layout and all components
		expectedTemplates := make(map[string]bool)
		expectedTemplates[p.Name+config.TemplateExt] = true
		expectedTemplates[p.Layout+config.TemplateExt] = true
		components, err := ioutil.ReadDir(c.TemplateRenderer.GetTemplatesPath() + "/components")
		require.NoError(t, err)
		for _, f := range components {
			expectedTemplates[f.Name()] = true
		}

		for _, v := range parsed.Templates() {
			delete(expectedTemplates, v.Name())
		}
		assert.Empty(t, expectedTemplates)
	})

	t.Run("page cache", func(t *testing.T) {
		ctx, rec, ctr, p := setup()
		p.Cache.Enabled = true
		p.Cache.Tags = []string{"tag1"}
		err := ctr.RenderPage(ctx, p)
		require.NoError(t, err)

		// Fetch from the cache
		res, err := marshaler.New(c.Cache).
			Get(context.Background(), p.URL, new(middleware.CachedPage))
		require.NoError(t, err)

		// Compare the cached page
		cp, ok := res.(*middleware.CachedPage)
		require.True(t, ok)
		assert.Equal(t, p.URL, cp.URL)
		assert.Equal(t, p.Headers, cp.Headers)
		assert.Equal(t, p.StatusCode, cp.StatusCode)
		assert.Equal(t, rec.Body.Bytes(), cp.HTML)

		// Clear the tag
		err = c.Cache.Invalidate(context.Background(), store.InvalidateOptions{
			Tags: []string{p.Cache.Tags[0]},
		})
		require.NoError(t, err)

		// Refetch from the cache and expect no results
		_, err = marshaler.New(c.Cache).
			Get(context.Background(), p.URL, new(middleware.CachedPage))
		assert.Error(t, err)
	})
}
