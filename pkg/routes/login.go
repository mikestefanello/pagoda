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
	login struct {
		controller.Controller
	}

	loginForm struct {
		Email      string `form:"email" validate:"required,email"`
		Password   string `form:"password" validate:"required"`
		Submission controller.FormSubmission
	}
)

func (c *login) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "auth"
	page.Name = "login"
	page.Title = "Log in"
	page.Form = loginForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*loginForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *login) Post(ctx echo.Context) error {
	var form loginForm
	ctx.Set(context.FormKey, &form)

	authFailed := func() error {
		form.Submission.SetFieldError("Email", "")
		form.Submission.SetFieldError("Password", "")
		msg.Danger(ctx, "Invalid credentials. Please try again.")
		return c.Get(ctx)
	}

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to parse login form")
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
	return c.Redirect(ctx, "home")
}
