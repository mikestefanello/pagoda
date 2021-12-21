package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"goweb/ent/passwordtoken"
	"goweb/ent/user"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
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

func TestPasswordHashing(t *testing.T) {
	pw := "testcheckpassword"
	hash, err := c.Auth.HashPassword(pw)
	assert.NoError(t, err)
	assert.NotEqual(t, hash, pw)
	err = c.Auth.CheckPassword(pw, hash)
	assert.NoError(t, err)
}

func TestGeneratePasswordResetToken(t *testing.T) {
	token, pt, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
	require.NoError(t, err)
	assert.Len(t, token, c.Config.App.PasswordToken.Length)
	assert.NoError(t, c.Auth.CheckPassword(token, pt.Hash))
}

func TestGetValidPasswordToken(t *testing.T) {
	// Check that a fake token is not valid
	_, err := c.Auth.GetValidPasswordToken(ctx, "faketoken", usr.ID)
	assert.Error(t, err)

	// Generate a valid token and check that it is returned
	token, pt, err := c.Auth.GeneratePasswordResetToken(ctx, usr.ID)
	require.NoError(t, err)
	pt2, err := c.Auth.GetValidPasswordToken(ctx, token, usr.ID)
	require.NoError(t, err)
	assert.Equal(t, pt.ID, pt2.ID)

	// Expire the token by pushed the date far enough back
	_, err = c.ORM.PasswordToken.
		Update().
		SetCreatedAt(time.Now().Add(-(c.Config.App.PasswordToken.Expiration + 10))).
		Where(passwordtoken.ID(pt.ID)).
		Save(context.Background())
	require.NoError(t, err)

	// Expired tokens should not be valid
	_, err = c.Auth.GetValidPasswordToken(ctx, token, usr.ID)
	assert.Error(t, err)
}

func TestDeletePasswordTokens(t *testing.T) {
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

func TestRandomToken(t *testing.T) {
	length := 64
	a, err := c.Auth.RandomToken(length)
	require.NoError(t, err)
	b, err := c.Auth.RandomToken(length)
	require.NoError(t, err)
	assert.Len(t, a, 64)
	assert.Len(t, b, 64)
	assert.NotEqual(t, a, b)
}
