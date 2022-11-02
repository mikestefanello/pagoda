package middleware

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/tests"

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
	tests.AssertHTTPErrorCode(t, err, http.StatusUnauthorized)

	// Login
	err = c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)
	_ = tests.ExecuteMiddleware(ctx, LoadAuthenticatedUser(c.Auth))

	// Logged in
	err = tests.ExecuteMiddleware(ctx, RequireAuthentication())
	assert.Nil(t, err)
}

func TestRequireNoAuthentication(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// Not logged in
	err := tests.ExecuteMiddleware(ctx, RequireNoAuthentication())
	assert.Nil(t, err)

	// Login
	err = c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)
	_ = tests.ExecuteMiddleware(ctx, LoadAuthenticatedUser(c.Auth))

	// Logged in
	err = tests.ExecuteMiddleware(ctx, RequireNoAuthentication())
	tests.AssertHTTPErrorCode(t, err, http.StatusForbidden)
}

func TestLoadValidPasswordToken(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)

	// Missing user context
	err := tests.ExecuteMiddleware(ctx, LoadValidPasswordToken(c.Auth))
	tests.AssertHTTPErrorCode(t, err, http.StatusInternalServerError)

	// Add user and password token context but no token and expect a redirect
	ctx.SetParamNames("user", "password_token")
	ctx.SetParamValues(fmt.Sprintf("%d", usr.ID), "1")
	_ = tests.ExecuteMiddleware(ctx, LoadUser(c.ORM))
	err = tests.ExecuteMiddleware(ctx, LoadValidPasswordToken(c.Auth))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, ctx.Response().Status)

	// Add user context and invalid password token and expect a redirect
	ctx.SetParamNames("user", "password_token", "token")
	ctx.SetParamValues(fmt.Sprintf("%d", usr.ID), "1", "faketoken")
	_ = tests.ExecuteMiddleware(ctx, LoadUser(c.ORM))
	err = tests.ExecuteMiddleware(ctx, LoadValidPasswordToken(c.Auth))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, ctx.Response().Status)

	// Create a valid token
	token, pt, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
	require.NoError(t, err)

	// Add user and valid password token
	ctx.SetParamNames("user", "password_token", "token")
	ctx.SetParamValues(fmt.Sprintf("%d", usr.ID), fmt.Sprintf("%d", pt.ID), token)
	_ = tests.ExecuteMiddleware(ctx, LoadUser(c.ORM))
	err = tests.ExecuteMiddleware(ctx, LoadValidPasswordToken(c.Auth))
	assert.Nil(t, err)
	ctxPt, ok := ctx.Get(context.PasswordTokenKey).(*ent.PasswordToken)
	require.True(t, ok)
	assert.Equal(t, pt.ID, ctxPt.ID)
}
