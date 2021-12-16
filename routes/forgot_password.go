package routes

import (
	"fmt"

	"goweb/context"
	"goweb/controller"
	"goweb/ent"
	"goweb/ent/user"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type (
	ForgotPassword struct {
		controller.Controller
	}

	ForgotPasswordForm struct {
		Email string `form:"email" validate:"required,email" label:"Email address"`
	}
)

func (f *ForgotPassword) Get(c echo.Context) error {
	p := controller.NewPage(c)
	p.Layout = "auth"
	p.Name = "forgot-password"
	p.Title = "Forgot password"
	p.Data = ForgotPasswordForm{}

	if form := c.Get(context.FormKey); form != nil {
		p.Data = form.(ForgotPasswordForm)
	}

	return f.RenderPage(c, p)
}

func (f *ForgotPassword) Post(c echo.Context) error {
	fail := func(message string, err error) error {
		c.Logger().Errorf("%s: %v", message, err)
		msg.Danger(c, "An error occurred. Please try again.")
		return f.Get(c)
	}

	succeed := func() error {
		c.Set(context.FormKey, nil)
		msg.Success(c, "An email containing a link to reset your password will be sent to this address if it exists in our system.")
		return f.Get(c)
	}

	// Parse the form values
	form := new(ForgotPasswordForm)
	if err := c.Bind(form); err != nil {
		return fail("unable to parse forgot password form", err)
	}
	c.Set(context.FormKey, *form)

	// Validate the form
	if err := c.Validate(form); err != nil {
		f.SetValidationErrorMessages(c, err, form)
		return f.Get(c)
	}

	// Attempt to load the user
	u, err := f.Container.ORM.User.
		Query().
		Where(user.Email(form.Email)).
		First(c.Request().Context())

	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return succeed()
		default:
			return fail("error querying user during forgot password", err)
		}
	}

	// Generate the token
	token, _, err := f.Container.Auth.GeneratePasswordResetToken(c, u.ID)
	if err != nil {
		return fail("error generating password reset token", err)
	}
	c.Logger().Infof("generated password reset token for user %d", u.ID)

	// Email the user
	err = f.Container.Mail.Send(c, u.Email, fmt.Sprintf("Go here to reset your password: %s", token)) // TODO: route
	if err != nil {
		return fail("error sending password reset email", err)
	}

	return succeed()
}
