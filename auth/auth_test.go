package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"goweb/config"
	"goweb/container"
	"goweb/ent"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

var (
	authClient *Client
	c          *container.Container
	ctx        echo.Context
	usr        *ent.User
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create an auth client
	c := container.NewContainer()
	authClient = c.Auth

	// Create a web context
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()
	ctx = c.Web.NewContext(req, rec)

	// Create a test uset
	var err error
	usr, err = c.ORM.User.
		Create().
		SetEmail("test@test.dev").
		SetPassword("abc").
		SetName("Test User").
		Save(context.Background())

	if err != nil {
		panic(err)
	}

	// Run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestLogin(t *testing.T) {

}

func TestLogout(t *testing.T) {

}

func TestGetAuthenticatedUserID(t *testing.T) {

}

func TestGetAuthenticatedUser(t *testing.T) {

}

func TestHashPassword(t *testing.T) {
	pw := "abcdef"
	hash, err := authClient.HashPassword(pw)
	assert.NoError(t, err)
	assert.NotEqual(t, hash, pw)
}

func TestCheckPassword(t *testing.T) {
	pw := "testcheckpassword"
	hash, err := authClient.HashPassword(pw)
	assert.NoError(t, err)
	err = authClient.CheckPassword(pw, hash)
	assert.NoError(t, err)
}

func TestGeneratePasswordResetToken(t *testing.T) {
	token, pt, err := authClient.GeneratePasswordResetToken(ctx, usr.ID)
	require.NoError(t, err)
	hash, err := authClient.HashPassword(token)
	require.NoError(t, err)
	assert.Len(t, token, c.Config.App.PasswordToken.Length)
	assert.Equal(t, hash, pt.Hash)
	assert.Equal(t, usr.ID, pt.Edges.User.ID)
}

func TestGetValidPasswordToken(t *testing.T) {

}

func TestDeletePasswordTokens(t *testing.T) {
}

func TestRandomToken(t *testing.T) {
	length := 64
	a, err := authClient.RandomToken(length)
	require.NoError(t, err)
	b, err := authClient.RandomToken(length)
	require.NoError(t, err)
	assert.Len(t, a, 64)
	assert.Len(t, b, 64)
	assert.NotEqual(t, a, b)
}
