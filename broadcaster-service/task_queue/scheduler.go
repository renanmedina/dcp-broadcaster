package task_queue

import (
	"github.com/hibiken/asynq"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

var queueClient *asynq.Client

func init() {
	if queueClient == nil {
		queueClient = NewQueueClient(utils.GetConfigs().TASKS_QUEUE_DB_URL)
	}
}

func NewQueueClient(dbUrl string) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: dbUrl})
}

func GetTasksScheduler() *asynq.Client {
	return queueClient
}
