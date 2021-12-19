package services

import (
	"goweb/config"

	"github.com/labstack/echo/v4"
)

// MailClient provides a client for sending email
// This is purposely not completed because there are many different methods and services
// for sending email, many of which are very different. Choose what works best for you
// and populate the methods below
type MailClient struct {
	// config stores application configuration
	config *config.Config
}

// NewMailClient creates a new MailClient
func NewMailClient(cfg *config.Config) (*MailClient, error) {
	return &MailClient{
		config: cfg,
	}, nil
}

// Send sends an email to a given email address with a given body
func (c *MailClient) Send(ctx echo.Context, to, body string) error {
	if c.skipSend() {
		ctx.Logger().Debugf("skipping email sent to: %s")
	}

	// TODO: Finish based on your mail sender of choice
	return nil
}

// SendTemplate sends an email to a given email address using a template and data which is passed to the template
// The template name should only include the filename without the extension or directory.
// The funcmap will be automatically added to the template and the data will be passed in.
func (c *MailClient) SendTemplate(ctx echo.Context, to, template string, data interface{}) error {
	if c.skipSend() {
		ctx.Logger().Debugf("skipping template email sent to: %s")
	}

	// TODO: Finish based on your mail sender of choice
	return nil
}

// skipSend determines if mail sending should be skipped
func (c *MailClient) skipSend() bool {
	return c.config.App.Environment != config.EnvProduction
}
