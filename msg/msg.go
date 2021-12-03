package msg

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Type string

const (
	Success Type = "success"
	Info    Type = "info"
	Warning Type = "warning"
	Danger  Type = "danger"
)

// TODO: Error handling

func getSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("msg", c)
	return sess
}

// Set adds a new message into the cookie storage.
func Set(c echo.Context, typ Type, value string) {
	sess := getSession(c)
	sess.AddFlash(value, string(typ))
	_ = sess.Save(c.Request(), c.Response())
}

// Get gets flash messages from the cookie storage.
func Get(c echo.Context, typ Type) []string {
	sess := getSession(c)
	fm := sess.Flashes(string(typ))
	// If we have some messages.
	if len(fm) > 0 {
		_ = sess.Save(c.Request(), c.Response())
		// Initiate a strings slice to return messages.
		var flashes []string
		for _, fl := range fm {
			// Add message to the slice.
			flashes = append(flashes, fl.(string))
		}

		return flashes
	}
	return nil
}
