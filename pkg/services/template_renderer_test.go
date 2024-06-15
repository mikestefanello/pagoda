package services

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/htmx"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/mikestefanello/pagoda/templates"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplateRenderer(t *testing.T) {
	group := "test"
	id := "parse"

	// Should not exist yet
	_, err := c.TemplateRenderer.Load(group, id)
	assert.Error(t, err)

	// Parse in to the cache
	tpl, err := c.TemplateRenderer.
		Parse().
		Group(group).
		Key(id).
		Base("htmx").
		Files("layouts/htmx", "pages/error").
		Directories("components").
		Store()
	require.NoError(t, err)

	// Should exist now
	parsed, err := c.TemplateRenderer.Load(group, id)
	require.NoError(t, err)

	// Check that all expected templates are included
	expectedTemplates := make(map[string]bool)
	expectedTemplates["htmx"+config.TemplateExt] = true
	expectedTemplates["error"+config.TemplateExt] = true
	components, err := templates.Get().ReadDir("components")
	require.NoError(t, err)
	for _, f := range components {
		expectedTemplates[f.Name()] = true
	}
	for _, v := range parsed.Template.Templates() {
		delete(expectedTemplates, v.Name())
	}
	assert.Empty(t, expectedTemplates)

	data := struct {
		StatusCode int
	}{
		StatusCode: 500,
	}
	buf, err := tpl.Execute(data)
	require.NoError(t, err)
	require.NotNil(t, buf)
	assert.Contains(t, buf.String(), "Please try again")

	buf, err = c.TemplateRenderer.
		Parse().
		Group(group).
		Key(id).
		Base("htmx").
		Files("htmx", "pages/error").
		Directories("components").
		Execute(data)

	require.NoError(t, err)
	require.NotNil(t, buf)
	assert.Contains(t, buf.String(), "Please try again")
}

func TestTemplateRenderer_RenderPage(t *testing.T) {
	setup := func() (echo.Context, *httptest.ResponseRecorder, page.Page) {
		ctx, rec := tests.NewContext(c.Web, "/test/TestController_RenderPage")
		tests.InitSession(ctx)

		p := page.New(ctx)
		p.Name = "home"
		p.Layout = "main"
		p.Cache.Enabled = false
		p.Headers["A"] = "b"
		p.Headers["C"] = "d"
		p.StatusCode = http.StatusCreated
		return ctx, rec, p
	}

	t.Run("missing name", func(t *testing.T) {
		// Rendering should fail if the Page has no name
		ctx, _, p := setup()
		p.Name = ""
		err := c.TemplateRenderer.RenderPage(ctx, p)
		assert.Error(t, err)
	})

	t.Run("no page cache", func(t *testing.T) {
		ctx, _, p := setup()
		err := c.TemplateRenderer.RenderPage(ctx, p)
		require.NoError(t, err)

		// Check status code and headers
		assert.Equal(t, http.StatusCreated, ctx.Response().Status)
		for k, v := range p.Headers {
			assert.Equal(t, v, ctx.Response().Header().Get(k))
		}

		// Check the template cache
		parsed, err := c.TemplateRenderer.Load("page", string(p.Name))
		require.NoError(t, err)

		// Check that all expected templates were parsed.
		// This includes the name, layout and all components
		expectedTemplates := make(map[string]bool)
		expectedTemplates[fmt.Sprintf("%s%s", p.Name, config.TemplateExt)] = true
		expectedTemplates[fmt.Sprintf("%s%s", p.Layout, config.TemplateExt)] = true
		components, err := templates.Get().ReadDir("components")
		require.NoError(t, err)
		for _, f := range components {
			expectedTemplates[f.Name()] = true
		}

		for _, v := range parsed.Template.Templates() {
			delete(expectedTemplates, v.Name())
		}
		assert.Empty(t, expectedTemplates)
	})

	t.Run("htmx rendering", func(t *testing.T) {
		ctx, _, p := setup()
		p.HTMX.Request.Enabled = true
		p.HTMX.Response = &htmx.Response{
			Trigger: "trigger",
		}
		err := c.TemplateRenderer.RenderPage(ctx, p)
		require.NoError(t, err)

		// Check HTMX header
		assert.Equal(t, "trigger", ctx.Response().Header().Get(htmx.HeaderTrigger))

		// Check the template cache
		parsed, err := c.TemplateRenderer.Load("page:htmx", string(p.Name))
		require.NoError(t, err)

		// Check that all expected templates were parsed.
		// This includes the name, htmx and all components
		expectedTemplates := make(map[string]bool)
		expectedTemplates[fmt.Sprintf("%s%s", p.Name, config.TemplateExt)] = true
		expectedTemplates["htmx"+config.TemplateExt] = true
		components, err := templates.Get().ReadDir("components")
		require.NoError(t, err)
		for _, f := range components {
			expectedTemplates[f.Name()] = true
		}

		for _, v := range parsed.Template.Templates() {
			delete(expectedTemplates, v.Name())
		}
		assert.Empty(t, expectedTemplates)
	})

	t.Run("page cache", func(t *testing.T) {
		ctx, rec, p := setup()
		p.Cache.Enabled = true
		p.Cache.Tags = []string{"tag1"}
		err := c.TemplateRenderer.RenderPage(ctx, p)
		require.NoError(t, err)

		// Fetch from the cache
		cp, err := c.TemplateRenderer.GetCachedPage(ctx, p.URL)
		require.NoError(t, err)

		// Compare the cached page
		assert.Equal(t, p.URL, cp.URL)
		assert.Equal(t, p.Headers, cp.Headers)
		assert.Equal(t, p.StatusCode, cp.StatusCode)
		assert.Equal(t, rec.Body.Bytes(), cp.HTML)

		// Clear the tag
		err = c.Cache.
			Flush().
			Tags(p.Cache.Tags[0]).
			Execute(context.Background())
		require.NoError(t, err)

		// Refetch from the cache and expect no results
		_, err = c.TemplateRenderer.GetCachedPage(ctx, p.URL)
		assert.Error(t, err)
	})
}
