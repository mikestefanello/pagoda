package auth

// NotAuthenticatedError is an error returned when a user is not authenticated
type NotAuthenticatedError struct{}

// Error implements the error interface.
func (e NotAuthenticatedError) Error() string {
	return "user not authenticated"
}

// InvalidTokenError is an error returned when an invalid token is provided
type InvalidTokenError struct{}

// Error implements the error interface.
func (e InvalidTokenError) Error() string {
	return "invalid token"
}
