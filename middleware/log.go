package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// LogRequestID includes the request ID in all logs for the given request
// This requires that middleware that includes the request ID first execute
func LogRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rID := c.Response().Header().Get(echo.HeaderXRequestID)
			format := `{"time":"${time_rfc3339_nano}","id":"%s","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}`
			c.Logger().SetHeader(fmt.Sprintf(format, rID))
			return next(c)
		}
	}
}
