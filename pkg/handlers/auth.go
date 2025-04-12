package handlers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/redirect"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/ui/emails"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/pages"
)

type Auth struct {
	auth *services.AuthClient
	mail *services.MailClient
	orm  *ent.Client
}

func init() {
	Register(new(Auth))
}

func (h *Auth) Init(c *services.Container) error {
	h.orm = c.ORM
	h.auth = c.Auth
	h.mail = c.Mail
	return nil
}

func (h *Auth) Routes(g *echo.Group) {
	g.GET("/logout", h.Logout, middleware.RequireAuthentication).Name = routenames.Logout
	g.GET("/email/verify/:token", h.VerifyEmail).Name = routenames.VerifyEmail

	noAuth := g.Group("/user", middleware.RequireNoAuthentication)
	noAuth.GET("/login", h.LoginPage).Name = routenames.Login
	noAuth.POST("/login", h.LoginSubmit).Name = routenames.LoginSubmit
	noAuth.GET("/register", h.RegisterPage).Name = routenames.Register
	noAuth.POST("/register", h.RegisterSubmit).Name = routenames.RegisterSubmit
	noAuth.GET("/password", h.ForgotPasswordPage).Name = routenames.ForgotPassword
	noAuth.POST("/password", h.ForgotPasswordSubmit).Name = routenames.ForgotPasswordSubmit

	resetGroup := noAuth.Group("/password/reset",
		middleware.LoadUser(h.orm),
		middleware.LoadValidPasswordToken(h.auth),
	)
	resetGroup.GET("/token/:user/:password_token/:token", h.ResetPasswordPage).Name = routenames.ResetPassword
	resetGroup.POST("/token/:user/:password_token/:token", h.ResetPasswordSubmit).Name = routenames.ResetPasswordSubmit
}

func (h *Auth) ForgotPasswordPage(ctx echo.Context) error {
	return pages.ForgotPassword(ctx, form.Get[forms.ForgotPassword](ctx))
}

func (h *Auth) ForgotPasswordSubmit(ctx echo.Context) error {
	var input forms.ForgotPassword

	succeed := func() error {
		form.Clear(ctx)
		msg.Success(ctx, "An email containing a link to reset your password will be sent to this address if it exists in our system.")
		return h.ForgotPasswordPage(ctx)
	}

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.ForgotPasswordPage(ctx)
	default:
		return err
	}

	// Attempt to load the user.
	u, err := h.orm.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return succeed()
	case nil:
	default:
		return fail(err, "error querying user during forgot password")
	}

	// Generate the token.
	token, pt, err := h.auth.GeneratePasswordResetToken(ctx, u.ID)
	if err != nil {
		return fail(err, "error generating password reset token")
	}

	log.Ctx(ctx).Info("generated password reset token",
		"user_id", u.ID,
	)

	// Email the user.
	url := ctx.Echo().Reverse(routenames.ResetPassword, u.ID, pt.ID, token)
	err = h.mail.
		Compose().
		To(u.Email).
		Subject("Reset your password").
		Body(fmt.Sprintf("Go here to reset your password: %s", url)).
		Send(ctx)

	if err != nil {
		return fail(err, "error sending password reset email")
	}

	return succeed()
}

func (h *Auth) LoginPage(ctx echo.Context) error {
	return pages.Login(ctx, form.Get[forms.Login](ctx))
}

func (h *Auth) LoginSubmit(ctx echo.Context) error {
	var input forms.Login

	authFailed := func() error {
		input.SetFieldError("Email", "")
		input.SetFieldError("Password", "")
		msg.Danger(ctx, "Invalid credentials. Please try again.")
		return h.LoginPage(ctx)
	}

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.LoginPage(ctx)
	default:
		return err
	}

	// Attempt to load the user.
	u, err := h.orm.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return authFailed()
	case nil:
	default:
		return fail(err, "error querying user during login")
	}

	// Check if the password is correct.
	err = h.auth.CheckPassword(input.Password, u.Password)
	if err != nil {
		return authFailed()
	}

	// Log the user in.
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		return fail(err, "unable to log in user")
	}

	msg.Success(ctx, fmt.Sprintf("Welcome back, %s. You are now logged in.", u.Name))

	return redirect.New(ctx).
		Route(routenames.Home).
		Go()
}

func (h *Auth) Logout(ctx echo.Context) error {
	if err := h.auth.Logout(ctx); err == nil {
		msg.Success(ctx, "You have been logged out successfully.")
	} else {
		msg.Danger(ctx, "An error occurred. Please try again.")
	}
	return redirect.New(ctx).
		Route(routenames.Home).
		Go()
}

func (h *Auth) RegisterPage(ctx echo.Context) error {
	return pages.Register(ctx, form.Get[forms.Register](ctx))
}

func (h *Auth) RegisterSubmit(ctx echo.Context) error {
	var input forms.Register

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.RegisterPage(ctx)
	default:
		return err
	}

	// Attempt creating the user.
	u, err := h.orm.User.
		Create().
		SetName(input.Name).
		SetEmail(input.Email).
		SetPassword(input.Password).
		Save(ctx.Request().Context())

	switch err.(type) {
	case nil:
		log.Ctx(ctx).Info("user created",
			"user_name", u.Name,
			"user_id", u.ID,
		)
	case *ent.ConstraintError:
		msg.Warning(ctx, "A user with this email address already exists. Please log in.")
		return redirect.New(ctx).
			Route(routenames.Login).
			Go()
	default:
		return fail(err, "unable to create user")
	}

	// Log the user in.
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		log.Ctx(ctx).Error("unable to log user in",
			"error", err,
			"user_id", u.ID,
		)
		msg.Info(ctx, "Your account has been created.")
		return redirect.New(ctx).
			Route(routenames.Login).
			Go()
	}

	msg.Success(ctx, "Your account has been created. You are now logged in.")

	// Send the verification email.
	h.sendVerificationEmail(ctx, u)

	return redirect.New(ctx).
		Route(routenames.Home).
		Go()
}

func (h *Auth) sendVerificationEmail(ctx echo.Context, usr *ent.User) {
	// Generate a token.
	token, err := h.auth.GenerateEmailVerificationToken(usr.Email)
	if err != nil {
		log.Ctx(ctx).Error("unable to generate email verification token",
			"user_id", usr.ID,
			"error", err,
		)
		return
	}

	// Send the email.
	err = h.mail.
		Compose().
		To(usr.Email).
		Subject("Confirm your email address").
		Component(emails.ConfirmEmailAddress(ctx, usr.Name, token)).
		Send(ctx)

	if err != nil {
		log.Ctx(ctx).Error("unable to send email verification link",
			"user_id", usr.ID,
			"error", err,
		)
		return
	}

	msg.Info(ctx, "An email was sent to you to verify your email address.")
}

func (h *Auth) ResetPasswordPage(ctx echo.Context) error {
	return pages.ResetPassword(ctx, form.Get[forms.ResetPassword](ctx))
}

func (h *Auth) ResetPasswordSubmit(ctx echo.Context) error {
	var input forms.ResetPassword

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.ResetPasswordPage(ctx)
	default:
		return err
	}

	// Get the requesting user.
	usr := ctx.Get(context.UserKey).(*ent.User)

	// Update the user.
	_, err = usr.
		Update().
		SetPassword(input.Password).
		Save(ctx.Request().Context())

	if err != nil {
		return fail(err, "unable to update password")
	}

	// Delete all password tokens for this user.
	err = h.auth.DeletePasswordTokens(ctx, usr.ID)
	if err != nil {
		return fail(err, "unable to delete password tokens")
	}

	msg.Success(ctx, "Your password has been updated.")
	return redirect.New(ctx).
		Route(routenames.Login).
		Go()
}

func (h *Auth) VerifyEmail(ctx echo.Context) error {
	var usr *ent.User

	// Validate the token.
	token := ctx.Param("token")
	email, err := h.auth.ValidateEmailVerificationToken(token)
	if err != nil {
		msg.Warning(ctx, "The link is either invalid or has expired.")
		return redirect.New(ctx).
			Route(routenames.Home).
			Go()
	}

	// Check if it matches the authenticated user.
	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		authUser := u.(*ent.User)

		if authUser.Email == email {
			usr = authUser
		}
	}

	// Query to find a matching user, if needed.
	if usr == nil {
		usr, err = h.orm.User.
			Query().
			Where(user.Email(email)).
			Only(ctx.Request().Context())

		if err != nil {
			return fail(err, "query failed loading email verification token user")
		}
	}

	// Verify the user, if needed.
	if !usr.Verified {
		usr, err = usr.
			Update().
			SetVerified(true).
			Save(ctx.Request().Context())

		if err != nil {
			return fail(err, "failed to set user as verified")
		}
	}

	msg.Success(ctx, "Your email has been successfully verified.")
	return redirect.New(ctx).
		Route(routenames.Home).
		Go()
}
