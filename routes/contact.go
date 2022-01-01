package routes

import (
	"github.com/mikestefanello/pagoda/context"
	"github.com/mikestefanello/pagoda/controller"

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
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "contact"
	page.Title = "Contact us"
	page.Form = ContactForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*ContactForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *Contact) Post(ctx echo.Context) error {
	var form ContactForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(ctx, err, "unable to bind form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(ctx, err, "unable to process form submission")
	}

	if !form.Submission.HasErrors() {
		if err := c.Container.Mail.Send(ctx, form.Email, "Hello!"); err != nil {
			return c.Fail(ctx, err, "unable to send email")
		}
	}

	return c.Get(ctx)
}
