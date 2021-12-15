package routes

import (
	"goweb/context"
	"goweb/controller"
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

	// TODO: generate and email a token

	return f.Redirect(c, "home")
}
