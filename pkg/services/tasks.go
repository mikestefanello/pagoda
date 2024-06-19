package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"github.com/maragudk/goqite"
	"github.com/maragudk/goqite/jobs"
	"github.com/mikestefanello/pagoda/config"
	"github.com/mikestefanello/pagoda/pkg/log"
)

type (
	// TaskClient is that client that allows you to queue or schedule task execution
	TaskClient struct {
		queue  *goqite.Queue
		runner *jobs.Runner
		db     *sql.DB
	}

	Task interface {
		Name() string
	}

	// TaskOp handles task creation operations
	TaskOp struct {
		client *TaskClient
		task   Task
		//payload    any
		//periodic   *string
		//queue      *string
		//maxRetries *int
		//timeout    *time.Duration
		//deadline   *time.Time
		at   *time.Time
		wait *time.Duration
		//retain     *time.Duration
	}

	Queuable interface {
		Name() string
		Receive(ctx context.Context, payload []byte) error
	}

	Queue[T any] struct {
		name       string
		subscriber QueueSubscriber[T]
	}

	QueueSubscriber[T any] func(context.Context, T) error
)

func NewQueue[T Task](subscriber QueueSubscriber[T]) *Queue[T] {
	var task T

	q := &Queue[T]{
		name:       task.Name(),
		subscriber: subscriber,
	}

	return q
}

func (q *Queue[T]) Name() string {
	return q.name
}

func (q *Queue[T]) Receive(ctx context.Context, payload []byte) error {
	var obj T
	err := json.Unmarshal(payload, &obj)
	if err != nil {
		return err
	}

	return q.subscriber(ctx, obj)
}

// NewTaskClient creates a new task client
func NewTaskClient(cfg *config.Config) (*TaskClient, error) {
	var connection string

	switch cfg.App.Environment {
	case config.EnvTest:
		connection = cfg.Tasks.TestConnection
	default:
		connection = cfg.Tasks.Connection
	}

	db, err := openDB(cfg.Tasks.Driver, connection)
	if err != nil {
		return nil, err
	}

	//db.SetMaxOpenConns(1)
	//db.SetMaxIdleConns(1)

	// Install the schema
	if err := goqite.Setup(context.Background(), db); err != nil {
		// An error is returned if we already ran this and there's no better way to check.
		// You can and probably should handle this via migrations
		if !strings.Contains(err.Error(), "already exists") {
			return nil, err
		}
	}

	t := &TaskClient{
		queue: goqite.New(goqite.NewOpts{
			DB:         db,
			Name:       "jobs",
			MaxReceive: cfg.Tasks.MaxRetries,
		}),
		db: db,
	}

	t.runner = jobs.NewRunner(jobs.NewRunnerOpts{
		Limit:        cfg.Tasks.Goroutines,
		Log:          log.Default(),
		PollInterval: cfg.Tasks.PollInterval,
		Queue:        t.queue,
	})

	return t, nil
}

// Close closes the connection to the task service
func (t *TaskClient) Close() error {
	// TODO close the runner
	return t.db.Close()
}

// StartRunner starts the scheduler service which adds scheduled tasks to the queue
// This must be running in order to queue tasks set for periodic execution
func (t *TaskClient) StartRunner(ctx context.Context) {
	t.runner.Start(ctx)
}

//func (t *TaskClient) Register(name string, processor jobs.Func) {
//	t.runner.Register(name, processor)
//}

func (t *TaskClient) Register(queue Queuable) {
	t.runner.Register(queue.Name(), queue.Receive)
}

// New starts a task creation operation
func (t *TaskClient) New(task Task) *TaskOp {
	return &TaskOp{
		client: t,
		task:   task,
	}
}

//// Payload sets the task payload data which will be sent to the task handler
//func (t *TaskOp) Payload(payload Task) *TaskOp {
//	t.payload = payload
//	return t
//}

// // Periodic sets the task to execute periodically according to a given interval
// // The interval can be either in cron form ("*/5 * * * *") or "@every 30s"
//
//	func (t *TaskOp) Periodic(interval string) *TaskOp {
//		t.periodic = &interval
//		return t
//	}
//
// // Queue specifies the name of the queue to add the task to
// // The default queue will be used if this is not set
//
//	func (t *TaskOp) Queue(queue string) *TaskOp {
//		t.queue = &queue
//		return t
//	}
//
// // Timeout sets the task timeout, meaning the task must execute within a given duration
//
//	func (t *TaskOp) Timeout(timeout time.Duration) *TaskOp {
//		t.timeout = &timeout
//		return t
//	}
//
// // Deadline sets the task execution deadline to a specific date and time
//
//	func (t *TaskOp) Deadline(deadline time.Time) *TaskOp {
//		t.deadline = &deadline
//		return t
//	}
//

// At sets the exact date and time the task should be executed
func (t *TaskOp) At(processAt time.Time) *TaskOp {
	until := time.Until(processAt)
	t.wait = &until
	return t
}

// Wait instructs the task to wait a given duration before it is executed
func (t *TaskOp) Wait(duration time.Duration) *TaskOp {
	t.wait = &duration
	return t
}

//
//// Retain instructs the task service to retain the task data for a given duration after execution is complete
//func (t *TaskOp) Retain(duration time.Duration) *TaskOp {
//	t.retain = &duration
//	return t
//}
//
//// MaxRetries sets the maximum amount of times to retry executing the task in the event of a failure
//func (t *TaskOp) MaxRetries(retries int) *TaskOp {
//	t.maxRetries = &retries
//	return t
//}

// Save saves the task so it can be executed
func (t *TaskOp) Save() error {
	var err error

	// Build the payload
	payload, err := json.Marshal(t.task)
	if err != nil {
		return err
	}

	// Build the task options
	//opts := make([]asynq.Option, 0)
	//if t.queue != nil {
	//	opts = append(opts, asynq.Queue(*t.queue))
	//}
	//if t.maxRetries != nil {
	//	opts = append(opts, asynq.MaxRetry(*t.maxRetries))
	//}
	//if t.timeout != nil {
	//	opts = append(opts, asynq.Timeout(*t.timeout))
	//}
	//if t.deadline != nil {
	//	opts = append(opts, asynq.Deadline(*t.deadline))
	//}
	//if t.wait != nil {
	//	opts = append(opts, asynq.ProcessIn(*t.wait))
	//}
	//if t.retain != nil {
	//	opts = append(opts, asynq.Retention(*t.retain))
	//}
	//if t.at != nil {
	//	opts = append(opts, asynq.ProcessAt(*t.at))
	//}

	msg := goqite.Message{
		Body: payload,
	}

	if t.wait != nil {
		msg.Delay = *t.wait
	}
	//return t.client.queue.Send(context.Background(), msg)
	return jobs.Create(context.Background(), t.client.queue, t.task.Name(), payload)
}
