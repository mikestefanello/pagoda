package ui

import (
	"fmt"
	"time"
)

var (
	// cacheBuster stores the current time as a cache buster for static files.
	cacheBuster = fmt.Sprint(time.Now().Unix())
)

// PublicFile generates a relative URL to a public file including a cache-buster query parameter.
func PublicFile(filepath string) string {
	return file("files", filepath)
}

// StaticFile generates a relative URL to a static file including a cache-buster query parameter.
func StaticFile(filepath string) string {
	return file("static", filepath)
}

func file(dir, filepath string) string {
	return fmt.Sprintf("/%s/%s?v=%s", dir, filepath, cacheBuster)
}
