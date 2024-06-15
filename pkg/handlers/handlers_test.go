package handlers

import (
	"errors"
	"net/http"
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

	t.Run("normal", func(t *testing.T) {
		ctx, _ := tests.NewContext(c.Web, "/abc")
		err := redirect(ctx, "redirect-test", "one", "two")
		require.NoError(t, err)
		assert.Equal(t, "/path/one/and/two", ctx.Response().Header().Get(echo.HeaderLocation))
		assert.Equal(t, http.StatusFound, ctx.Response().Status)
	})

	t.Run("htmx boosted", func(t *testing.T) {
		ctx, _ := tests.NewContext(c.Web, "/abc")
		ctx.Request().Header.Set(htmx.HeaderBoosted, "true")
		err := redirect(ctx, "redirect-test", "one", "two")
		require.NoError(t, err)
		assert.Equal(t, "/path/one/and/two", ctx.Response().Header().Get(htmx.HeaderRedirect))
	})
}

func TestFail(t *testing.T) {
	err := fail(errors.New("err message"), "log message")
	require.IsType(t, new(echo.HTTPError), err)
	he := err.(*echo.HTTPError)
	assert.Equal(t, http.StatusInternalServerError, he.Code)
	assert.Equal(t, "log message: err message", he.Message)
}
