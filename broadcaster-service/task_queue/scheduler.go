package task_queue

import (
	"github.com/hibiken/asynq"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

var queueClient *asynq.Client

func init() {
	if queueClient == nil {
		queueClient = NewQueueClient()
	}
}

func NewQueueClient() *asynq.Client {
	return asynq.NewClient(GetQueueClientOptions())
}

func GetTasksScheduler() *asynq.Client {
	return queueClient
}

func GetQueueClientOptions() asynq.RedisClientOpt {
	addr := utils.GetConfigs().TASKS_QUEUE_DB_URL
	return asynq.RedisClientOpt{Addr: addr}
}
