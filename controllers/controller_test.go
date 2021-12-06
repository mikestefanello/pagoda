package controllers

import (
	"net/http"
	"net/http/httptest"
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

func GetRequest(t *testing.T, route string, routeParams ...interface{}) *http.Response {
	cli := http.Client{}
	resp, err := cli.Get(srv.URL + c.Web.Reverse(route, routeParams))
	require.NoError(t, err)
	return resp
}

func GetGoqueryDoc(t *testing.T, resp *http.Response) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	require.NoError(t, err)
	err = resp.Body.Close()
	assert.NoError(t, err)
	return doc
}
