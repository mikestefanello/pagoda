package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbout_Get(t *testing.T) {
	resp := GetRequest(t, "about")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	doc := GetGoqueryDoc(t, resp)
	h1 := doc.Find("h1.title")
	assert.Len(t, h1.Nodes, 1)
	assert.Equal(t, "About", h1.Text())
}
