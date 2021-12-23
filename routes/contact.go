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
		Email      string `form:"email" validate:"required,email"`
		Message    string `form:"message" validate:"required"`
		Submission controller.FormSubmission
	}
)

func (c *Contact) Get(ctx echo.Context) error {
	p := controller.NewPage(ctx)
	p.Layout = "main"
	p.Name = "contact"
	p.Title = "Contact us"
	p.Form = ContactForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		p.Form = form.(ContactForm)
	}

	return c.RenderPage(ctx, p)
}

func (c *Contact) Post(ctx echo.Context) error {
	//fail := func(message string, err error) error {
	//	ctx.Logger().Errorf("%s: %v", message, err)
	//	msg.Danger(ctx, "An error occurred. Please try again.")
	//	return c.Get(ctx)
	//}

	// Parse the form values
	var form ContactForm
	if err := ctx.Bind(&form); err != nil {
		ctx.Logger().Error(err)
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		// TOOD
	}

	ctx.Set(context.FormKey, form)

	if form.Submission.HasErrors() {
		return c.Get(ctx)
	}

	htmx := controller.GetHTMXRequest(ctx)

	if htmx.Enabled {
		return ctx.String(http.StatusOK, "<b>HELLO!</b>")
	} else {
		msg.Success(ctx, "Thank you for contacting us!")
		msg.Info(ctx, "We will respond to you shortly.")
		return c.Redirect(ctx, "home")
	}
}
