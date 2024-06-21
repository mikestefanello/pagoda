package handlers

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/services"
	"github.com/mikestefanello/pagoda/pkg/tasks"
	"github.com/mikestefanello/pagoda/templates"
)

const (
	routeNameContact       = "contact"
	routeNameContactSubmit = "contact.submit"
)

type (
	Contact struct {
		mail  *services.MailClient
		tasks *services.TaskClient
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

func (h *Contact) Init(c *services.Container) error {
	h.TemplateRenderer = c.TemplateRenderer
	h.mail = c.Mail
	h.tasks = c.Tasks
	return nil
}

func (h *Contact) Routes(g *echo.Group) {
	g.GET("/contact", h.Page).Name = routeNameContact
	g.POST("/contact", h.Submit).Name = routeNameContactSubmit
}

func (h *Contact) Page(ctx echo.Context) error {
	p := page.New(ctx)
	p.Layout = templates.LayoutMain
	p.Name = templates.PageContact
	p.Title = "Contact us"
	p.Form = form.Get[contactForm](ctx)

	return h.RenderPage(ctx, p)
}

func (h *Contact) Submit(ctx echo.Context) error {
	var input contactForm

	err := form.Submit(ctx, &input)

	switch err.(type) {
	case nil:
	case validator.ValidationErrors:
		return h.Page(ctx)
	default:
		return err
	}

	// TODO create a new page for this
	err = h.tasks.New(tasks.ExampleTask{
		Message: input.Message,
	}).
		Wait(10 * time.Second).
		Save()
	if err != nil {
		return err
	}

	err = h.mail.
		Compose().
		To(input.Email).
		Subject("Contact form submitted").
		Body(fmt.Sprintf("The message is: %s", input.Message)).
		Send(ctx)

	if err != nil {
		return fail(err, "unable to send email")
	}

	return h.Page(ctx)
}
