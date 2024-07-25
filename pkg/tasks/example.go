package tasks

import (
	"context"
	"github.com/mikestefanello/backlite"
	"time"

	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/services"
)

// ExampleTask is an example implementation of backlite.Task
// This represents the task that can be queued for execution via the task client and should contain everything
// that your queue processor needs to process the task.
type ExampleTask struct {
	Message string
}

// Config satisfies the backlite.Task interface by providing configuration for the queue that these items will be
// placed into for execution.
func (t ExampleTask) Config() backlite.QueueConfig {
	return backlite.QueueConfig{
		Name:        "ExampleTask",
		MaxAttempts: 3,
		Timeout:     5 * time.Second,
		Backoff:     10 * time.Second,
		Retention: &backlite.Retention{
			Duration:   24 * time.Hour,
			OnlyFailed: false,
			Data: &backlite.RetainData{
				OnlyFailed: false,
			},
		},
	}
}

// NewExampleTaskQueue provides a Queue that can process ExampleTask tasks
// The service container is provided so the subscriber can have access to the app dependencies.
// All queues must be registered in the Register() function.
// Whenever an ExampleTask is added to the task client, it will be queued and eventually sent here for execution.
func NewExampleTaskQueue(c *services.Container) backlite.Queue {
	return backlite.NewQueue[ExampleTask](func(ctx context.Context, task ExampleTask) error {
		log.Default().Info("Example task received",
			"message", task.Message,
		)
		log.Default().Info("This can access the container for dependencies",
			"echo", c.Web.Reverse("home"),
		)
		return nil
	})
}
