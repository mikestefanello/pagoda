package form

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContextFuncs(t *testing.T) {
	e := echo.New()

	type example struct {
		Name string `form:"name"`
	}

	t.Run("get empty context", func(t *testing.T) {
		// Empty context, still return a form
		ctx, _ := tests.NewContext(e, "/")
		form := Get[example](ctx)
		assert.NotNil(t, form)
	})

	t.Run("set bad request", func(t *testing.T) {
		// Set with a bad request
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("abc=abc"))
		ctx := e.NewContext(req, httptest.NewRecorder())
		var form example
		err := Set(ctx, &form)
		assert.Error(t, err)
	})

	t.Run("set", func(t *testing.T) {
		// Set and parse the values
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("name=abc"))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		ctx := e.NewContext(req, httptest.NewRecorder())
		var form example
		err := Set(ctx, &form)
		require.NoError(t, err)
		assert.Equal(t, "abc", form.Name)

		// Get again and expect the values were stored
		got := Get[example](ctx)
		require.NotNil(t, got)
		assert.Equal(t, "abc", form.Name)

		// Clear
		Clear(ctx)
		got = Get[example](ctx)
		require.NotNil(t, got)
		assert.Empty(t, got.Name)
	})
}
