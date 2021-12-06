package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbout_Get(t *testing.T) {
	doc := request(t).
		setRoute("about").
		get().
		assertStatusCode(http.StatusOK).
		toDoc()

	h1 := doc.Find("h1.title")
	assert.Len(t, h1.Nodes, 1)
	assert.Equal(t, "About", h1.Text())
}
