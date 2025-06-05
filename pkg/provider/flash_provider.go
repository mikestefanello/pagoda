package provider

import (
	"context"
	"errors"

	"github.com/labstack/echo-contrib/session"
	inertia "github.com/romsar/gonertia/v2"
)

const (
	FlashErrorsKey = "_flash_errors"
)

type SessionFlashProvider struct{}

func NewSessionFlashProvider() *SessionFlashProvider {
	return &SessionFlashProvider{}
}

func (p *SessionFlashProvider) FlashErrors(ctx context.Context, validationErrors inertia.ValidationErrors) error {
	eCtx, ok := FromEchoContext(ctx)
	if !ok {
		return errors.New("echo.Context not found in context")
	}
	sess, err := session.Get(FlashErrorsKey, eCtx)
	if err != nil {
		return err
	}
	sess.AddFlash(validationErrors)
	return sess.Save(eCtx.Request(), eCtx.Response())
}

func (p *SessionFlashProvider) GetErrors(ctx context.Context) (inertia.ValidationErrors, error) {
	eCtx, ok := FromEchoContext(ctx)
	if !ok {
		return nil, errors.New("echo.Context not found in context")
	}
	sess, err := session.Get(FlashErrorsKey, eCtx)
	if err != nil {
		return nil, err
	}
	if flashes := sess.Flashes(); len(flashes) > 0 {
		err := sess.Save(eCtx.Request(), eCtx.Response())
		if err != nil {
			return nil, err
		}
		return flashes[0].(inertia.ValidationErrors), nil
	}
	return nil, nil
}
