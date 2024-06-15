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
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/templates"
)

const (
	routeNameForgotPassword       = "forgot_password"
	routeNameForgotPasswordSubmit = "forgot_password.submit"
	routeNameLogin                = "login"
	routeNameLoginSubmit          = "login.submit"
	routeNameLogout               = "logout"
	routeNameRegister             = "register"
	routeNameRegisterSubmit       = "register.submit"
	routeNameResetPassword        = "reset_password"
	routeNameResetPasswordSubmit  = "reset_password.submit"
	routeNameVerifyEmail          = "verify_email"
)

type (
	Auth struct {
		auth *services.AuthClient
		mail *services.MailClient
		orm  *ent.Client
		*services.TemplateRenderer
	}

	forgotPasswordForm struct {
		Email string `form:"email" validate:"required,email"`
		form.Submission
	}

	loginForm struct {
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required"`
		form.Submission
	}

	registerForm struct {
		Name            string `form:"name" validate:"required"`
		Email           string `form:"email" validate:"required,email"`
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		form.Submission
	}

	resetPasswordForm struct {
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		form.Submission
	}
)

func init() {
	Register(new(Auth))
}

func (h *Auth) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.orm = c.ORM
	h.auth = c.Auth
	h.mail = c.Mail
	return nil
}

func (h *Auth) Routes(g *echo.Group) {
	g.GET("/logout", h.Logout, middleware.RequireAuthentication()).Name = routeNameLogout
	g.GET("/email/verify/:token", h.VerifyEmail).Name = routeNameVerifyEmail

	noAuth := g.Group("/user", middleware.RequireNoAuthentication())
	noAuth.GET("/login", h.LoginPage).Name = routeNameLogin
	noAuth.POST("/login", h.LoginSubmit).Name = routeNameLoginSubmit
	noAuth.GET("/register", h.RegisterPage).Name = routeNameRegister
	noAuth.POST("/register", h.RegisterSubmit).Name = routeNameRegisterSubmit
	noAuth.GET("/password", h.ForgotPasswordPage).Name = routeNameForgotPassword
	noAuth.POST("/password", h.ForgotPasswordSubmit).Name = routeNameForgotPasswordSubmit

	resetGroup := noAuth.Group("/password/reset",
		middleware.LoadUser(h.orm),
		middleware.LoadValidPasswordToken(h.auth),
	)
	resetGroup.GET("/token/:user/:password_token/:token", h.ResetPasswordPage).Name = routeNameResetPassword
	resetGroup.POST("/token/:user/:password_token/:token", h.ResetPasswordSubmit).Name = routeNameResetPasswordSubmit
}

func (h *Auth) ForgotPasswordPage(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutAuth
	p.Name = templates.PageForgotPassword
	p.Title = "Forgot password"
	p.Form = form.Get[forgotPasswordForm](ctx)

	return h.RenderPage(ctx, p)
}

func (h *Auth) ForgotPasswordSubmit(ctx echo.Context) error {
	var input forgotPasswordForm

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

	// Attempt to load the user
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

	// Generate the token
	token, pt, err := h.auth.GeneratePasswordResetToken(ctx, u.ID)
	if err != nil {
		return fail(err, "error generating password reset token")
	}

	log.Ctx(ctx).Info("generated password reset token",
		"user_id", u.ID,
	)

	// Email the user
	url := ctx.Echo().Reverse(routeNameResetPassword, u.ID, pt.ID, token)
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
	p := page.New(ctx)
	p.Layout = templates.LayoutAuth
	p.Name = templates.PageLogin
	p.Title = "Log in"
	p.Form = form.Get[loginForm](ctx)

	return h.RenderPage(ctx, p)
}

func (h *Auth) LoginSubmit(ctx echo.Context) error {
	var input loginForm

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

	// Attempt to load the user
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

	// Check if the password is correct
	err = h.auth.CheckPassword(input.Password, u.Password)
	if err != nil {
		return authFailed()
	}

	// Log the user in
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		return fail(err, "unable to log in user")
	}

	msg.Success(ctx, fmt.Sprintf("Welcome back, <strong>%s</strong>. You are now logged in.", u.Name))
	return redirect(ctx, routeNameHome)
}

func (h *Auth) Logout(ctx echo.Context) error {
	if err := h.auth.Logout(ctx); err == nil {
		msg.Success(ctx, "You have been logged out successfully.")
	} else {
		msg.Danger(ctx, "An error occurred. Please try again.")
	}
	return redirect(ctx, routeNameHome)
}

func (h *Auth) RegisterPage(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutAuth
	p.Name = templates.PageRegister
	p.Title = "Register"
	p.Form = form.Get[registerForm](ctx)

	return h.RenderPage(ctx, p)
}

func (h *Auth) RegisterSubmit(ctx echo.Context) error {
	var input registerForm

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.RegisterPage(ctx)
	default:
		return err
	}

	// Hash the password
	pwHash, err := h.auth.HashPassword(input.Password)
	if err != nil {
		return fail(err, "unable to hash password")
	}

	// Attempt creating the user
	u, err := h.orm.User.
		Create().
		SetName(input.Name).
		SetEmail(input.Email).
		SetPassword(pwHash).
		Save(ctx.Request().Context())

	switch err.(type) {
	case nil:
		log.Ctx(ctx).Info("user created",
			"user_name", u.Name,
			"user_id", u.ID,
		)
	case *ent.ConstraintError:
		msg.Warning(ctx, "A user with this email address already exists. Please log in.")
		return redirect(ctx, routeNameLogin)
	default:
		return fail(err, "unable to create user")
	}

	// Log the user in
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		log.Ctx(ctx).Error("unable to log user in",
			"error", err,
			"user_id", u.ID,
		)
		msg.Info(ctx, "Your account has been created.")
		return redirect(ctx, routeNameLogin)
	}

	msg.Success(ctx, "Your account has been created. You are now logged in.")

	// Send the verification email
	h.sendVerificationEmail(ctx, u)

	return redirect(ctx, routeNameHome)
}

func (h *Auth) sendVerificationEmail(ctx echo.Context, usr *ent.User) {
	// Generate a token
	token, err := h.auth.GenerateEmailVerificationToken(usr.Email)
	if err != nil {
		log.Ctx(ctx).Error("unable to generate email verification token",
			"user_id", usr.ID,
			"error", err,
		)
		return
	}

	// Send the email
	url := ctx.Echo().Reverse(routeNameVerifyEmail, token)
	err = h.mail.
		Compose().
		To(usr.Email).
		Subject("Confirm your email address").
		Body(fmt.Sprintf("Click here to confirm your email address: %s", url)).
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
	p := page.New(ctx)
	p.Layout = templates.LayoutAuth
	p.Name = templates.PageResetPassword
	p.Title = "Reset password"
	p.Form = form.Get[resetPasswordForm](ctx)

	return h.RenderPage(ctx, p)
}

func (h *Auth) ResetPasswordSubmit(ctx echo.Context) error {
	var input resetPasswordForm

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.ResetPasswordPage(ctx)
	default:
		return err
	}

	// Hash the new password
	hash, err := h.auth.HashPassword(input.Password)
	if err != nil {
		return fail(err, "unable to hash password")
	}

	// Get the requesting user
	usr := ctx.Get(context.UserKey).(*ent.User)

	// Update the user
	_, err = usr.
		Update().
		SetPassword(hash).
		Save(ctx.Request().Context())

	if err != nil {
		return fail(err, "unable to update password")
	}

	// Delete all password tokens for this user
	err = h.auth.DeletePasswordTokens(ctx, usr.ID)
	if err != nil {
		return fail(err, "unable to delete password tokens")
	}

	msg.Success(ctx, "Your password has been updated.")
	return redirect(ctx, routeNameLogin)
}

func (h *Auth) VerifyEmail(ctx echo.Context) error {
	var usr *ent.User

	// Validate the token
	token := ctx.Param("token")
	email, err := h.auth.ValidateEmailVerificationToken(token)
	if err != nil {
		msg.Warning(ctx, "The link is either invalid or has expired.")
		return redirect(ctx, routeNameHome)
	}

	// Check if it matches the authenticated user
	if u := ctx.Get(context.AuthenticatedUserKey); u != nil {
		authUser := u.(*ent.User)

		if authUser.Email == email {
			usr = authUser
		}
	}

	// Query to find a matching user, if needed
	if usr == nil {
		usr, err = h.orm.User.
			Query().
			Where(user.Email(email)).
			Only(ctx.Request().Context())

		if err != nil {
			return fail(err, "query failed loading email verification token user")
		}
	}

	// Verify the user, if needed
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
	return redirect(ctx, routeNameHome)
}
