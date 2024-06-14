package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/templates"
)

type Error struct {
	controller.Controller
}

func (e *Error) Page(err error, ctx echo.Context) {
	if ctx.Response().Committed || context.IsCanceledError(err) {
		return
	}

	// Determine the error status code
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	// Log the error
	msg := "request failed"
	sub := log.Ctx(ctx).With("error", err)
	if code >= 500 {
		sub.Error(msg)
	} else {
		sub.Warn(msg)
	}

	// Render the error page
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutMain
	page.Name = templates.PageError
	page.Title = http.StatusText(code)
	page.StatusCode = code
	page.HTMX.Request.Enabled = false

	if err = e.RenderPage(ctx, page); err != nil {
		log.Ctx(ctx).Error("failed to render error page", "error", err)
	}
}
