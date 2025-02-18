package daily_questions

import (
	"github.com/hibiken/asynq"
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/task_queue"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type StoreQuestionSolutionFileHandler struct {
	logger *utils.ApplicationLogger
}

func (handler StoreQuestionSolutionFileHandler) Handle(evt event_store.PublishableEvent) {
	solutionId := evt.ObjectId()
	scheduler := task_queue.GetTasksScheduler()
	task, _ := NewStoreQuestionSolutionFileTask(solutionId)
	scheduler.Enqueue(task, asynq.Queue(task_queue.QUEUE_QUESTIONS_STORAGE))
}

func NewStoreQuestionSolutionFileHandler() StoreQuestionSolutionFileHandler {
	return StoreQuestionSolutionFileHandler{
		utils.GetApplicationLogger(),
	}
}
