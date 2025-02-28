package ui

import (
	"fmt"

	"github.com/labstack/gommon/random"
	"github.com/mikestefanello/pagoda/config"
)

// AppName is the name of the application.
const AppName = "Pagoda"

var (
	// cacheBuster stores a random string used as a cache buster for static files.
	cacheBuster = random.String(10)
)

// File generates a relative URL to a static file including a cache-buster query parameter.
func File(filepath string) string {
	return fmt.Sprintf("/%s/%s?v=%s", config.StaticPrefix, filepath, cacheBuster)
}
