package routes

import (
	"goweb/controller"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type (
	ForgotPassword struct {
		controller.Controller
		form ForgotPasswordForm
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
	p.Data = f.form
	return f.RenderPage(c, p)
}

func (f *ForgotPassword) Post(c echo.Context) error {
	fail := func(message string, err error) error {
		c.Logger().Errorf("%s: %v", message, err)
		msg.Danger(c, "An error occurred. Please try again.")
		return f.Get(c)
	}

	// Parse the form values
	if err := c.Bind(&f.form); err != nil {
		return fail("unable to parse forgot password form", err)
	}

	// Validate the form
	if err := c.Validate(f.form); err != nil {
		f.SetValidationErrorMessages(c, err, f.form)
		return f.Get(c)
	}

	// TODO
	return f.Redirect(c, "home")
}
