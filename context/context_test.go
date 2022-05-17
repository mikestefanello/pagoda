package context

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	assert.False(t, IsCanceledError(ctx.Err()))
	cancel()
	assert.True(t, IsCanceledError(ctx.Err()))

	ctx, cancel = context.WithTimeout(context.Background(), time.Microsecond)
	time.Sleep(time.Microsecond * 5)
	cancel()
	assert.False(t, IsCanceledError(ctx.Err()))

	assert.False(t, IsCanceledError(errors.New("test error")))
}
