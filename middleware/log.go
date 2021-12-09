package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// LogRequestID includes the request ID in all logs for the given request
func LogRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rid := c.Response().Header().Get(echo.HeaderXRequestID)
			format := fmt.Sprintf(`{"time":"${time_rfc3339_nano}","id":"%s","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}`, rid)
			c.Logger().SetHeader(format)
			return next(c)
		}
	}
}
