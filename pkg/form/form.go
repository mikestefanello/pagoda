package form

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
)

// Get gets a form from the context or initializes a new copy if one is not set
func Get[T any](ctx echo.Context) *T {
	if v := ctx.Get(context.FormKey); v != nil {
		return v.(*T)
	}
	var v T
	return &v
}

// Set sets a form in the context and binds the request values to it
func Set(ctx echo.Context, form any) error {
	ctx.Set(context.FormKey, form)

	if err := ctx.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("unable to bind form: %v", err))
	}

	return nil
}

// Clear removes the form set in the context
func Clear(ctx echo.Context) {
	ctx.Set(context.FormKey, nil)
}
