package services

import (
	"io/ioutil"
	"testing"

	"github.com/mikestefanello/pagoda/config"

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
		Files("htmx", "pages/error").
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
	components, err := ioutil.ReadDir(c.TemplateRenderer.GetTemplatesPath() + "/components")
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
