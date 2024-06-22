package services

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type testTask struct {
	Val int
}

func (t testTask) Name() string {
	return "test_task"
}

func TestTaskClient_New(t *testing.T) {
	var subCalled bool

	queue := NewQueue[testTask](func(ctx context.Context, task testTask) error {
		subCalled = true
		assert.Equal(t, 123, task.Val)
		return nil
	})
	c.Tasks.Register(queue)

	task := testTask{Val: 123}

	tx := &sql.Tx{}

	op := c.Tasks.
		New(task).
		Wait(5 * time.Second).
		Tx(tx)

	// Check that the task op was built correctly
	assert.Equal(t, task, op.task)
	assert.Equal(t, tx, op.tx)
	assert.Equal(t, 5*time.Second, *op.wait)

	// Remove the transaction and delay so we can process the task immediately
	op.tx, op.wait = nil, nil
	err := op.Save()
	require.NoError(t, err)

	// Start the runner
	ctx, cancel := context.WithCancel(context.Background())
	go c.Tasks.StartRunner(ctx)
	defer cancel()

	// Check for up to 5 seconds if the task executed
	start := time.Now()
waitLoop:
	for {
		switch {
		case subCalled:
			break waitLoop
		case time.Since(start) > (5 * time.Second):
			break waitLoop
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}

	assert.True(t, subCalled)
}
