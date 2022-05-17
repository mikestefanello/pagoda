package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/user"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestAuthClient_Auth(t *testing.T) {
	assertNoAuth := func() {
		_, err := c.Auth.GetAuthenticatedUserID(ctx)
		assert.True(t, errors.Is(err, NotAuthenticatedError{}))
		_, err = c.Auth.GetAuthenticatedUser(ctx)
		assert.True(t, errors.Is(err, NotAuthenticatedError{}))
	}

	assertNoAuth()

	err := c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)

	uid, err := c.Auth.GetAuthenticatedUserID(ctx)
	require.NoError(t, err)
	assert.Equal(t, usr.ID, uid)

	u, err := c.Auth.GetAuthenticatedUser(ctx)
	require.NoError(t, err)
	assert.Equal(t, u.ID, usr.ID)

	err = c.Auth.Logout(ctx)
	require.NoError(t, err)

	assertNoAuth()
}

func TestAuthClient_PasswordHashing(t *testing.T) {
	pw := "testcheckpassword"
	hash, err := c.Auth.HashPassword(pw)
	assert.NoError(t, err)
	assert.NotEqual(t, hash, pw)
	err = c.Auth.CheckPassword(pw, hash)
	assert.NoError(t, err)
}

func TestAuthClient_GeneratePasswordResetToken(t *testing.T) {
	token, pt, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
	require.NoError(t, err)
	assert.Len(t, token, c.Config.App.PasswordToken.Length)
	assert.NoError(t, c.Auth.CheckPassword(token, pt.Hash))
}

func TestAuthClient_GetValidPasswordToken(t *testing.T) {
	// Check that a fake token is not valid
	_, err := c.Auth.GetValidPasswordToken(ctx, usr.ID, 1, "faketoken")
	assert.Error(t, err)

	// Generate a valid token and check that it is returned
	token, pt, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
	require.NoError(t, err)
	pt2, err := c.Auth.GetValidPasswordToken(ctx, usr.ID, pt.ID, token)
	require.NoError(t, err)
	assert.Equal(t, pt.ID, pt2.ID)

	// Expire the token by pushing the date far enough back
	count, err := c.ORM.PasswordToken.
		Update().
		SetCreatedAt(time.Now().Add(-(c.Config.App.PasswordToken.Expiration + time.Hour))).
		Where(passwordtoken.ID(pt.ID)).
		Save(context.Background())
	require.NoError(t, err)
	require.Equal(t, 1, count)

	// Expired tokens should not be valid
	_, err = c.Auth.GetValidPasswordToken(ctx, usr.ID, pt.ID, token)
	assert.Error(t, err)
}

func TestAuthClient_DeletePasswordTokens(t *testing.T) {
	// Create three tokens for the user
	for i := 0; i < 3; i++ {
		_, _, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
		require.NoError(t, err)
	}

	// Delete all tokens for the user
	err := c.Auth.DeletePasswordTokens(ctx, usr.ID)
	require.NoError(t, err)

	// Check that no tokens remain
	count, err := c.ORM.PasswordToken.
		Query().
		Where(passwordtoken.HasUserWith(user.ID(usr.ID))).
		Count(context.Background())

	require.NoError(t, err)
	assert.Equal(t, 0, count)
}

func TestAuthClient_RandomToken(t *testing.T) {
	length := c.Config.App.PasswordToken.Length
	a, err := c.Auth.RandomToken(length)
	require.NoError(t, err)
	b, err := c.Auth.RandomToken(length)
	require.NoError(t, err)
	assert.Len(t, a, length)
	assert.Len(t, b, length)
	assert.NotEqual(t, a, b)
}

func TestAuthClient_EmailVerificationToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		email := "test@localhost.com"
		token, err := c.Auth.GenerateEmailVerificationToken(email)
		require.NoError(t, err)

		tokenEmail, err := c.Auth.ValidateEmailVerificationToken(token)
		require.NoError(t, err)
		assert.Equal(t, email, tokenEmail)
	})

	t.Run("invalid token", func(t *testing.T) {
		badToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbG9jYWxob3N0LmNvbSIsImV4cCI6MTkxNzg2NDAwMH0.ScJCpfEEzlilKfRs_aVouzwPNKI28M3AIm-hyImQHUQ"
		_, err := c.Auth.ValidateEmailVerificationToken(badToken)
		assert.Error(t, err)
	})

	t.Run("expired token", func(t *testing.T) {
		c.Config.App.EmailVerificationTokenExpiration = -time.Hour
		email := "test@localhost.com"
		token, err := c.Auth.GenerateEmailVerificationToken(email)
		require.NoError(t, err)

		_, err = c.Auth.ValidateEmailVerificationToken(token)
		assert.Error(t, err)

		c.Config.App.EmailVerificationTokenExpiration = time.Hour * 12
	})
}
