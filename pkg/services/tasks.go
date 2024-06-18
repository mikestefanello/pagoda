package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log/slog"
	"strings"
	"time"

	"github.com/maragudk/goqite"
	"github.com/maragudk/goqite/jobs"
	"github.com/mikestefanello/pagoda/config"
)

type (
	// TaskClient is that client that allows you to queue or schedule task execution
	TaskClient struct {
		queue  *goqite.Queue
		runner *jobs.Runner
		db     *sql.DB
	}

	// task handles task creation operations
	task struct {
		client     *TaskClient
		typ        string
		payload    any
		periodic   *string
		queue      *string
		maxRetries *int
		timeout    *time.Duration
		deadline   *time.Time
		at         *time.Time
		wait       *time.Duration
		retain     *time.Duration
	}

	Queue[T any] struct {
		name       string
		q          *goqite.Queue
		subscriber func(context.Context, T) error
	}
)

var queues = make(map[string]Queuable)

func NewQueue[T any](name string) *Queue[T] {
	q := &Queue[T]{name: name}
	queues[name] = q
	return q
}

func GetQueue[T any](name string) *Queue[T] {
	return queues[name].(*Queue[T])
}

type Queuable interface {
	Receive(ctx context.Context, payload []byte) error
}

func (q *Queue[T]) Add(item T) error {
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return jobs.Create(context.Background(), q.q, q.name, b)
}

func (q *Queue[T]) Receive(ctx context.Context, payload []byte) error {
	var obj T
	err := json.Unmarshal(payload, &obj)
	if err != nil {
		return err
	}

	return q.subscriber(ctx, obj)
}

func (q *Queue[T]) Register(r *jobs.Runner) {
	r.Register(q.name, q.Receive)
}

// NewTaskClient creates a new task client
func NewTaskClient(cfg *config.Config) (*TaskClient, error) {
	db, err := openDB("sqlite3", "dbs/tasks.db?_journal=WAL&_timeout=5000&_fk=true")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	// Install the schema
	if err := goqite.Setup(context.Background(), db); err != nil {
		// An error is returned if we already ran this
		if !strings.Contains(err.Error(), "already exists") {
			return nil, err
		}
	}
	// Determine the database based on the environment
	//db := cfg.Cache.Database
	//if cfg.App.Environment == config.EnvTest {
	//	db = cfg.Cache.TestDatabase
	//}
	// TODO test db

	t := &TaskClient{
		queue: goqite.New(goqite.NewOpts{
			DB:         db,
			Name:       "jobs",
			MaxReceive: 10,
		}),
		db: db,
	}

	t.runner = jobs.NewRunner(jobs.NewRunnerOpts{
		Limit:        1,
		Log:          slog.Default(),
		PollInterval: 10 * time.Millisecond,
		Queue:        t.queue,
	})

	return t, nil
}

// Close closes the connection to the task service
func (t *TaskClient) Close() error {
	return t.db.Close()
}

// StartRunner starts the scheduler service which adds scheduled tasks to the queue
// This must be running in order to queue tasks set for periodic execution
func (t *TaskClient) StartRunner(ctx context.Context) {
	t.runner.Start(ctx)
}

func (t *TaskClient) Register(name string, processor jobs.Func) {
	t.runner.Register(name, processor)
}

// New starts a task creation operation
func (t *TaskClient) New(typ string) *task {
	return &task{
		client: t,
		typ:    typ,
	}
}

// Payload sets the task payload data which will be sent to the task handler
func (t *task) Payload(payload any) *task {
	t.payload = payload
	return t
}

// // Periodic sets the task to execute periodically according to a given interval
// // The interval can be either in cron form ("*/5 * * * *") or "@every 30s"
//
//	func (t *task) Periodic(interval string) *task {
//		t.periodic = &interval
//		return t
//	}
//
// // Queue specifies the name of the queue to add the task to
// // The default queue will be used if this is not set
//
//	func (t *task) Queue(queue string) *task {
//		t.queue = &queue
//		return t
//	}
//
// // Timeout sets the task timeout, meaning the task must execute within a given duration
//
//	func (t *task) Timeout(timeout time.Duration) *task {
//		t.timeout = &timeout
//		return t
//	}
//
// // Deadline sets the task execution deadline to a specific date and time
//
//	func (t *task) Deadline(deadline time.Time) *task {
//		t.deadline = &deadline
//		return t
//	}
//
// // At sets the exact date and time the task should be executed
//
//	func (t *task) At(processAt time.Time) *task {
//		t.at = &processAt
//		return t
//	}
//
// Wait instructs the task to wait a given duration before it is executed
func (t *task) Wait(duration time.Duration) *task {
	t.wait = &duration
	return t
}

//
//// Retain instructs the task service to retain the task data for a given duration after execution is complete
//func (t *task) Retain(duration time.Duration) *task {
//	t.retain = &duration
//	return t
//}
//
//// MaxRetries sets the maximum amount of times to retry executing the task in the event of a failure
//func (t *task) MaxRetries(retries int) *task {
//	t.maxRetries = &retries
//	return t
//}

// Save saves the task so it can be executed
func (t *task) Save() error {
	var err error

	// Build the payload
	var payload []byte
	if t.payload != nil {
		if payload, err = json.Marshal(t.payload); err != nil {
			return err
		}
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
	return t.client.queue.Send(context.Background(), msg)
	//return jobs.Create(context.Background(), t.client.queue, t.typ, payload)
}
