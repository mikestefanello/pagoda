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
		Email      string `form:"email" validate:"required,email"`
		Submission controller.FormSubmission
	}
)

func (c *ForgotPassword) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "auth"
	page.Name = "forgot-password"
	page.Title = "Forgot password"
	page.Form = ForgotPasswordForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*ForgotPasswordForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *ForgotPassword) Post(ctx echo.Context) error {
	var form ForgotPasswordForm
	ctx.Set(context.FormKey, &form)

	succeed := func() error {
		ctx.Set(context.FormKey, nil)
		msg.Success(ctx, "An email containing a link to reset your password will be sent to this address if it exists in our system.")
		return c.Get(ctx)
	}

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(ctx, err, "unable to parse forgot password form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(ctx, err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	// Attempt to load the user
	u, err := c.Container.ORM.User.
		Query().
		Where(user.Email(form.Email)).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return succeed()
	case nil:
	default:
		return c.Fail(ctx, err, "error querying user during forgot password")
	}

	// Generate the token
	token, _, err := c.Container.Auth.GeneratePasswordResetToken(ctx, u.ID)
	if err != nil {
		return c.Fail(ctx, err, "error generating password reset token")
	}

	ctx.Logger().Infof("generated password reset token for user %d", u.ID)

	// Email the user
	body := fmt.Sprintf(
		"Go here to reset your password: %s",
		ctx.Echo().Reverse("reset_password", u.ID, token),
	)
	ctx.Logger().Info(body)
	err = c.Container.Mail.Send(ctx, u.Email, body)
	if err != nil {
		return c.Fail(ctx, err, "error sending password reset email")
	}

	return succeed()
}
