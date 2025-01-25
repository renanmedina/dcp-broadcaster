package daily_questions

import (
	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions/questions_solver"
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/task_queue"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type SolveQuestionEventHandler struct {
	logger *utils.ApplicationLogger
}

func (handler SolveQuestionEventHandler) Handle(evt event_store.PublishableEvent) {
	questionId := evt.ObjectId()
	scheduler := task_queue.GetTasksScheduler()
	for _, language := range questions_solver.SolvingLanguages {
		task, _ := NewSolveQuestionTask(questionId, language.LanguageName)
		scheduler.Enqueue(task)
	}
}

func NewSolveQuestionEventHandler() SolveQuestionEventHandler {
	return SolveQuestionEventHandler{
		utils.GetApplicationLogger(),
	}
}
