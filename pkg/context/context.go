package context

import (
	"context"
	"errors"
)

const (
	// AuthenticatedUserKey is the key value used to store the authenticated user in context
	AuthenticatedUserKey = "auth_user"

	// UserKey is the key value used to store a user in context
	UserKey = "user"

	// FormKey is the key value used to store a form in context
	FormKey = "form"

	// PasswordTokenKey is the key value used to store a password token in context
	PasswordTokenKey = "password_token"
)

// IsCanceledError determines if an error is due to a context cancelation
func IsCanceledError(err error) bool {
	return errors.Is(err, context.Canceled)
}
