package funcmap

import (
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/config"

	"github.com/stretchr/testify/assert"
)

func TestNewFuncMap(t *testing.T) {
	f := NewFuncMap(echo.New())
	assert.NotNil(t, f["link"])
	assert.NotNil(t, f["file"])
	assert.NotNil(t, f["url"])
}

func TestLink(t *testing.T) {
	f := new(funcMap)

	link := string(f.link("/abc", "Text", "/abc"))
	expected := `<a class="is-active" href="/abc">Text</a>`
	assert.Equal(t, expected, link)

	link = string(f.link("/abc", "Text", "/abc", "first", "second"))
	expected = `<a class="first second is-active" href="/abc">Text</a>`
	assert.Equal(t, expected, link)

	link = string(f.link("/abc", "Text", "/def"))
	expected = `<a class="" href="/abc">Text</a>`
	assert.Equal(t, expected, link)
}

func TestFile(t *testing.T) {
	f := new(funcMap)

	file := f.file("test.png")
	expected := fmt.Sprintf("/%s/test.png?v=%s", config.StaticPrefix, CacheBuster)
	assert.Equal(t, expected, file)
}

func TestUrl(t *testing.T) {
	f := new(funcMap)
	f.web = echo.New()
	f.web.GET("/mypath/:id", func(c echo.Context) error {
		return nil
	}).Name = "test"
	out := f.url("test", 5)
	assert.Equal(t, "/mypath/5", out)
}
