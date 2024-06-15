package handlers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/msg"
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
		controller.Controller
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

func (c *Auth) Init(ct *services.Container) error {
	c.Controller = controller.NewController(ct)
	c.orm = ct.ORM
	c.auth = ct.Auth
	c.mail = ct.Mail
	return nil
}

func (c *Auth) Routes(g *echo.Group) {
	g.GET("/logout", c.Logout, middleware.RequireAuthentication()).Name = routeNameLogout
	g.GET("/email/verify/:token", c.VerifyEmail).Name = routeNameVerifyEmail

	noAuth := g.Group("/user", middleware.RequireNoAuthentication())
	noAuth.GET("/login", c.LoginPage).Name = routeNameLogin
	noAuth.POST("/login", c.LoginSubmit).Name = routeNameLoginSubmit
	noAuth.GET("/register", c.RegisterPage).Name = routeNameRegister
	noAuth.POST("/register", c.RegisterSubmit).Name = routeNameRegisterSubmit
	noAuth.GET("/password", c.ForgotPasswordPage).Name = routeNameForgotPassword
	noAuth.POST("/password", c.ForgotPasswordSubmit).Name = routeNameForgotPasswordSubmit

	resetGroup := noAuth.Group("/password/reset",
		middleware.LoadUser(c.orm),
		middleware.LoadValidPasswordToken(c.auth),
	)
	resetGroup.GET("/token/:user/:password_token/:token", c.ResetPasswordPage).Name = routeNameResetPassword
	resetGroup.POST("/token/:user/:password_token/:token", c.ResetPasswordSubmit).Name = routeNameResetPasswordSubmit
}

func (c *Auth) ForgotPasswordPage(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutAuth
	page.Name = templates.PageForgotPassword
	page.Title = "Forgot password"
	page.Form = form.Get[forgotPasswordForm](ctx)

	return c.RenderPage(ctx, page)
}

func (c *Auth) ForgotPasswordSubmit(ctx echo.Context) error {
	var input forgotPasswordForm

	succeed := func() error {
		form.Clear(ctx)
		msg.Success(ctx, "An email containing a link to reset your password will be sent to this address if it exists in our system.")
		return c.ForgotPasswordPage(ctx)
	}

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return c.ForgotPasswordPage(ctx)
	default:
		return err
	}

	// Attempt to load the user
	u, err := c.orm.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return succeed()
	case nil:
	default:
		return c.Fail(err, "error querying user during forgot password")
	}

	// Generate the token
	token, pt, err := c.auth.GeneratePasswordResetToken(ctx, u.ID)
	if err != nil {
		return c.Fail(err, "error generating password reset token")
	}

	log.Ctx(ctx).Info("generated password reset token",
		"user_id", u.ID,
	)

	// Email the user
	url := ctx.Echo().Reverse(routeNameResetPassword, u.ID, pt.ID, token)
	err = c.mail.
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

func (c *Auth) LoginPage(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutAuth
	page.Name = templates.PageLogin
	page.Title = "Log in"
	page.Form = form.Get[loginForm](ctx)

	return c.RenderPage(ctx, page)
}

func (c *Auth) LoginSubmit(ctx echo.Context) error {
	var input loginForm

	authFailed := func() error {
		input.SetFieldError("Email", "")
		input.SetFieldError("Password", "")
		msg.Danger(ctx, "Invalid credentials. Please try again.")
		return c.LoginPage(ctx)
	}

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return c.LoginPage(ctx)
	default:
		return err
	}

	// Attempt to load the user
	u, err := c.orm.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return authFailed()
	case nil:
	default:
		return c.Fail(err, "error querying user during login")
	}

	// Check if the password is correct
	err = c.auth.CheckPassword(input.Password, u.Password)
	if err != nil {
		return authFailed()
	}

	// Log the user in
	err = c.auth.Login(ctx, u.ID)
	if err != nil {
		return c.Fail(err, "unable to log in user")
	}

	msg.Success(ctx, fmt.Sprintf("Welcome back, <strong>%s</strong>. You are now logged in.", u.Name))
	return c.Redirect(ctx, routeNameHome)
}

func (c *Auth) Logout(ctx echo.Context) error {
	if err := c.auth.Logout(ctx); err == nil {
		msg.Success(ctx, "You have been logged out successfully.")
	} else {
		msg.Danger(ctx, "An error occurred. Please try again.")
	}
	return c.Redirect(ctx, routeNameHome)
}

func (c *Auth) RegisterPage(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutAuth
	page.Name = templates.PageRegister
	page.Title = "Register"
	page.Form = form.Get[registerForm](ctx)

	return c.RenderPage(ctx, page)
}

func (c *Auth) RegisterSubmit(ctx echo.Context) error {
	var input registerForm

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return c.RegisterPage(ctx)
	default:
		return err
	}

	// Hash the password
	pwHash, err := c.auth.HashPassword(input.Password)
	if err != nil {
		return c.Fail(err, "unable to hash password")
	}

	// Attempt creating the user
	u, err := c.orm.User.
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
		return c.Redirect(ctx, routeNameLogin)
	default:
		return c.Fail(err, "unable to create user")
	}

	// Log the user in
	err = c.auth.Login(ctx, u.ID)
	if err != nil {
		log.Ctx(ctx).Error("unable to log user in",
			"error", err,
			"user_id", u.ID,
		)
		msg.Info(ctx, "Your account has been created.")
		return c.Redirect(ctx, routeNameLogin)
	}

	msg.Success(ctx, "Your account has been created. You are now logged in.")

	// Send the verification email
	c.sendVerificationEmail(ctx, u)

	return c.Redirect(ctx, routeNameHome)
}

func (c *Auth) sendVerificationEmail(ctx echo.Context, usr *ent.User) {
	// Generate a token
	token, err := c.auth.GenerateEmailVerificationToken(usr.Email)
	if err != nil {
		log.Ctx(ctx).Error("unable to generate email verification token",
			"user_id", usr.ID,
			"error", err,
		)
		return
	}

	// Send the email
	url := ctx.Echo().Reverse(routeNameVerifyEmail, token)
	err = c.mail.
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

func (c *Auth) ResetPasswordPage(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutAuth
	page.Name = templates.PageResetPassword
	page.Title = "Reset password"
	page.Form = form.Get[resetPasswordForm](ctx)

	return c.RenderPage(ctx, page)
}

func (c *Auth) ResetPasswordSubmit(ctx echo.Context) error {
	var input resetPasswordForm

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return c.ResetPasswordPage(ctx)
	default:
		return err
	}

	// Hash the new password
	hash, err := c.auth.HashPassword(input.Password)
	if err != nil {
		return c.Fail(err, "unable to hash password")
	}

	// Get the requesting user
	usr := ctx.Get(context.UserKey).(*ent.User)

	// Update the user
	_, err = usr.
		Update().
		SetPassword(hash).
		Save(ctx.Request().Context())

	if err != nil {
		return c.Fail(err, "unable to update password")
	}

	// Delete all password tokens for this user
	err = c.auth.DeletePasswordTokens(ctx, usr.ID)
	if err != nil {
		return c.Fail(err, "unable to delete password tokens")
	}

	msg.Success(ctx, "Your password has been updated.")
	return c.Redirect(ctx, routeNameLogin)
}

func (c *Auth) VerifyEmail(ctx echo.Context) error {
	var usr *ent.User

	// Validate the token
	token := ctx.Param("token")
	email, err := c.auth.ValidateEmailVerificationToken(token)
	if err != nil {
		msg.Warning(ctx, "The link is either invalid or has expired.")
		return c.Redirect(ctx, routeNameHome)
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
		usr, err = c.orm.User.
			Query().
			Where(user.Email(email)).
			Only(ctx.Request().Context())

		if err != nil {
			return c.Fail(err, "query failed loading email verification token user")
		}
	}

	// Verify the user, if needed
	if !usr.Verified {
		usr, err = usr.
			Update().
			SetVerified(true).
			Save(ctx.Request().Context())

		if err != nil {
			return c.Fail(err, "failed to set user as verified")
		}
	}

	msg.Success(ctx, "Your email has been successfully verified.")
	return c.Redirect(ctx, routeNameHome)
}
