package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func CacheControl(maxAge int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			v := "no-cache, no-store"
			if maxAge > 0 {
				v = fmt.Sprintf("public, max-age=%d", maxAge)
			}
			c.Response().Header().Set("Cache-Control", v)
			return next(c)
		}
	}
}
