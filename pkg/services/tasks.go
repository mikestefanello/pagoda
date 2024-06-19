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
	// TaskClient is that client that allows you to queue or schedule task execution.
	// Under the hood we create only a single queue using goqite for all tasks because we do not want more than one
	// runner to process the tasks. The TaskClient wrapper provides abstractions for separate, type-safe queues.
	TaskClient struct {
		queue  *goqite.Queue
		runner *jobs.Runner
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
		tx   *sql.Tx
		at   *time.Time
		wait *time.Duration
		//retain     *time.Duration
	}

	Queue interface {
		Name() string
		Receive(ctx context.Context, payload []byte) error
	}

	queue[T any] struct {
		name       string
		subscriber QueueSubscriber[T]
	}

	QueueSubscriber[T any] func(context.Context, T) error
)

func NewQueue[T Task](subscriber QueueSubscriber[T]) Queue {
	var task T

	q := &queue[T]{
		name:       task.Name(),
		subscriber: subscriber,
	}

	return q
}

func (q *queue[T]) Name() string {
	return q.name
}

func (q *queue[T]) Receive(ctx context.Context, payload []byte) error {
	var obj T
	err := json.Unmarshal(payload, &obj)
	if err != nil {
		return err
	}

	return q.subscriber(ctx, obj)
}

// NewTaskClient creates a new task client
func NewTaskClient(cfg config.TasksConfig, db *sql.DB) (*TaskClient, error) {
	//var connection string
	//
	//switch cfg.App.Environment {
	//case config.EnvTest:
	//	connection = cfg.Tasks.TestConnection
	//default:
	//	connection = cfg.Tasks.Connection
	//}
	//
	//db, err := openDB(cfg.Tasks.Driver, connection)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// TODO is this correct?
	//db.SetMaxOpenConns(cfg.Tasks.Goroutines)
	//db.SetMaxIdleConns(cfg.Tasks.Goroutines)

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
			Name:       "tasks",
			MaxReceive: cfg.MaxRetries,
		}),
	}

	t.runner = jobs.NewRunner(jobs.NewRunnerOpts{
		Limit:        cfg.Goroutines,
		Log:          log.Default(),
		PollInterval: cfg.PollInterval,
		Queue:        t.queue,
	})

	return t, nil
}

//// Close closes the connection to the task service
//func (t *TaskClient) Close() error {
//	// TODO close the runner
//	return t.db.Close()
//}

// StartRunner starts the scheduler service which adds scheduled tasks to the queue.
// This must be running in order to execute queued tasked.
// To stop the runner, cancel the context.
func (t *TaskClient) StartRunner(ctx context.Context) {
	t.runner.Start(ctx)
}

func (t *TaskClient) Register(queue Queue) {
	t.runner.Register(queue.Name(), queue.Receive)
}

// New starts a task creation operation
func (t *TaskClient) New(task Task) *TaskOp {
	return &TaskOp{
		client: t,
		task:   task,
	}
}

// At sets the exact date and time the task should be executed
func (t *TaskOp) At(processAt time.Time) *TaskOp {
	t.Wait(time.Until(processAt))
	return t
}

// Wait instructs the task to wait a given duration before it is executed
func (t *TaskOp) Wait(duration time.Duration) *TaskOp {
	t.wait = &duration
	return t
}

// Tx will insert the task as part of a given database transaction
func (t *TaskOp) Tx(tx *sql.Tx) *TaskOp {
	t.tx = tx
	return t
}

// Save saves the task so it can be executed
func (t *TaskOp) Save() error {
	// Build the payload
	// TODO use gob?
	payload, err := json.Marshal(t.task)
	if err != nil {
		return err
	}

	//msg := goqite.Message{
	//	Body: payload,
	//}
	//
	//if t.wait != nil {
	//	msg.Delay = *t.wait
	//}
	// TODO support delay
	//return t.client.queue.Send(context.Background(), msg)
	if t.tx == nil {
		return jobs.Create(context.Background(), t.client.queue, t.task.Name(), payload)
	} else {
		return jobs.CreateTx(context.Background(), t.tx, t.client.queue, t.task.Name(), payload)
	}
}
