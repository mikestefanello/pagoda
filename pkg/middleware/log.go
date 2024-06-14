package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/log"
)

// SetLogger initializes a logger for the current request and stores it in the context.
// It's recommended to have this executed after Echo's RequestID() middleware because it will add
// the request ID to the logger so that all log messages produced from this request have the
// request ID in it. You can modify this code to include any other fields that you want to always
// appear.
func SetLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Include the request ID in the logger
			rID := ctx.Response().Header().Get(echo.HeaderXRequestID)
			logger := log.Ctx(ctx).With("request_id", rID)

			// TODO include other fields you may want in all logs for this request
			log.Set(ctx, logger)
			return next(ctx)
		}
	}
}

// LogRequest logs the current request
// Echo provides middleware similar to this, but we want to use our own logger
func LogRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			req := ctx.Request()
			res := ctx.Response()

			// Track how long the request takes to complete
			start := time.Now()
			if err = next(ctx); err != nil {
				ctx.Error(err)
			}
			stop := time.Now()

			sub := log.Ctx(ctx).With(
				"ip", ctx.RealIP(),
				"host", req.Host,
				"method", req.Method,
				"path", func() string {
					p := req.URL.Path
					if p == "" {
						p = "/"
					}
					return p
				}(),
				"referer", req.Referer(),
				"status", res.Status,
				"bytes_in", func() string {
					cl := req.Header.Get(echo.HeaderContentLength)
					if cl == "" {
						cl = "0"
					}
					return cl
				}(),
				"bytes_out", strconv.FormatInt(res.Size, 10),
				"latency", stop.Sub(start).String(),
			)

			// TODO is there a (better) way to log without a message?

			if res.Status >= 500 {
				sub.Error("")
			} else {
				sub.Info("")
			}

			return nil
		}
	}
}
