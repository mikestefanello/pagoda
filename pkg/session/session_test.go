package session

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetStore(t *testing.T) {
	const sessName = "test"
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())

	// No store is set on the context yet.
	_, err := Get(ctx, sessName)
	assert.Equal(t, ErrStoreNotFound, err)

	// Set up the store.
	Store(ctx, sessions.NewCookieStore([]byte("secret")))
	sess, err := Get(ctx, sessName)
	require.NoError(t, err)

	// Save session data.
	sess.Values["a"] = 1
	err = sess.Save(ctx.Request(), ctx.Response())
	require.NoError(t, err)

	// Ensure the session data can be fetched.
	sess, err = Get(ctx, sessName)
	require.NoError(t, err)
	assert.Equal(t, sess.Values["a"], 1)

	// Create a new request with the session cookie.
	cookie := ctx.Response().Header().Get("Set-Cookie")
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("Cookie", cookie)
	ctx = e.NewContext(req, httptest.NewRecorder())

	// Change the encryption key and ensure no error is returned.
	Store(ctx, sessions.NewCookieStore([]byte("newSecret")))
	_, err = Get(ctx, sessName)
	require.NoError(t, err)

	// The session data shouldn't be present now.
	sess, err = Get(ctx, sessName)
	require.NoError(t, err)
	_, exists := sess.Values["a"]
	assert.False(t, exists)
}
