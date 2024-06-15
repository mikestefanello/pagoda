package middleware

import (
	"context"
	"log/slog"
	"testing"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

type mockLogHandler struct {
	msg   string
	level string
	group string
	attr  []slog.Attr
}

func (m *mockLogHandler) Enabled(_ context.Context, l slog.Level) bool {
	return true
}

func (m *mockLogHandler) Handle(_ context.Context, r slog.Record) error {
	m.level = r.Level.String()
	m.msg = r.Message
	return nil
}

func (m *mockLogHandler) WithAttrs(as []slog.Attr) slog.Handler {
	if m.attr == nil {
		m.attr = make([]slog.Attr, 0)
	}
	m.attr = append(m.attr, as...)
	return m
}

func (m *mockLogHandler) WithGroup(name string) slog.Handler {
	m.group = name
	return m
}

func (m *mockLogHandler) GetAttr(key string) string {
	if m.attr == nil {
		return ""
	}
	for _, attr := range m.attr {
		if attr.Key == key {
			return attr.Value.String()
		}
	}

	return ""
}

func TestLogRequestID(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")

	h := new(mockLogHandler)
	logger := slog.New(h)
	log.Set(ctx, logger)

	require.NoError(t, tests.ExecuteMiddleware(ctx, echomw.RequestID()))
	require.NoError(t, tests.ExecuteMiddleware(ctx, SetLogger()))

	log.Ctx(ctx).Info("test")
	rID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	assert.Equal(t, rID, h.GetAttr("request_id"))
}

func TestLogRequest(t *testing.T) {
	statusCode := 200
	h := new(mockLogHandler)

	exec := func() {
		ctx, _ := tests.NewContext(c.Web, "http://test.localhost/abc")
		logger := slog.New(h).With("previous", "param")
		log.Set(ctx, logger)
		ctx.Request().Header.Set("Referer", "ref.com")
		ctx.Request().Header.Set(echo.HeaderXRealIP, "21.12.12.21")

		require.NoError(t, tests.ExecuteHandler(ctx, func(ctx echo.Context) error {
			return ctx.String(statusCode, "hello")
		},
			SetLogger(),
			LogRequest(),
		))
	}

	exec()
	assert.Equal(t, "param", h.GetAttr("previous"))
	assert.Equal(t, "21.12.12.21", h.GetAttr("ip"))
	assert.Equal(t, "test.localhost", h.GetAttr("host"))
	assert.Equal(t, "ref.com", h.GetAttr("referer"))
	assert.Equal(t, "200", h.GetAttr("status"))
	assert.Equal(t, "0", h.GetAttr("bytes_in"))
	assert.Equal(t, "5", h.GetAttr("bytes_out"))
	assert.NotEmpty(t, h.GetAttr("latency"))
	assert.Equal(t, "INFO", h.level)
	assert.Equal(t, "GET /abc", h.msg)

	statusCode = 500
	exec()
	assert.Equal(t, "ERROR", h.level)
}
