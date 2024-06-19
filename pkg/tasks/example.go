package tasks

type ExampleTask struct {
	Message string
}

func (t ExampleTask) Name() string {
	return "example_task"
}
