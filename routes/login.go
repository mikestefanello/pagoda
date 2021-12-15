package routes

import (
	"fmt"

	"goweb/auth"
	"goweb/controller"
	"goweb/ent"
	"goweb/ent/user"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type (
	Login struct {
		controller.Controller
		form LoginForm
	}

	LoginForm struct {
		Username string `form:"username" validate:"required" label:"Username"`
		Password string `form:"password" validate:"required" label:"Password"`
	}
)

func (l *Login) Get(c echo.Context) error {
	p := controller.NewPage(c)
	p.Layout = "auth"
	p.Name = "login"
	p.Title = "Log in"
	p.Data = l.form
	return l.RenderPage(c, p)
}

func (l *Login) Post(c echo.Context) error {
	fail := func(message string, err error) error {
		c.Logger().Errorf("%s: %v", message, err)
		msg.Danger(c, "An error occurred. Please try again.")
		return l.Get(c)
	}

	// Parse the form values
	if err := c.Bind(&l.form); err != nil {
		return fail("unable to parse login form", err)
	}

	// Validate the form
	if err := c.Validate(l.form); err != nil {
		l.SetValidationErrorMessages(c, err, l.form)
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
			return fail("error querying user during login", err)
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
		return fail("unable to log in user", err)
	}

	msg.Success(c, fmt.Sprintf("Welcome back, %s. You are now logged in.", u.Username))
	return l.Redirect(c, "home")
}
