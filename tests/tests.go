package tests

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mikestefanello/pagoda/ent"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func NewContext(e *echo.Echo, url string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, url, strings.NewReader(""))
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func InitSession(ctx echo.Context) {
	mw := session.Middleware(sessions.NewCookieStore([]byte("secret")))
	_ = ExecuteMiddleware(ctx, mw)
}

func ExecuteMiddleware(ctx echo.Context, mw echo.MiddlewareFunc) error {
	handler := mw(func(c echo.Context) error {
		return nil
	})
	return handler(ctx)
}

func AssertHTTPErrorCode(t *testing.T, err error, code int) {
	httpError, ok := err.(*echo.HTTPError)
	require.True(t, ok)
	assert.Equal(t, code, httpError.Code)
}

func CreateUser(orm *ent.Client) (*ent.User, error) {
	seed := fmt.Sprintf("%d-%d", time.Now().UnixMilli(), rand.Intn(1000000))
	return orm.User.
		Create().
		SetEmail(fmt.Sprintf("testuser-%s@localhost.localhost", seed)).
		SetPassword("password").
		SetName(fmt.Sprintf("Test User %s", seed)).
		Save(context.Background())
}
