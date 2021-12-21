package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"goweb/ent"

	"k8s.io/apimachinery/pkg/util/rand"

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
	handler := mw(echo.NotFoundHandler)
	return handler(ctx)
}

func CreateUser(orm *ent.Client) (*ent.User, error) {
	seed := fmt.Sprintf("%d-%d", time.Now().UnixMilli(), rand.IntnRange(10, 1000000))
	return orm.User.
		Create().
		SetEmail(fmt.Sprintf("testuser-%s@localhost.localhost", seed)).
		SetPassword("password").
		SetName(fmt.Sprintf("Test User %s", seed)).
		Save(context.Background())
}
