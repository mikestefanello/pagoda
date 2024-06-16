package redirect

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedirect(t *testing.T) {
	e := echo.New()
	e.GET("/path/:first/and/:second", func(c echo.Context) error {
		return nil
	}).Name = "test"

	redirect := func() (*Redirect, echo.Context) {
		ctx, _ := tests.NewContext(e, "/")
		return New(ctx), ctx
	}

	t.Run("route", func(t *testing.T) {
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		r, ctx := redirect()
		r.Route("test")
		r.Params("one", "two")
		r.Query(q)
		r.StatusCode(http.StatusTemporaryRedirect)
		require.NoError(t, r.Go())
		assert.Equal(t, "/path/one/and/two?a=1&b=2", ctx.Response().Header().Get(echo.HeaderLocation))
		assert.Equal(t, http.StatusTemporaryRedirect, ctx.Response().Status)
	})

	t.Run("route htmx", func(t *testing.T) {
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		r, ctx := redirect()
		ctx.Request().Header.Set(htmx.HeaderBoosted, "true")
		r.Route("test")
		r.Params("one", "two")
		r.Query(q)
		require.NoError(t, r.Go())
		assert.Equal(t, "/path/one/and/two?a=1&b=2", ctx.Response().Header().Get(htmx.HeaderRedirect))
	})

	t.Run("url", func(t *testing.T) {
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		r, ctx := redirect()
		r.URL("https://localhost.dev")
		r.Query(q)
		r.StatusCode(http.StatusTemporaryRedirect)
		require.NoError(t, r.Go())
		assert.Equal(t, "https://localhost.dev?a=1&b=2", ctx.Response().Header().Get(echo.HeaderLocation))
		assert.Equal(t, http.StatusTemporaryRedirect, ctx.Response().Status)
	})

	t.Run("url htmx", func(t *testing.T) {
		q := url.Values{}
		q.Add("a", "1")
		q.Add("b", "2")
		r, ctx := redirect()
		ctx.Request().Header.Set(htmx.HeaderBoosted, "true")
		r.URL("https://localhost.dev")
		r.Query(q)
		require.NoError(t, r.Go())
		assert.Equal(t, "https://localhost.dev?a=1&b=2", ctx.Response().Header().Get(htmx.HeaderRedirect))
	})
}
