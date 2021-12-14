package mail

import (
	"goweb/config"

	"github.com/labstack/echo/v4"
)

type Client struct {
	config *config.Config
}

func NewClient(cfg *config.Config) (*Client, error) {
	return &Client{
		config: cfg,
	}, nil
}

func (c *Client) Send(ctx echo.Context, to, body string) error {
	if c.config.App.Environment != config.EnvProduction {
		// IE, skip sending email..
	}
	ctx.Logger().Debugf("Mock email sent. To: %s  Body: %s", to, body)
	return nil
}

// TODO: Send with template?
