package log

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/tests"
	"github.com/stretchr/testify/assert"
)

func TestCtxSet(t *testing.T) {
	ctx, _ := tests.NewContext(echo.New(), "/")
	logger := Ctx(ctx)
	assert.NotNil(t, logger)

	logger = logger.With("a", "b")
	Set(ctx, logger)

	got := Ctx(ctx)
	assert.Equal(t, got, logger)
}
