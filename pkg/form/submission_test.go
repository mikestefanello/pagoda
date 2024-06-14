package form

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormSubmission(t *testing.T) {
	type formTest struct {
		Name  string `form:"name" validate:"required"`
		Email string `form:"email" validate:"required,email"`
		Submission
	}

	e := echo.New()
	e.Validator = services.NewValidator()

	t.Run("valid request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("email=a@a.com"))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		ctx := e.NewContext(req, httptest.NewRecorder())

		var form formTest
		err := form.Submit(ctx, &form)
		assert.IsType(t, validator.ValidationErrors{}, err)

		assert.Empty(t, form.Name)
		assert.Equal(t, "a@a.com", form.Email)
		assert.False(t, form.IsValid())
		assert.True(t, form.FieldHasErrors("Name"))
		assert.False(t, form.FieldHasErrors("Email"))
		require.Len(t, form.GetFieldErrors("Name"), 1)
		assert.Len(t, form.GetFieldErrors("Email"), 0)
		assert.Equal(t, "This field is required.", form.GetFieldErrors("Name")[0])
		assert.Equal(t, "is-danger", form.GetFieldStatusClass("Name"))
		assert.Equal(t, "is-success", form.GetFieldStatusClass("Email"))
		assert.False(t, form.IsDone())

		formInCtx := Get[formTest](ctx)
		require.NotNil(t, formInCtx)
		assert.Equal(t, form.Email, formInCtx.Email)
	})

	t.Run("invalid request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("abc=abc"))
		ctx := e.NewContext(req, httptest.NewRecorder())
		var form formTest
		err := form.Submit(ctx, &form)
		assert.IsType(t, new(echo.HTTPError), err)
	})
}
