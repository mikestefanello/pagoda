package tasks

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
)

// TypeExample is the type for the example task.
// This is what is passed in to TaskClient.New() when creating a new task
const TypeExample = "example_task"

// ExampleProcessor processes example tasks
type ExampleProcessor struct {
}

// ProcessTask handles the processing of the task
func (p *ExampleProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	log.Printf("executing task: %s", t.Type())
	return nil
}
