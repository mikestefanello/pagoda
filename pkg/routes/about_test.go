package routes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple example of how to test routes and their markup using the test HTTP server spun up within
// this test package
func TestAbout_Get(t *testing.T) {
	doc := request(t).
		setRoute("about").
		get().
		assertStatusCode(http.StatusOK).
		toDoc()

	// Goquery is an excellent package to use for testing HTML markup
	h1 := doc.Find("h1.title")
	assert.Len(t, h1.Nodes, 1)
	assert.Equal(t, "About", h1.Text())
}
