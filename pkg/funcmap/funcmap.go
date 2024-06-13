package funcmap

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
	"github.com/mikestefanello/pagoda/config"
)

var (
	// CacheBuster stores a random string used as a cache buster for static files.
	CacheBuster = random.String(10)
)

type funcMap struct {
	web *echo.Echo
}

// GetFuncMap provides a template function map
func GetFuncMap(web *echo.Echo) template.FuncMap {
	fm := &funcMap{web: web}

	// See http://masterminds.github.io/sprig/ for all provided funcs
	funcs := sprig.FuncMap()

	// Include all the custom functions
	funcs["hasField"] = fm.HasField
	funcs["file"] = fm.File
	funcs["link"] = fm.Link
	funcs["url"] = fm.url

	return funcs
}

// HasField checks if an interface contains a given field
func (fm *funcMap) HasField(v any, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

// File appends a cache buster to a given filepath so it can remain cached until the app is restarted
func (fm *funcMap) File(filepath string) string {
	return fmt.Sprintf("/%s/%s?v=%s", config.StaticPrefix, filepath, CacheBuster)
}

// Link outputs HTML for a link element, providing the ability to dynamically set the active class
func (fm *funcMap) Link(url, text, currentPath string, classes ...string) template.HTML {
	if currentPath == url {
		classes = append(classes, "is-active")
	}

	html := fmt.Sprintf(`<a class="%s" href="%s">%s</a>`, strings.Join(classes, " "), url, text)
	return template.HTML(html)
}

// Url generates a URL from a given route name and optional parameters
func (fm *funcMap) url(routeName string, params ...any) string {
	return fm.web.Reverse(routeName, params)
}
