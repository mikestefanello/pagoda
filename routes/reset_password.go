package routes

import (
	"goweb/controller"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type (
	ResetPassword struct {
		controller.Controller
	}

	ResetPasswordForm struct {
		Password        string `form:"password" validate:"required" label:"Password"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password" label:"Confirm password"`
	}
)

func (r *ResetPassword) Get(c echo.Context) error {
	p := controller.NewPage(c)
	p.Layout = "auth"
	p.Name = "reset-password"
	p.Title = "Reset password"
	return r.RenderPage(c, p)
}

func (r *ResetPassword) Post(c echo.Context) error {
	fail := func(message string, err error) error {
		c.Logger().Errorf("%s: %v", message, err)
		msg.Danger(c, "An error occurred. Please try again.")
		return r.Get(c)
	}

	succeed := func() error {
		msg.Success(c, "Your password has been updated.")
		return r.Redirect(c, "login")
	}

	// Parse the form values
	form := new(ResetPassword)
	if err := c.Bind(form); err != nil {
		return fail("unable to parse forgot password form", err)
	}

	// Validate the form
	if err := c.Validate(form); err != nil {
		r.SetValidationErrorMessages(c, err, form)
		return r.Get(c)
	}

	// Attempt to load the user
	//u, err := f.Container.ORM.User.
	//	Query().
	//	Where(user.Email(form.Email)).
	//	First(c.Request().Context())
	//
	//if err != nil {
	//	switch err.(type) {
	//	case *ent.NotFoundError:
	//		return succeed()
	//	default:
	//		return fail("error querying user during forgot password", err)
	//	}
	//}
	//
	//// Generate the token
	//token, _, err := f.Container.Auth.GeneratePasswordResetToken(c, u.ID)
	//if err != nil {
	//	return fail("error generating password reset token", err)
	//}
	//c.Logger().Infof("generated password reset token for user %d", u.ID)
	//
	//// Email the user
	//err = f.Container.Mail.Send(c, u.Email, fmt.Sprintf("Go here to reset your password: %s", token)) // TODO: route
	//if err != nil {
	//	return fail("error sending password reset email", err)
	//}

	return succeed()
}
