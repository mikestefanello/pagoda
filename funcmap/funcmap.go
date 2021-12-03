package funcmap

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/labstack/gommon/random"
)

// CacheKey stores a random string used as a cache key for static files
var CacheKey = random.String(10)

func GetFuncMap() template.FuncMap {
	// See http://masterminds.github.io/sprig/ for available funcs
	funcMap := sprig.FuncMap()

	// Provide a list of custom functions
	f := template.FuncMap{
		"hasField": HasField,
		"file":     File,
		"link":     Link,
	}

	for k, v := range f {
		funcMap[k] = v
	}

	return funcMap
}

// HasField checks if an interface contains a given field
func HasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

// File appends a cache key to a given filepath so it can remain cached until the app is restarted
func File(filepath string) string {
	return fmt.Sprintf("%s?v=%s", filepath, CacheKey)
}

func Link(url, text, currentPath string, classes ...string) template.HTML {
	if currentPath == url {
		classes = append(classes, "is-active")
	}

	html := fmt.Sprintf(`<a class="%s" href="%s">%s</a>`, strings.Join(classes, " "), url, text)
	return template.HTML(html)
}
