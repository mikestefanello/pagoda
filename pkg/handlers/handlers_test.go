package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSetHandlers(t *testing.T) {
	handlers = []Handler{}
	assert.Empty(t, GetHandlers())
	h := new(Pages)
	Register(h)
	got := GetHandlers()
	require.Len(t, got, 1)
	assert.Equal(t, h, got[0])
}

func TestRedirect(t *testing.T) {
	c.Web.GET("/path/:first/and/:second", func(c echo.Context) error {
		return nil
	}).Name = "redirect-test"

	t.Run("no query", func(t *testing.T) {
		ctx, _ := tests.NewContext(c.Web, "/abc")
		err := redirect(ctx, "redirect-test", "one", "two")
		require.NoError(t, err)
		assert.Equal(t, "/path/one/and/two", ctx.Response().Header().Get(echo.HeaderLocation))
		assert.Equal(t, http.StatusFound, ctx.Response().Status)
	})

	t.Run("no query htmx", func(t *testing.T) {
		ctx, _ := tests.NewContext(c.Web, "/abc")
		ctx.Request().Header.Set(htmx.HeaderBoosted, "true")
		err := redirect(ctx, "redirect-test", "one", "two")
		require.NoError(t, err)
		assert.Equal(t, "/path/one/and/two", ctx.Response().Header().Get(htmx.HeaderRedirect))
	})

	t.Run("query", func(t *testing.T) {
		ctx, _ := tests.NewContext(c.Web, "/abc")
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		err := redirectWithQuery(ctx, q, "redirect-test", "one", "two")
		require.NoError(t, err)
		assert.Equal(t, "/path/one/and/two?a=1&b=2", ctx.Response().Header().Get(echo.HeaderLocation))
		assert.Equal(t, http.StatusFound, ctx.Response().Status)
	})

	t.Run("query htmx", func(t *testing.T) {
		ctx, _ := tests.NewContext(c.Web, "/abc")
		ctx.Request().Header.Set(htmx.HeaderBoosted, "true")
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		err := redirectWithQuery(ctx, q, "redirect-test", "one", "two")
		require.NoError(t, err)
		assert.Equal(t, "/path/one/and/two?a=1&b=2", ctx.Response().Header().Get(htmx.HeaderRedirect))
		assert.Equal(t, http.StatusFound, ctx.Response().Status)
	})
}

func TestFail(t *testing.T) {
	err := fail(errors.New("err message"), "log message")
	require.IsType(t, new(echo.HTTPError), err)
	he := err.(*echo.HTTPError)
	assert.Equal(t, http.StatusInternalServerError, he.Code)
	assert.Equal(t, "log message: err message", he.Message)
}
