package controller

import (
	"net/http"
	"testing"

	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/tests"

	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPage(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
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

	ctx, _ = tests.NewContext(c.Web, "/abc?def=123")
	usr, err := tests.CreateUser(c.ORM)
	require.NoError(t, err)
	ctx.Set(context.AuthenticatedUserKey, usr)
	ctx.Set(echomw.DefaultCSRFConfig.ContextKey, "csrf")
	p = NewPage(ctx)
	assert.Equal(t, "/abc", p.Path)
	assert.Equal(t, "/abc?def=123", p.URL)
	assert.False(t, p.IsHome)
	assert.True(t, p.IsAuth)
	assert.Equal(t, usr, p.AuthUser)
	assert.Equal(t, "csrf", p.CSRF)
}

func TestPage_GetMessages(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	tests.InitSession(ctx)
	p := NewPage(ctx)

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
