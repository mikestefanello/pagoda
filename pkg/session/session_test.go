package session

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetStore(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, httptest.NewRecorder())
	_, err := Get(ctx, "test")
	assert.Equal(t, ErrStoreNotFound, err)

	Store(ctx, sessions.NewCookieStore([]byte("secret")))
	_, err = Get(ctx, "test")
	assert.NoError(t, err)
}
