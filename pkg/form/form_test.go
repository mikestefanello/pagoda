package form

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockForm struct {
	called bool
	Submission
}

func (m *mockForm) Submit(_ echo.Context, _ any) error {
	m.called = true
	return nil
}

func TestSubmit(t *testing.T) {
	m := mockForm{}
	ctx, _ := tests.NewContext(echo.New(), "/")
	err := Submit(ctx, &m)
	require.NoError(t, err)
	assert.True(t, m.called)
}

func TestGetClear(t *testing.T) {
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

	t.Run("get non-empty context", func(t *testing.T) {
		form := example{
			Name: "test",
		}
		ctx, _ := tests.NewContext(e, "/")
		ctx.Set(context.FormKey, &form)

		// Get again and expect the values were stored
		got := Get[example](ctx)
		require.NotNil(t, got)
		assert.Equal(t, "test", form.Name)

		// Clear
		Clear(ctx)
		got = Get[example](ctx)
		require.NotNil(t, got)
		assert.Empty(t, got.Name)
	})
}
