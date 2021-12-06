package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func CacheControl(maxAge time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			v := "no-cache, no-store"
			if maxAge > 0 {
				v = fmt.Sprintf("public, max-age=%.0f", maxAge.Seconds())
			}
			c.Response().Header().Set("Cache-Control", v)
			return next(c)
		}
	}
}
