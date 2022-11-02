package middleware

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mikestefanello/pagoda/pkg/tests"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"

	echomw "github.com/labstack/echo/v4/middleware"
)

func TestLogRequestID(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	_ = tests.ExecuteMiddleware(ctx, echomw.RequestID())
	_ = tests.ExecuteMiddleware(ctx, LogRequestID())

	var buf bytes.Buffer
	ctx.Logger().SetOutput(&buf)
	ctx.Logger().Info("test")
	rID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	assert.Contains(t, buf.String(), fmt.Sprintf(`id":"%s"`, rID))
}
