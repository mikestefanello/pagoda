package services

import (
	"fmt"

	"github.com/mikestefanello/pagoda/config"

	"github.com/labstack/echo/v4"
)

type (
	// MailClient provides a client for sending email
	// This is purposely not completed because there are many different methods and services
	// for sending email, many of which are very different. Choose what works best for you
	// and populate the methods below
	MailClient struct {
		// config stores application configuration
		config *config.Config

		// templates stores the template renderer
		templates *TemplateRenderer
	}

	mail struct {
		client       *MailClient
		from         string
		to           string
		subject      string
		body         string
		template     string
		templateData interface{}
	}
)

// NewMailClient creates a new MailClient
func NewMailClient(cfg *config.Config, templates *TemplateRenderer) (*MailClient, error) {
	return &MailClient{
		config:    cfg,
		templates: templates,
	}, nil
}

// Compose creates a new email
func (m *MailClient) Compose() *mail {
	return &mail{
		client: m,
		from:   m.config.Mail.FromAddress,
	}
}

// skipSend determines if mail sending should be skipped
func (m *MailClient) skipSend() bool {
	return m.config.App.Environment != config.EnvProduction
}

// From sets the email from address
func (m *mail) From(from string) *mail {
	m.from = from
	return m
}

// To sets the email address this email will be sent to
func (m *mail) To(to string) *mail {
	m.to = to
	return m
}

// Subject sets the subject line of the email
func (m *mail) Subject(subject string) *mail {
	m.subject = subject
	return m
}

// Body sets the body of the email
// This is not required and will be ignored if a template via Template()
func (m *mail) Body(body string) *mail {
	m.body = body
	return m
}

// Template sets the template to be used to produce the body of the email
// The template name should only include the filename without the extension or directory.
// The template must reside within the emails sub-directory.
// The funcmap will be automatically added to the template.
// Use TemplateData() to supply the data that will be passed in to the template.
func (m *mail) Template(template string) *mail {
	m.template = template
	return m
}

// TemplateData sets the data that will be passed to the template specified when calling Template()
func (m *mail) TemplateData(data interface{}) *mail {
	m.templateData = data
	return m
}

// Send attempts to send the email
func (m *mail) Send(ctx echo.Context) error {
	// Check if a template was supplied
	if m.template != "" {
		// Parse and execute template
		buf, err := m.client.templates.ParseAndExecute(
			"mail",
			m.template,
			m.template,
			[]string{fmt.Sprintf("emails/%s", m.template)},
			[]string{},
			m.templateData,
		)
		if err != nil {
			return err
		}

		m.body = buf.String()
	}

	// Check if mail sending should be skipped
	if m.client.skipSend() {
		ctx.Logger().Debugf("skipping email sent to: %s", m.to)
		return nil
	}

	// TODO: Finish based on your mail sender of choice!
	return nil
}
