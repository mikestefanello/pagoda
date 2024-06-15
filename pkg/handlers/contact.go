package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/page"
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
		*services.TemplateRenderer
	}

	contactForm struct {
		Email      string `form:"email" validate:"required,email"`
		Department string `form:"department" validate:"required,oneof=sales marketing hr"`
		Message    string `form:"message" validate:"required"`
		form.Submission
	}
)

func init() {
	Register(new(Contact))
}

func (c *Contact) Init(ct *services.Container) error {
	c.TemplateRenderer = ct.TemplateRenderer
	c.mail = ct.Mail
	return nil
}

func (c *Contact) Routes(g *echo.Group) {
	g.GET("/contact", c.Page).Name = routeNameContact
	g.POST("/contact", c.Submit).Name = routeNameContactSubmit
}

func (c *Contact) Page(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageContact
	p.Title = "Contact us"
	p.Form = form.Get[contactForm](ctx)

	return c.RenderPage(ctx, p)
}

func (c *Contact) Submit(ctx echo.Context) error {
	var input contactForm

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return c.Page(ctx)
	default:
		return err
	}

	err = c.mail.
		Compose().
		To(input.Email).
		Subject("Contact form submitted").
		Body(fmt.Sprintf("The message is: %s", input.Message)).
		Send(ctx)

	if err != nil {
		return fail(err, "unable to send email")
	}

	return c.Page(ctx)
}
