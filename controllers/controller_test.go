package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"goweb/container"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/require"
)

var (
	srv *httptest.Server
	c   *container.Container
)

func TestMain(m *testing.M) {
	// Start a test HTTP server
	c = container.NewContainer()
	BuildRouter(c)
	c.Web.Logger.SetLevel(log.DEBUG)
	srv = httptest.NewServer(c.Web)

	exitVal := m.Run()
	srv.Close()
	os.Exit(exitVal)
}

type httpResponse struct {
	*http.Response
	t *testing.T
}

func (h *httpResponse) assertStatusCode(code int) *httpResponse {
	assert.Equal(h.t, code, h.Response.StatusCode)
	return h
}

func (h *httpResponse) assertRedirect(t *testing.T, destination string) *httpResponse {
	assert.Equal(t, destination, h.Header.Get("Location"))
	return h
}

func (h *httpResponse) toDoc() *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(h.Body)
	require.NoError(h.t, err)
	err = h.Body.Close()
	assert.NoError(h.t, err)
	return doc
}

func getRequest(t *testing.T, route string, routeParams ...interface{}) *httpResponse {
	cli := http.Client{}
	resp, err := cli.Get(srv.URL + c.Web.Reverse(route, routeParams))
	require.NoError(t, err)
	h := httpResponse{
		t:        t,
		Response: resp,
	}
	return &h
}

func postRequest(t *testing.T, values url.Values, route string, routeParams ...interface{}) *httpResponse {
	cli := http.Client{}
	resp, err := cli.PostForm(srv.URL+c.Web.Reverse(route, routeParams), values)
	require.NoError(t, err)
	h := httpResponse{
		t:        t,
		Response: resp,
	}
	return &h
}
