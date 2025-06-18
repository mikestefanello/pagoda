package ui

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicFile(t *testing.T) {
	path := "abc.txt"
	got := PublicFile(path)
	expected := fmt.Sprintf("/%s/%s", "files", path)
	assert.Equal(t, expected, got)
}

func TestStaticFile(t *testing.T) {
	path := "abc.txt"
	got := StaticFile(path)
	expected := fmt.Sprintf("/%s/%s?v=%s", "static", path, cacheBuster)
	assert.Equal(t, expected, got)
}
