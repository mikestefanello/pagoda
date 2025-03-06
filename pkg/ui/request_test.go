package ui

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func TestNewRequest(t *testing.T) {
	e := echo.New()
	ctx, _ := tests.NewContext(e, "/")
	r := NewRequest(ctx)
	assert.Same(t, ctx, r.Context)
	assert.Equal(t, "/", r.CurrentPath)
	assert.True(t, r.IsHome)
	assert.False(t, r.IsAuth)
	assert.Nil(t, r.AuthUser)
	assert.Empty(t, r.CSRF)
	assert.Nil(t, r.Config)
	assert.Same(t, htmx.GetRequest(ctx), r.Htmx)

	ctx, _ = tests.NewContext(e, "/abc")
	usr := &ent.User{
		ID: 1,
	}
	ctx.Set(context.AuthenticatedUserKey, usr)
	ctx.Set(context.CSRFKey, "12345")
	ctx.Set(context.ConfigKey, &config.Config{
		App: config.AppConfig{
			Name: "testing",
		},
	})
	r = NewRequest(ctx)
	assert.Equal(t, "/abc", r.CurrentPath)
	assert.False(t, r.IsHome)
	assert.True(t, r.IsAuth)
	assert.Equal(t, usr, r.AuthUser)
	assert.Equal(t, "12345", r.CSRF)
	assert.Equal(t, "testing", r.Config.App.Name)
}

func TestRequest_UrlPath(t *testing.T) {
	e := echo.New()
	e.GET("/abc/:id", func(c echo.Context) error { return nil }).Name = "test"
	ctx, _ := tests.NewContext(e, "/")
	r := NewRequest(ctx)
	r.Config = &config.Config{
		App: config.AppConfig{
			Host: "http://localhost",
		},
	}

	assert.Equal(t, "http://localhost/abc/123", r.Url("test", 123))
	assert.Equal(t, "/abc/123", r.Path("test", 123))
}

func TestRequest_Render(t *testing.T) {
	e := echo.New()
	layout := func(r *Request, n gomponents.Node) gomponents.Node {
		return html.Div(html.Class("test"), n)
	}
	node := html.P(gomponents.Text("hello"))

	t.Run("no htmx", func(t *testing.T) {
		ctx, rec := tests.NewContext(e, "/")
		r := NewRequest(ctx)
		r.Htmx = &htmx.Request{}
		err := r.Render(layout, node)
		require.NoError(t, err)
		assert.Equal(t, `<div class="test"><p>hello</p></div>`, rec.Body.String())
	})

	t.Run("htmx", func(t *testing.T) {
		ctx, rec := tests.NewContext(e, "/")
		r := NewRequest(ctx)
		r.Htmx = &htmx.Request{
			Enabled: true,
			Boosted: false,
		}
		err := r.Render(layout, node)
		require.NoError(t, err)
		assert.Equal(t, `<p>hello</p>`, rec.Body.String())
	})
}
