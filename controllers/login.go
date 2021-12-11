package controllers

import (
	"goweb/ent/user"
	"goweb/msg"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type Login struct {
	Controller
}

func (l *Login) Get(c echo.Context) error {
	p := NewPage(c)
	p.Layout = "auth"
	p.Name = "login"
	p.Title = "Log in"
	p.Data = "This is the login page"
	return l.RenderPage(c, p)
}

func (l *Login) Post(c echo.Context) error {
	name := c.FormValue("username")
	pw := c.FormValue("password")

	if name == "" || pw == "" {
		msg.Warning(c, "All fields are required.")
		return l.Get(c)
	}

	u, err := l.Container.Ent.User.
		Query().
		Where(user.Username(name)).
		First(c.Request().Context())

	if err != nil {
		c.Logger().Errorf("error querying user during login: %v", err)
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
		if err != nil {
			msg.Danger(c, "Invalid credentials. Please try again.")
		}
	}

	return l.Get(c)
}
