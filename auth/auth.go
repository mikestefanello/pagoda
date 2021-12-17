package auth

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"goweb/config"
	"goweb/ent"
	"goweb/ent/passwordtoken"
	"goweb/ent/user"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	sessionName             = "ua"
	sessionKeyUserID        = "user_id"
	sessionKeyAuthenticated = "authenticated"
	passwordTokenLength     = 64
)

type NotAuthenticatedError struct{}

// Error implements the error interface.
func (e NotAuthenticatedError) Error() string {
	return "user not authenticated"
}

type InvalidTokenError struct{}

// Error implements the error interface.
func (e InvalidTokenError) Error() string {
	return "invalid token"
}

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

	return 0, NotAuthenticatedError{}
}

func (c *Client) GetAuthenticatedUser(ctx echo.Context) (*ent.User, error) {
	if userID, err := c.GetAuthenticatedUserID(ctx); err == nil {
		return c.orm.User.Query().
			Where(user.ID(userID)).
			Only(ctx.Request().Context())
	}

	return nil, NotAuthenticatedError{}
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

func (c *Client) GeneratePasswordResetToken(ctx echo.Context, userID int) (string, *ent.PasswordToken, error) {
	// Generate the token, which is what will go in the URL, but not the database
	token, err := c.RandomToken(passwordTokenLength)
	if err != nil {
		return "", nil, err
	}

	// Hash the token, which is what will be stored in the database
	hash, err := c.HashPassword(token)
	if err != nil {
		return "", nil, err
	}

	// Create and save the password reset token
	pt, err := c.orm.PasswordToken.
		Create().
		SetHash(hash).
		SetUserID(userID).
		Save(ctx.Request().Context())

	return token, pt, err
}

func (c *Client) GetValidPasswordToken(ctx echo.Context, token string, userID int) (*ent.PasswordToken, error) {
	// Ensure expired tokens are never returned
	expiration := time.Now().Add(-c.config.App.PasswordTokenExpiration)

	// Query to find all tokens for te user that haven't expired
	// We need to get all of them in order to properly match the token to the hashes
	pts, err := c.orm.PasswordToken.
		Query().
		Where(passwordtoken.HasUserWith(user.ID(userID))).
		Where(passwordtoken.CreatedAtGTE(expiration)).
		All(ctx.Request().Context())

	if err != nil {
		ctx.Logger().Error(err)
		return nil, err
	}

	// Check all tokens for a hash match
	for _, pt := range pts {
		if err := c.CheckPassword(token, pt.Hash); err == nil {
			return pt, nil
		}
	}

	return nil, InvalidTokenError{}
}

func (c *Client) DeletePasswordTokens(ctx echo.Context, userID int) error {
	_, err := c.orm.PasswordToken.
		Delete().
		Where(passwordtoken.HasUserWith(user.ID(userID))).
		Exec(ctx.Request().Context())

	return err
}

func (c *Client) RandomToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
