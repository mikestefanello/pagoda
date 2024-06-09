package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/controller"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/templates"
)

const (
	routeNameContact       = "contact"
	routeNameContactSubmit = "contact.submit"
)

type (
	Contact struct {
		mail *services.MailClient
		controller.Controller
	}

	contactForm struct {
		Email      string `form:"email" validate:"required,email"`
		Department string `form:"department" validate:"required,oneof=sales marketing hr"`
		Message    string `form:"message" validate:"required"`
		Submission controller.FormSubmission
	}
)

func init() {
	Register(new(Contact))
}

func (c *Contact) Init(ct *services.Container) error {
	c.Controller = controller.NewController(ct)
	c.mail = ct.Mail
	return nil
}

func (c *Contact) Routes(g *echo.Group) {
	g.GET("/contact", c.Page).Name = routeNameContact
	g.POST("/contact", c.Submit).Name = routeNameContactSubmit
}

func (c *Contact) Page(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = templates.LayoutMain
	page.Name = templates.PageContact
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
		err := c.mail.
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
