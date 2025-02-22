package ui

import (
	"fmt"

	"github.com/labstack/gommon/random"
	"github.com/mikestefanello/pagoda/config"
)

var (
	// CacheBuster stores a random string used as a cache buster for static files.
	cacheBuster = random.String(10)
)

func file(filepath string) string {
	return fmt.Sprintf("/%s/%s?v=%s", config.StaticPrefix, filepath, cacheBuster)
}
