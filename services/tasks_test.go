package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskClient_New(t *testing.T) {
	now := time.Now()
	tk := c.Tasks.
		New("task1").
		Payload("payload").
		Queue("queue").
		Periodic("@every 5s").
		MaxRetries(5).
		Timeout(5 * time.Second).
		Deadline(now).
		At(now).
		Wait(6 * time.Second).
		Retain(7 * time.Second)

	assert.Equal(t, "task1", tk.typ)
	assert.Equal(t, "payload", tk.payload.(string))
	assert.Equal(t, "queue", *tk.queue)
	assert.Equal(t, "@every 5s", *tk.periodic)
	assert.Equal(t, 5, *tk.maxRetries)
	assert.Equal(t, 5*time.Second, *tk.timeout)
	assert.Equal(t, now, *tk.deadline)
	assert.Equal(t, now, *tk.at)
	assert.Equal(t, 6*time.Second, *tk.wait)
	assert.Equal(t, 7*time.Second, *tk.retain)
	assert.NoError(t, tk.Save())
}
