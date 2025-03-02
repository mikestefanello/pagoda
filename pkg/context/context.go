package context

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
)

const (
	// AuthenticatedUserKey is the key used to store the authenticated user in context.
	AuthenticatedUserKey = "auth_user"

	// UserKey is the key used to store a user in context.
	UserKey = "user"

	// FormKey is the key used to store a form in context.
	FormKey = "form"

	// PasswordTokenKey is the key used to store a password token in context.
	PasswordTokenKey = "password_token"

	// LoggerKey is the key used to store a structured logger in context.
	LoggerKey = "logger"

	// SessionKey is the key used to store the session data in context.
	SessionKey = "session"

	// HTMXRequestKey is the key used to store the HTMX request data in context.
	HTMXRequestKey = "htmx"

	// CSRFKey is the key used to store the CSRF token in context.
	CSRFKey = "csrf"

	// ConfigKey is the key used to store the configuration in context.
	ConfigKey = "config"
)

// IsCanceledError determines if an error is due to a context cancellation.
func IsCanceledError(err error) bool {
	return errors.Is(err, context.Canceled)
}

// Cache checks if a value of a given type exists in the Echo context for a given key and returns that, otherwise
// it will use a callback to generate a value, which is stored in the context then returned. This allows you to
// only generate items only once for a given request.
func Cache[T any](ctx echo.Context, key string, gen func(echo.Context) T) T {
	if val := ctx.Get(key); val != nil {
		if v, ok := val.(T); ok {
			return v
		}
	}
	val := gen(ctx)
	ctx.Set(key, val)
	return val
}
