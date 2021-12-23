package routes

import (
	"goweb/context"
	"goweb/controller"

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

	// TODO: Error handling w/ HTMX support

	// Parse the form values
	var form ContactForm
	if err := ctx.Bind(&form); err != nil {
		ctx.Logger().Error(err)
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		// TOOD
	}

	ctx.Set(context.FormKey, form)

	if !form.Submission.HasErrors() {
		if err := c.Container.Mail.Send(ctx, form.Email, "Hello!"); err != nil {
			ctx.Logger().Error(err)
		}
	}

	return c.Get(ctx)
}
