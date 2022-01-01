package routes

import (
	"github.com/mikestefanello/pagoda/context"
	"github.com/mikestefanello/pagoda/controller"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/msg"

	"github.com/labstack/echo/v4"
)

type (
	ResetPassword struct {
		controller.Controller
	}

	ResetPasswordForm struct {
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		Submission      controller.FormSubmission
	}
)

func (c *ResetPassword) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "auth"
	page.Name = "reset-password"
	page.Title = "Reset password"
	page.Form = ResetPasswordForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*ResetPasswordForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *ResetPassword) Post(ctx echo.Context) error {
	var form ResetPasswordForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(ctx, err, "unable to parse password reset form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(ctx, err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	// Hash the new password
	hash, err := c.Container.Auth.HashPassword(form.Password)
	if err != nil {
		return c.Fail(ctx, err, "unable to hash password")
	}

	// Get the requesting user
	usr := ctx.Get(context.UserKey).(*ent.User)

	// Update the user
	_, err = c.Container.ORM.User.
		Update().
		SetPassword(hash).
		Where(user.ID(usr.ID)).
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(ctx, err, "unable to update password")
	}

	// Delete all password tokens for this user
	err = c.Container.Auth.DeletePasswordTokens(ctx, usr.ID)
	if err != nil {
		return c.Fail(ctx, err, "unable to delete password tokens")
	}

	msg.Success(ctx, "Your password has been updated.")
	return c.Redirect(ctx, "login")
}
