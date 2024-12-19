package daily_questions

import (
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
)

type StoreQuestionSolutionFileHandler struct{}

func (handler StoreQuestionSolutionFileHandler) Handle(evt event_store.PublishableEvent) {
	solutionId := evt.ObjectId()
	use_case := NewStoreQuestionSolutionFile()
	use_case.Execute(solutionId)
}

func NewStoreQuestionSolutionFileHandler() StoreQuestionSolutionFileHandler {
	return StoreQuestionSolutionFileHandler{}
}
