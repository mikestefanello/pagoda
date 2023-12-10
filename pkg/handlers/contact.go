package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
)

const (
	routeNameContact       = "contact"
	routeNameContactSubmit = "contact.submit"
)

type (
	Contact struct {
		controller.Controller
	}

	contactForm struct {
		Email      string `form:"email" validate:"required,email"`
		Message    string `form:"message" validate:"required"`
		Submission controller.FormSubmission
	}
)

func (c *Contact) Routes(g *echo.Group) {
	g.GET("/contact", c.Page).Name = routeNameContact
	g.POST("/contact", c.Submit).Name = routeNameContactSubmit
}

func (c *Contact) Page(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "contact"
	page.Title = "Contact us"
	page.Form = contactForm{}

	if form := ctx.Get(context.FormKey); form != nil {
		page.Form = form.(*contactForm)
	}

	return c.RenderPage(ctx, page)
}

func (c *Contact) Submit(ctx echo.Context) error {
	var form contactForm
	ctx.Set(context.FormKey, &form)

	// Parse the form values
	if err := ctx.Bind(&form); err != nil {
		return c.Fail(err, "unable to bind form")
	}

	if err := form.Submission.Process(ctx, form); err != nil {
		return c.Fail(err, "unable to process form submission")
	}

	if !form.Submission.HasErrors() {
		err := c.Container.Mail.
			Compose().
			To(form.Email).
			Subject("Contact form submitted").
			Body(fmt.Sprintf("The message is: %s", form.Message)).
			Send(ctx)

		if err != nil {
			return c.Fail(err, "unable to send email")
		}
	}

	return c.Page(ctx)
}
