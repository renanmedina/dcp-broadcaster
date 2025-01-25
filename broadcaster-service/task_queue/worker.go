package task_queue

import (
	"github.com/hibiken/asynq"
	"github.com/renanmedina/dcp-broadcaster/utils"
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
	srv := newQueueServer(utils.GetConfigs().TASKS_QUEUE_DB_URL)
	mux := asynq.NewServeMux()
	return &QueueWorker{srv, mux}
}

func newQueueServer(redisAddr string) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"questions-message-deliveries": 3,
				"questions-solutions":          4,
				"questions-storage":            3,
			},
		},
	)
}
