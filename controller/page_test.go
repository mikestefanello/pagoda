package controller

import (
	"net/http"
	"testing"

	"goweb/context"
	"goweb/msg"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestNewPage(t *testing.T) {
	ctx := newContext("/")
	p := NewPage(ctx)
	assert.Same(t, ctx, p.Context)
	assert.NotNil(t, p.ToURL)
	assert.Equal(t, "/", p.Path)
	assert.Equal(t, "/", p.URL)
	assert.Equal(t, http.StatusOK, p.StatusCode)
	assert.Equal(t, NewPager(ctx, DefaultItemsPerPage), p.Pager)
	assert.Empty(t, p.Headers)
	assert.True(t, p.IsHome)
	assert.False(t, p.IsAuth)
	assert.Empty(t, p.CSRF)
	assert.Empty(t, p.RequestID)
	assert.False(t, p.Cache.Enabled)

	ctx = newContext("/abc?def=123")
	ctx.Set(context.AuthenticatedUserKey, 1)
	ctx.Set(echomw.DefaultCSRFConfig.ContextKey, "csrf")
	p = NewPage(ctx)
	assert.Equal(t, "/abc", p.Path)
	assert.Equal(t, "/abc?def=123", p.URL)
	assert.False(t, p.IsHome)
	assert.True(t, p.IsAuth)
	assert.Equal(t, "csrf", p.CSRF)
}

func TestPage_GetMessages(t *testing.T) {
	ctx := newContext("/")
	p := NewPage(ctx)

	// Simulate an HTTP request through the session middleware to initiate the session
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	handler := mw(echo.NotFoundHandler)
	assert.Error(t, handler(ctx))

	// Set messages
	msgTests := make(map[msg.Type][]string)
	msgTests[msg.TypeWarning] = []string{
		"abc",
		"def",
	}
	msgTests[msg.TypeInfo] = []string{
		"123",
		"456",
	}
	for typ, values := range msgTests {
		for _, value := range values {
			msg.Set(ctx, typ, value)
		}
	}

	// Get the messages
	for typ, values := range msgTests {
		msgs := p.GetMessages(typ)

		for i, message := range msgs {
			assert.Equal(t, values[i], string(message))
		}
	}
}
