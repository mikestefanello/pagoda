package routes

import (
	"goweb/context"
	"goweb/controller"
	"goweb/ent"
	"goweb/ent/user"
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

	// Parse the form values
	var form ResetPasswordForm
	if err := c.Bind(&form); err != nil {
		return fail("unable to parse forgot password form", err)
	}

	// Validate the form
	if err := c.Validate(form); err != nil {
		r.SetValidationErrorMessages(c, err, form)
		return r.Get(c)
	}

	// Hash the new password
	hash, err := r.Container.Auth.HashPassword(form.Password)
	if err != nil {
		return fail("unable to hash password", err)
	}

	// Get the requesting user
	usr := c.Get(context.UserKey).(*ent.User)

	// Update the user
	_, err = r.Container.ORM.User.
		Update().
		SetPassword(hash).
		Where(user.ID(usr.ID)).
		Save(c.Request().Context())

	if err != nil {
		return fail("unable to update password", err)
	}

	// Delete all password tokens for this user
	err = r.Container.Auth.DeletePasswordTokens(c, usr.ID)
	if err != nil {
		return fail("unable to delete password tokens", err)
	}

	msg.Success(c, "Your password has been updated.")
	return r.Redirect(c, "login")
}
