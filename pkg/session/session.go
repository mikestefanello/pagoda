package session

import (
	"errors"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/context"
	"github.com/mikestefanello/pagoda/pkg/log"
)

// ErrStoreNotFound indicates that the session store was not present in the context
var ErrStoreNotFound = errors.New("session store not found")

// Get returns a session. If the incoming request has a session cookie that can't be decoded -
// eg. because it was encrypted with a since-rotated App.EncryptionKey, has expired, or was
// tampered with - this does not return an error. gorilla/sessions always hands back a valid,
// brand new session in that situation (see the sessions.Store.Get documentation), so it's both
// safe and expected to treat the caller as if they simply don't have a session yet, rather than
// surfacing what is normally a benign, one-time situation as a fatal error.
func Get(ctx echo.Context, name string) (*sessions.Session, error) {
	s := ctx.Get(context.SessionKey)
	if s == nil {
		return nil, ErrStoreNotFound
	}
	store := s.(sessions.Store)
	sess, err := store.Get(ctx.Request(), name)
	if err != nil {
		log.Ctx(ctx).Debug("discarding undecodable session cookie", "name", name, "error", err)
	}
	return sess, nil
}

// Store sets the session storage in the context
func Store(ctx echo.Context, store sessions.Store) {
	ctx.Set(context.SessionKey, store)
}
