package handlers

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/user"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/middleware"
	"github.com/mikestefanello/pagoda/pkg/msg"
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
		controller.Controller
	}

	forgotPasswordForm struct {
		Email      string `form:"email" validate:"required,email"`
		Submission controller.FormSubmission
	}

	loginForm struct {
		Email      string `form:"email" validate:"required,email"`
		Password   string `form:"password" validate:"required"`
		Submission controller.FormSubmission
	}

	registerForm struct {
		Name            string `form:"name" validate:"required"`
		Email           string `form:"email" validate:"required,email"`
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		Submission      controller.FormSubmission
	}

	resetPasswordForm struct {
		Password        string `form:"password" validate:"required"`
		ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
		Submission      controller.FormSubmission
	}
)

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
		middleware.LoadUser(c.Container.ORM),
		middleware.LoadValidPasswordToken(c.Container.Auth),
	)
	resetGroup.GET("/token/:user/:password_token/:token", c.ResetPasswordPage).Name = routeNameResetPassword
	resetGroup.POST("/token/:user/:password_token/:token", c.ResetPasswordSubmit).Name = routeNameResetPasswordSubmit
}

func (c *Auth) ForgotPasswordPage(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutAuth
	page.Name = templates.PageForgotPassword
	page.Title = "Forgot password"
	page.Form = forgotPasswordForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*forgotPasswordForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *Auth) ForgotPasswordSubmit(ctx echo.Context) error {
	var form forgotPasswordForm
	ctx.Set(context.FormKey, &form)

	succeed := func() error {
		ctx.Set(context.FormKey, nil)
		msg.Success(ctx, "An email containing a link to reset your password will be sent to this address if it exists in our system.")
		return c.ForgotPasswordPage(ctx)
	}

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse forgot password form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.ForgotPasswordPage(ctx)
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
	url := ctx.Echo().Reverse(routeNameResetPassword, u.ID, pt.ID, token)
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

func (c *Auth) LoginPage(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutAuth
	page.Name = templates.PageLogin
	page.Title = "Log in"
	page.Form = loginForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*loginForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *Auth) LoginSubmit(ctx echo.Context) error {
	var form loginForm
	ctx.Set(context.FormKey, &form)

	authFailed := func() error {
		form.Submission.SetFieldError("Email", "")
		form.Submission.SetFieldError("Password", "")
		msg.Danger(ctx, "Invalid credentials. Please try again.")
		return c.LoginPage(ctx)
	}

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse login form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.LoginPage(ctx)
	}

	// Attempt to load the user
	u, err := c.Container.ORM.User.
		Query().
		Where(user.Email(strings.ToLower(form.Email))).
		Only(ctx.Request().Context())

	switch err.(type) {
	case *ent.NotFoundError:
		return authFailed()
	case nil:
	default:
		return c.Fail(err, "error querying user during login")
	}

	// Check if the password is correct
	err = c.Container.Auth.CheckPassword(form.Password, u.Password)
	if err != nil {
		return authFailed()
	}

	// Log the user in
	err = c.Container.Auth.Login(ctx, u.ID)
	if err != nil {
		return c.Fail(err, "unable to log in user")
	}

	msg.Success(ctx, fmt.Sprintf("Welcome back, <strong>%s</strong>. You are now logged in.", u.Name))
	return c.Redirect(ctx, routeNameHome)
}

func (c *Auth) Logout(ctx echo.Context) error {
	if err := c.Container.Auth.Logout(ctx); err == nil {
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
	page.Form = registerForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*registerForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *Auth) RegisterSubmit(ctx echo.Context) error {
	var form registerForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse register form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.RegisterPage(ctx)
	}

	// Hash the password
	pwHash, err := c.Container.Auth.HashPassword(form.Password)
	if err != nil {
		return c.Fail(err, "unable to hash password")
	}

	// Attempt creating the user
	u, err := c.Container.ORM.User.
		Create().
		SetName(form.Name).
		SetEmail(form.Email).
		SetPassword(pwHash).
		Save(ctx.Request().Context())

	switch err.(type) {
	case nil:
		ctx.Logger().Infof("user created: %s", u.Name)
	case *ent.ConstraintError:
		msg.Warning(ctx, "A user with this email address already exists. Please log in.")
		return c.Redirect(ctx, routeNameLogin)
	default:
		return c.Fail(err, "unable to create user")
	}

	// Log the user in
	err = c.Container.Auth.Login(ctx, u.ID)
	if err != nil {
		ctx.Logger().Errorf("unable to log in: %v", err)
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
	token, err := c.Container.Auth.GenerateEmailVerificationToken(usr.Email)
	if err != nil {
		ctx.Logger().Errorf("unable to generate email verification token: %v", err)
		return
	}

	// Send the email
	url := ctx.Echo().Reverse(routeNameVerifyEmail, token)
	err = c.Container.Mail.
		Compose().
		To(usr.Email).
		Subject("Confirm your email address").
		Body(fmt.Sprintf("Click here to confirm your email address: %s", url)).
		Send(ctx)

	if err != nil {
		ctx.Logger().Errorf("unable to send email verification link: %v", err)
		return
	}

	msg.Info(ctx, "An email was sent to you to verify your email address.")
}

func (c *Auth) ResetPasswordPage(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutAuth
	page.Name = templates.PageResetPassword
	page.Title = "Reset password"
	page.Form = resetPasswordForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*resetPasswordForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *Auth) ResetPasswordSubmit(ctx echo.Context) error {
	var form resetPasswordForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse password reset form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if form.Submission.HasErrors() {
		return c.ResetPasswordPage(ctx)
	}

	// Hash the new password
	hash, err := c.Container.Auth.HashPassword(form.Password)
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
	err = c.Container.Auth.DeletePasswordTokens(ctx, usr.ID)
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
	email, err := c.Container.Auth.ValidateEmailVerificationToken(token)
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
		usr, err = c.Container.ORM.User.
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
