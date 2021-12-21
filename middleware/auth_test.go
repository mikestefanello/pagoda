package middleware

import (
	"net/http"
	"testing"

	"goweb/context"
	"goweb/ent"
	"goweb/tests"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestLoadAuthenticatedUser(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)
	mw := LoadAuthenticatedUser(c.Auth)

	// Not authenticated
	_ = tests.ExecuteMiddleware(ctx, mw)
	assert.Nil(t, ctx.Get(context.AuthenticatedUserKey))

	// Login
	err := c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)

	// Verify the midldeware returns the authenticated user
	_ = tests.ExecuteMiddleware(ctx, mw)
	require.NotNil(t, ctx.Get(context.AuthenticatedUserKey))
	ctxUsr, ok := ctx.Get(context.AuthenticatedUserKey).(*ent.User)
	require.True(t, ok)
	assert.Equal(t, usr.ID, ctxUsr.ID)
}

func TestRequireAuthentication(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// Not logged in
	err := tests.ExecuteMiddleware(ctx, RequireAuthentication())
	httpError, ok := err.(*echo.HTTPError)
	require.True(t, ok)
	assert.Equal(t, http.StatusUnauthorized, httpError.Code)

	// Login
	err = c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)
	_ = tests.ExecuteMiddleware(ctx, LoadAuthenticatedUser(c.Auth))

	// Logged in
	err = tests.ExecuteMiddleware(ctx, RequireAuthentication())
	httpError, ok = err.(*echo.HTTPError)
	require.True(t, ok)
	assert.NotEqual(t, http.StatusUnauthorized, httpError.Code)
}

func TestRequireNoAuthentication(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// Not logged in
	err := tests.ExecuteMiddleware(ctx, RequireNoAuthentication())
	httpError, ok := err.(*echo.HTTPError)
	require.True(t, ok)
	assert.NotEqual(t, http.StatusForbidden, httpError.Code)

	// Login
	err = c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)
	_ = tests.ExecuteMiddleware(ctx, LoadAuthenticatedUser(c.Auth))

	// Logged in
	err = tests.ExecuteMiddleware(ctx, RequireNoAuthentication())
	httpError, ok = err.(*echo.HTTPError)
	require.True(t, ok)
	assert.Equal(t, http.StatusForbidden, httpError.Code)
}

func TestLoadValidPasswordToken(t *testing.T) {

}
