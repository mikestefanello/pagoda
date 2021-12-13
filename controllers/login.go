package controllers

import (
	"fmt"

	"goweb/auth"
	"goweb/ent"
	"goweb/ent/user"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type (
	Login struct {
		Controller
		form LoginForm
	}

	LoginForm struct {
		Username string `form:"username" validate:"required"`
		Password string `form:"password" validate:"required"`
	}
)

func (l *Login) Get(c echo.Context) error {
	p := NewPage(c)
	p.Layout = "auth"
	p.Name = "login"
	p.Title = "Log in"
	p.Data = l.form
	return l.RenderPage(c, p)
}

func (l *Login) Post(c echo.Context) error {
	// Parse the form values
	if err := c.Bind(&l.form); err != nil {
		c.Logger().Errorf("unable to parse login form: %v", err)
		msg.Danger(c, "An error occurred. Please try again.")
		return l.Get(c)
	}

	// Validate the form
	if err := c.Validate(l.form); err != nil {
		msg.Danger(c, "All fields are required.")
		return l.Get(c)
	}

	// Attempt to load the user
	u, err := l.Container.ORM.User.
		Query().
		Where(user.Username(l.form.Username)).
		First(c.Request().Context())

	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			msg.Danger(c, "Invalid credentials. Please try again.")
			return l.Get(c)
		default:
			c.Logger().Errorf("error querying user during login: %v", err)
			msg.Danger(c, "An error occurred. Please try again.")
			return l.Get(c)
		}

	}

	// Check if the password is correct
	err = auth.CheckPassword(l.form.Password, u.Password)
	if err != nil {
		msg.Danger(c, "Invalid credentials. Please try again.")
		return l.Get(c)
	}

	// Log the user in
	err = auth.Login(c, u.ID)
	if err != nil {
		c.Logger().Errorf("unable to log in user %d: %v", u.ID, err)
		msg.Danger(c, "An error occurred. Please try again.")
		return l.Get(c)
	}

	msg.Success(c, fmt.Sprintf("Welcome back, %s. You are now logged in.", u.Username))
	return l.Redirect(c, "home")
}
