package tasks

// TypeExample is the type for the example task.
// This is what is passed in to TaskClient.New() when creating a new task
const TypeExample = "example_task"

type ExampleTask struct {
	Message string
}
