package controllers

import (
	"goweb/auth"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type Logout struct {
	Controller
}

func (l *Logout) Get(c echo.Context) error {
	if err := auth.Logout(c); err == nil {
		msg.Success(c, "You have been logged out successfully.")
	}
	return l.Redirect(c, "home")
}
