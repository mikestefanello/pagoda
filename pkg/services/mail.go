package services

import (
	"bytes"
	"errors"

	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/log"
	"maragu.dev/gomponents"

	"github.com/labstack/echo/v4"
)

type (
	// MailClient provides a client for sending email
	// This is purposely not completed because there are many different methods and services
	// for sending email, many of which are very different. Choose what works best for you
	// and populate the methods below. For now, emails will just be logged.
	MailClient struct {
		// config stores application configuration.
		config *config.Config
	}

	// mail represents an email to be sent.
	mail struct {
		client    *MailClient
		from      string
		to        string
		subject   string
		body      string
		component gomponents.Node
	}
)

// NewMailClient creates a new MailClient.
func NewMailClient(cfg *config.Config) (*MailClient, error) {
	return &MailClient{
		config: cfg,
	}, nil
}

// Compose creates a new email.
func (m *MailClient) Compose() *mail {
	return &mail{
		client: m,
		from:   m.config.Mail.FromAddress,
	}
}

// skipSend determines if mail sending should be skipped.
func (m *MailClient) skipSend() bool {
	return m.config.App.Environment != config.EnvProduction
}

// send attempts to send the email.
func (m *MailClient) send(email *mail, ctx echo.Context) error {
	switch {
	case email.to == "":
		return errors.New("email cannot be sent without a to address")
	case email.body == "" && email.component == nil:
		return errors.New("email cannot be sent without a body or component to render")
	}

	// Check if a component was supplied.
	if email.component != nil {
		// Render the component and use as the body.
		// TODO pool the buffers?
		buf := bytes.NewBuffer(nil)
		if err := email.component.Render(buf); err != nil {
			return err
		}

		email.body = buf.String()
	}

	// Check if mail sending should be skipped.
	if m.skipSend() {
		log.Ctx(ctx).Debug("skipping email delivery",
			"to", email.to,
		)
		return nil
	}

	// TODO: Finish based on your mail sender of choice or stop logging below!
	log.Ctx(ctx).Info("sending email",
		"to", email.to,
		"subject", email.subject,
		"body", email.body,
	)
	return nil
}

// From sets the email from address.
func (m *mail) From(from string) *mail {
	m.from = from
	return m
}

// To sets the email address this email will be sent to.
func (m *mail) To(to string) *mail {
	m.to = to
	return m
}

// Subject sets the subject line of the email.
func (m *mail) Subject(subject string) *mail {
	m.subject = subject
	return m
}

// Body sets the body of the email.
// This is not required and will be ignored if a component is set via Component().
func (m *mail) Body(body string) *mail {
	m.body = body
	return m
}

// Component sets a renderable component to use as the body of the email.
func (m *mail) Component(component gomponents.Node) *mail {
	m.component = component
	return m
}

// Send attempts to send the email.
func (m *mail) Send(ctx echo.Context) error {
	return m.client.send(m, ctx)
}
