package msg

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Type is a message type
type Type string

const (
	// TypeSuccess represents a success message type
	TypeSuccess Type = "success"

	// TypeInfo represents a info message type
	TypeInfo Type = "info"

	// TypeWarning represents a warning message type
	TypeWarning Type = "warning"

	// TypeDanger represents a danger message type
	TypeDanger Type = "danger"
)

const (
	// sessionName stores the name of the session which contains flash messages
	sessionName = "msg"
)

// Success sets a success flash message
func Success(ctx echo.Context, message string) {
	Set(ctx, TypeSuccess, message)
}

// Info sets an info flash message
func Info(ctx echo.Context, message string) {
	Set(ctx, TypeInfo, message)
}

// Warning sets a warning flash message
func Warning(ctx echo.Context, message string) {
	Set(ctx, TypeWarning, message)
}

// Danger sets a danger flash message
func Danger(ctx echo.Context, message string) {
	Set(ctx, TypeDanger, message)
}

// Set adds a new flash message of a given type into the session storage
// Errors will logged and not returned
func Set(ctx echo.Context, typ Type, message string) {
	if sess, err := getSession(ctx); err == nil {
		sess.AddFlash(message, string(typ))
		save(ctx, sess)
	}
}

// Get gets flash messages of a given type from the session storage
// Errors will logged and not returned
func Get(ctx echo.Context, typ Type) []string {
	var msgs []string

	if sess, err := getSession(ctx); err == nil {
		if flash := sess.Flashes(string(typ)); len(flash) > 0 {
			save(ctx, sess)

			for _, m := range flash {
				msgs = append(msgs, m.(string))
			}
		}
	}

	return msgs
}

// getSession gets the flash message session
func getSession(ctx echo.Context) (*sessions.Session, error) {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		ctx.Logger().Errorf("cannot load flash message session: %v", err)
	}
	return sess, err
}

// save saves the flash message session
func save(ctx echo.Context, sess *sessions.Session) {
	if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
		ctx.Logger().Errorf("failed to set flash message: %v", err)
	}
}
