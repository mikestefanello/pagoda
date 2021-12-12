package controllers

import (
	"goweb/auth"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type (
	Register struct {
		Controller
		form RegisterForm
	}

	RegisterForm struct {
		Username string `form:"username" validate:"required"`
		Password string `form:"password" validate:"required"`
	}
)

func (r *Register) Get(c echo.Context) error {
	p := NewPage(c)
	p.Layout = "auth"
	p.Name = "register"
	p.Title = "Register"
	p.Data = r.form

	return r.RenderPage(c, p)
}

func (r *Register) Post(c echo.Context) error {
	fail := func(message string, err error) error {
		c.Logger().Errorf("%s: %v", message, err)
		msg.Danger(c, "An error occurred. Please try again.")
		return r.Get(c)
	}

	// Parse the form values
	if err := c.Bind(&r.form); err != nil {
		return fail("unable to parse form values", err)
	}

	// Validate the form
	if err := c.Validate(r.form); err != nil {
		msg.Danger(c, "All fields are required.")
		return r.Get(c)
	}

	// Hash the password
	pwHash, err := auth.HashPassword(r.form.Password)
	if err != nil {
		return fail("unable to hash password", err)
	}

	// Attempt creating the user
	u, err := r.Container.ORM.User.
		Create().
		SetUsername(r.form.Username).
		SetPassword(pwHash).
		Save(c.Request().Context())

	if err != nil {
		return fail("unable to create user", err)
	}

	c.Logger().Infof("user created: %s", u.Username)

	err = auth.Login(c, u.ID)
	if err != nil {
		// TODO
	}

	msg.Info(c, "Your account has been created. You are now logged in.")
	return r.Redirect(c, "home")
}
