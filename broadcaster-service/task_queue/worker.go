package task_queue

import (
	"github.com/hibiken/asynq"
)

type QueueWorker struct {
	server *asynq.Server
	muxes  *asynq.ServeMux
}

func (qw *QueueWorker) Run() error {
	return qw.server.Run(qw.muxes)
}

func (qw *QueueWorker) RegisterTaskProcessor(taskType string, handler asynq.HandlerFunc) {
	qw.muxes.HandleFunc(taskType, handler)
}

func (qw *QueueWorker) RegisterTask(taskType string, handler asynq.HandlerFunc) {
	qw.muxes.HandleFunc(taskType, handler)
}

var queueWorker *QueueWorker

func init() {
	if queueWorker == nil {
		queueWorker = InitializeQueueServer()
	}
}

func InitializeQueueServer() *QueueWorker {
	srv := newQueueServer()
	mux := asynq.NewServeMux()
	return &QueueWorker{srv, mux}
}

func newQueueServer() *asynq.Server {
	return asynq.NewServer(
		GetQueueClientOptions(),
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				QUEUE_MESSAGES_DELIVERIES: 3,
				QUEUE_QUESTIONS_SOLUTIONS: 3,
				QUEUE_QUESTIONS_STORAGE:   3,
				QUEUE_DEFAULT:             1,
			},
		},
	)
}
