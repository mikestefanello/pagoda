package ui

import (
	"fmt"
	"time"

	"github.com/mikestefanello/pagoda/config"
)

// AppName is the name of the application.
const AppName = "Pagoda"

var (
	// cacheBuster stores the current time as a cache buster for static files.
	cacheBuster = fmt.Sprint(time.Now().Unix())
)

// File generates a relative URL to a static file including a cache-buster query parameter.
func File(filepath string) string {
	return fmt.Sprintf("/%s/%s?v=%s", config.StaticPrefix, filepath, cacheBuster)
}
