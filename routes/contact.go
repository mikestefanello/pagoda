package routes

import (
	"net/http"

	"goweb/context"
	"goweb/controller"
	"goweb/msg"

	"github.com/labstack/echo/v4"
)

type (
	Contact struct {
		controller.Controller
	}

	ContactForm struct {
		Email   string `form:"email" validate:"required,email" label:"Email address"`
		Message string `form:"message" validate:"required" label:"Message"`
	}
)

func (c *Contact) Get(ctx echo.Context) error {
	p := controller.NewPage(ctx)
	p.Layout = "main"
	p.Name = "contact"
	p.Title = "Contact us"
	p.Data = ContactForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		p.Data = form.(ContactForm)
	}

	return c.RenderPage(ctx, p)
}

func (c *Contact) Post(ctx echo.Context) error {
	fail := func(message string, err error) error {
		ctx.Logger().Errorf("%s: %v", message, err)
		msg.Danger(ctx, "An error occurred. Please try again.")
		return c.Get(ctx)
	}

	// Parse the form values
	var form ContactForm
	if err := ctx.Bind(&form); err != nil {
		return fail("unable to parse contact form", err)
	}
	ctx.Set(context.FormKey, form)

	// Validate the form
	if err := ctx.Validate(form); err != nil {
		c.SetValidationErrorMessages(ctx, err, form)
		return c.Get(ctx)
	}

	p := controller.NewHTMX(ctx)

	if p.Request.Enabled {
		return ctx.String(http.StatusOK, "<b>HELLO!</b>")
	} else {
		msg.Success(ctx, "Thank you for contacting us!")
		msg.Info(ctx, "We will respond to you shortly.")
		return c.Redirect(ctx, "home")
	}
}
