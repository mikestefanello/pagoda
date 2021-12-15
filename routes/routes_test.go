package routes

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"goweb/config"
	"goweb/container"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

var (
	srv *httptest.Server
	c   *container.Container
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Start a test HTTP server
	c = container.NewContainer()
	BuildRouter(c)
	srv = httptest.NewServer(c.Web)

	// Run tests
	exitVal := m.Run()
	srv.Close()
	os.Exit(exitVal)
}

type httpRequest struct {
	route  string
	client http.Client
	body   url.Values
	t      *testing.T
}

func request(t *testing.T) *httpRequest {
	r := httpRequest{
		t: t,
	}
	return &r
}

func (h *httpRequest) setClient(client http.Client) *httpRequest {
	h.client = client
	return h
}

func (h *httpRequest) setRoute(route string, params ...interface{}) *httpRequest {
	h.route = srv.URL + c.Web.Reverse(route, params)
	return h
}

func (h *httpRequest) setBody(body url.Values) *httpRequest {
	h.body = body
	return h
}

func (h *httpRequest) get() *httpResponse {
	resp, err := h.client.Get(h.route)
	require.NoError(h.t, err)
	r := httpResponse{
		t:        h.t,
		Response: resp,
	}
	return &r
}

func (h *httpRequest) post() *httpResponse {
	resp, err := h.client.PostForm(h.route, h.body)
	require.NoError(h.t, err)
	r := httpResponse{
		t:        h.t,
		Response: resp,
	}
	return &r
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
