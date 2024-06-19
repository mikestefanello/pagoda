package tasks

import (
	"context"

	"github.com/mikestefanello/pagoda/pkg/log"
	"github.com/mikestefanello/pagoda/pkg/services"
)

// ExampleTask is an example implementation of services.Task
// This represents the task that can be queued for execution via the task client and should contain everything
// that your queue subscriber needs to process the task.
type ExampleTask struct {
	Message string
}

// Name satisfies the services.Task interface by proviing a unique name for this Task type
func (t ExampleTask) Name() string {
	return "example_task"
}

// NewExampleTaskQueue provides a Queue that can process ExampleTask tasks
// The service container is provided so the subscriber can have access to the app dependencies.
// All queues must be registered in the Register() function.
// Whenever an ExampleTask is added to the task client, it will be queued and eventually sent here for execution.
func NewExampleTaskQueue(c *services.Container) services.Queue {
	return services.NewQueue[ExampleTask](func(ctx context.Context, task ExampleTask) error {
		log.Default().Info("Example task received",
			"message", task.Message,
		)
		log.Default().Info("This can access the container for dependencies",
			"echo", c.Web.Reverse("home"),
		)
		return nil
	})
}
