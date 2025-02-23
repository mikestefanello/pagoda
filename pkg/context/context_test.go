package context

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	assert.False(t, IsCanceledError(ctx.Err()))
	cancel()
	assert.True(t, IsCanceledError(ctx.Err()))

	ctx, cancel = context.WithTimeout(context.Background(), time.Microsecond*5)
	<-ctx.Done()
	cancel()
	assert.False(t, IsCanceledError(ctx.Err()))

	assert.False(t, IsCanceledError(errors.New("test error")))
}

func TestCache(t *testing.T) {
	ctx := echo.New().NewContext(nil, nil)

	key := "testing"
	value := "hello"
	called := 0
	callback := func(ctx echo.Context) string {
		called++
		return value
	}

	assert.Nil(t, ctx.Get(key))

	got := Cache(ctx, key, callback)
	assert.Equal(t, value, got)
	assert.Equal(t, 1, called)

	got = Cache(ctx, key, callback)
	assert.Equal(t, value, got)
	assert.Equal(t, 1, called)
}
