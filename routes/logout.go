package routes

import (
	"goweb/controller"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type Logout struct {
	controller.Controller
}

func (l *Logout) Get(c echo.Context) error {
	if err := l.Container.Auth.Logout(c); err == nil {
		msg.Success(c, "You have been logged out successfully.")
	}
	return l.Redirect(c, "home")
}
