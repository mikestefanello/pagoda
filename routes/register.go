package routes

import (
	"goweb/context"
	"goweb/controller"
	"goweb/ent"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type (
	Register struct {
		controller.Controller
	}

	RegisterForm struct {
		Name            string `form:"name" validate:"required" label:"Name"`
		Email           string `form:"email" validate:"required,email" label:"Email address"`
		Password        string `form:"password" validate:"required" label:"Password"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password" label:"Confirm password"`
	}
)

func (r *Register) Get(c echo.Context) error {
	p := controller.NewPage(c)
	p.Layout = "auth"
	p.Name = "register"
	p.Title = "Register"
	p.Data = RegisterForm{}

	if form := c.Get(context.FormKey); form != nil {
		p.Data = form.(RegisterForm)
	}

	return r.RenderPage(c, p)
}

func (r *Register) Post(c echo.Context) error {
	fail := func(message string, err error) error {
		c.Logger().Errorf("%s: %v", message, err)
		msg.Danger(c, "An error occurred. Please try again.")
		return r.Get(c)
	}

	// Parse the form values
	form := new(RegisterForm)
	if err := c.Bind(form); err != nil {
		return fail("unable to parse form values", err)
	}
	c.Set(context.FormKey, *form)

	// Validate the form
	if err := c.Validate(form); err != nil {
		r.SetValidationErrorMessages(c, err, form)
		return r.Get(c)
	}

	// Hash the password
	pwHash, err := r.Container.Auth.HashPassword(form.Password)
	if err != nil {
		return fail("unable to hash password", err)
	}

	// Attempt creating the user
	u, err := r.Container.ORM.User.
		Create().
		SetName(form.Name).
		SetEmail(form.Email).
		SetPassword(pwHash).
		Save(c.Request().Context())

	switch err.(type) {
	case nil:
		c.Logger().Infof("user created: %s", u.Name)
	case *ent.ConstraintError:
		msg.Warning(c, "A user with this email address already exists. Please log in.")
		return r.Redirect(c, "login")
	default:
		return fail("unable to create user", err)
	}

	// Log the user in
	err = r.Container.Auth.Login(c, u.ID)
	if err != nil {
		c.Logger().Errorf("unable to log in: %v", err)
		msg.Info(c, "Your account has been created.")
		return r.Redirect(c, "login")
	}

	msg.Info(c, "Your account has been created. You are now logged in.")
	return r.Redirect(c, "home")
}
