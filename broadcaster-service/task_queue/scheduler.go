package task_queue

import (
	"fmt"

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
	utils.GetApplicationLogger().Info(fmt.Sprintf("Connecting to redis at %s for task queue processing", addr))
	return asynq.RedisClientOpt{Addr: addr}
}
