package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/gob"
	"strings"
	"sync"
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
		queue   *goqite.Queue
		runner  *jobs.Runner
		buffers sync.Pool
	}

	// Task is a job that can be added to a queue and later passed to and executed by a QueueSubscriber.
	// See pkg/tasks for an example of how this can be used with a queue.
	Task interface {
		Name() string
	}

	// TaskSaveOp handles task save operations
	TaskSaveOp struct {
		client *TaskClient
		task   Task
		tx     *sql.Tx
		at     *time.Time
		wait   *time.Duration
	}

	// Queue is a queue that a Task can be pushed to for execution.
	// While this can be implemented directly, it's recommended to use NewQueue() which uses generics in
	// order to provide type-safe queues and queue subscriber callbacks for task execution.
	Queue interface {
		// Name returns the name of the task this queue processes
		Name() string

		// Receive receives the Task payload to be processed
		Receive(ctx context.Context, payload []byte) error
	}

	// queue provides a type-safe implementation of Queue
	queue[T Task] struct {
		name       string
		subscriber QueueSubscriber[T]
	}

	// QueueSubscriber is a generic subscriber callback for a given queue to process Tasks
	QueueSubscriber[T Task] func(context.Context, T) error
)

// NewTaskClient creates a new task client
func NewTaskClient(cfg config.TasksConfig, db *sql.DB) (*TaskClient, error) {
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
		buffers: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(nil)
			},
		},
	}

	t.runner = jobs.NewRunner(jobs.NewRunnerOpts{
		Limit:        cfg.Goroutines,
		Log:          log.Default(),
		PollInterval: cfg.PollInterval,
		Queue:        t.queue,
	})

	return t, nil
}

// StartRunner starts the scheduler service which adds scheduled tasks to the queue.
// This must be running in order to execute queued tasked.
// To stop the runner, cancel the context.
// This is a blocking call.
func (t *TaskClient) StartRunner(ctx context.Context) {
	t.runner.Start(ctx)
}

// Register registers a queue so tasks can be added to it and processed
func (t *TaskClient) Register(queue Queue) {
	t.runner.Register(queue.Name(), queue.Receive)
}

// New starts a task save operation
func (t *TaskClient) New(task Task) *TaskSaveOp {
	return &TaskSaveOp{
		client: t,
		task:   task,
	}
}

// At sets the exact date and time the task should be executed
func (t *TaskSaveOp) At(processAt time.Time) *TaskSaveOp {
	t.Wait(time.Until(processAt))
	return t
}

// Wait instructs the task to wait a given duration before it is executed
func (t *TaskSaveOp) Wait(duration time.Duration) *TaskSaveOp {
	t.wait = &duration
	return t
}

// Tx will include the task as part of a given database transaction
func (t *TaskSaveOp) Tx(tx *sql.Tx) *TaskSaveOp {
	t.tx = tx
	return t
}

// Save saves the task, so it can be queued for execution
func (t *TaskSaveOp) Save() error {
	type message struct {
		Name    string
		Message []byte
	}

	// Encode the task
	taskBuf := t.client.buffers.Get().(*bytes.Buffer)
	if err := gob.NewEncoder(taskBuf).Encode(t.task); err != nil {
		return err
	}

	// Wrap and encode the message
	// This is needed as a workaround because goqite doesn't support delays using the jobs package,
	// so we format the message the way it expects but use the queue to supply the delay
	msgBuf := t.client.buffers.Get().(*bytes.Buffer)
	wrapper := message{Name: t.task.Name(), Message: taskBuf.Bytes()}
	if err := gob.NewEncoder(msgBuf).Encode(wrapper); err != nil {
		return err
	}

	msg := goqite.Message{
		Body: msgBuf.Bytes(),
	}

	if t.wait != nil {
		msg.Delay = *t.wait
	}

	// Put the buffers back in the pool for re-use
	taskBuf.Reset()
	msgBuf.Reset()
	t.client.buffers.Put(taskBuf)
	t.client.buffers.Put(msgBuf)

	if t.tx == nil {
		return t.client.queue.Send(context.Background(), msg)
	} else {
		return t.client.queue.SendTx(context.Background(), t.tx, msg)
	}
}

// NewQueue queues a new type-safe Queue of a given Task type
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
	err := gob.NewDecoder(bytes.NewReader(payload)).Decode(&obj)
	if err != nil {
		return err
	}

	return q.subscriber(ctx, obj)
}
