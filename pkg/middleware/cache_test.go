package middleware

import (
	"net/http"
	"testing"
	"time"

	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/mikestefanello/pagoda/templates"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestServeCachedPage(t *testing.T) {
	// Cache a page
	ctx, rec := tests.NewContext(c.Web, "/cache")
	p := page.New(ctx)
	p.Layout = templates.LayoutHTMX
	p.Name = templates.PageHome
	p.Cache.Enabled = true
	p.Cache.Expiration = time.Minute
	p.StatusCode = http.StatusCreated
	p.Headers["a"] = "b"
	p.Headers["c"] = "d"
	err := c.TemplateRenderer.RenderPage(ctx, p)
	output := rec.Body.Bytes()
	require.NoError(t, err)

	// Request the URL of the cached page
	ctx, rec = tests.NewContext(c.Web, "/cache")
	err = tests.ExecuteMiddleware(ctx, ServeCachedPage(c.TemplateRenderer))
	assert.NoError(t, err)
	assert.Equal(t, p.StatusCode, ctx.Response().Status)
	assert.Equal(t, p.Headers["a"], ctx.Response().Header().Get("a"))
	assert.Equal(t, p.Headers["c"], ctx.Response().Header().Get("c"))
	assert.Equal(t, output, rec.Body.Bytes())

	// Login and try again
	tests.InitSession(ctx)
	err = c.Auth.Login(ctx, usr.ID)
	require.NoError(t, err)
	_ = tests.ExecuteMiddleware(ctx, LoadAuthenticatedUser(c.Auth))
	err = tests.ExecuteMiddleware(ctx, ServeCachedPage(c.TemplateRenderer))
	assert.Nil(t, err)
}

func TestCacheControl(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	_ = tests.ExecuteMiddleware(ctx, CacheControl(time.Second*5))
	assert.Equal(t, "public, max-age=5", ctx.Response().Header().Get("Cache-Control"))
	_ = tests.ExecuteMiddleware(ctx, CacheControl(0))
	assert.Equal(t, "no-cache, no-store", ctx.Response().Header().Get("Cache-Control"))
}
