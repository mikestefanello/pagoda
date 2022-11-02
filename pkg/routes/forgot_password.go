package routes

import (
	"fmt"
	"strings"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/msg"

	"github.com/labstack/echo/v4"
)

type (
	forgotPassword struct {
		controller.Controller
	}

	forgotPasswordForm struct {
		Email      string `form:"email" validate:"required,email"`
		Submission controller.FormSubmission
	}
)

func (c *forgotPassword) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "auth"
	page.Name = "forgot-password"
	page.Title = "Forgot password"
	page.Form = forgotPasswordForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*forgotPasswordForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *forgotPassword) Post(ctx echo.Context) error {
	var form forgotPasswordForm
	ctx.Set(context.FormKey, &form)

	succeed := func() error {
		ctx.Set(context.FormKey, nil)
		msg.Success(ctx, "An email containing a link to reset your password will be sent to this address if it exists in our system.")
		return c.Get(ctx)
	}

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse forgot password form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	// Attempt to load the user
	u, err := c.Container.ORM.User.
		Query().
		Where(user.Email(strings.ToLower(form.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return succeed()
	case nil:
	default:
		return c.Fail(err, "error querying user during forgot password")
	}

	// Generate the token
	token, pt, err := c.Container.Auth.GeneratePasswordResetToken(ctx, u.ID)
	if err != nil {
		return c.Fail(err, "error generating password reset token")
	}

	ctx.Logger().Infof("generated password reset token for user %d", u.ID)

	// Email the user
	url := ctx.Echo().Reverse("reset_password", u.ID, pt.ID, token)
	err = c.Container.Mail.
		Compose().
		To(u.Email).
		Subject("Reset your password").
		Body(fmt.Sprintf("Go here to reset your password: %s", url)).
		Send(ctx)

	if err != nil {
		return c.Fail(err, "error sending password reset email")
	}

	return succeed()
}
