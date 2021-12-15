package auth

import (
	"errors"

	"goweb/config"
	"goweb/ent"
	"goweb/ent/user"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	sessionName             = "ua"
	sessionKeyUserID        = "user_id"
	sessionKeyAuthenticated = "authenticated"
)

type Client struct {
	config *config.Config
	orm    *ent.Client
}

func NewClient(cfg *config.Config, orm *ent.Client) *Client {
	return &Client{
		config: cfg,
		orm:    orm,
	}
}

func (c *Client) Login(ctx echo.Context, userID int) error {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return err
	}
	sess.Values[sessionKeyUserID] = userID
	sess.Values[sessionKeyAuthenticated] = true
	return sess.Save(ctx.Request(), ctx.Response())
}

func (c *Client) Logout(ctx echo.Context) error {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return err
	}
	sess.Values[sessionKeyAuthenticated] = false
	return sess.Save(ctx.Request(), ctx.Response())
}

func (c *Client) GetAuthenticatedUserID(ctx echo.Context) (int, error) {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return 0, err
	}

	if sess.Values[sessionKeyAuthenticated] == true {
		return sess.Values[sessionKeyUserID].(int), nil
	}

	return 0, errors.New("user not authenticated")
}

func (c *Client) GetAuthenticatedUser(ctx echo.Context) (*ent.User, error) {
	if userID, err := c.GetAuthenticatedUserID(ctx); err == nil {
		return c.orm.User.Query().
			Where(user.ID(userID)).
			First(ctx.Request().Context())
	}

	return nil, errors.New("user not authenticated")
}

func (c *Client) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (c *Client) CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
