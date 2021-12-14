package routes

import (
	"net/http"

	"goweb/controller"

	"github.com/labstack/echo/v4"
)

type Error struct {
	controller.Controller
}

func (e *Error) Get(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code >= 500 {
		c.Logger().Error(err)
	} else {
		c.Logger().Info(err)
	}

	p := controller.NewPage(c)
	p.Layout = "main"
	p.Title = http.StatusText(code)
	p.Name = "error"
	p.StatusCode = code

	if err = e.RenderPage(c, p); err != nil {
		c.Logger().Error(err)
	}
}
