package daily_questions

import (
	"time"

	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type StoreQuestionSolutionFileHandler struct {
	logger *utils.ApplicationLogger
}

const STORE_QUESTION_RETRY_TIME = time.Second * 5

func (handler StoreQuestionSolutionFileHandler) Handle(evt event_store.PublishableEvent) {
	solutionId := evt.ObjectId()
	use_case := NewStoreQuestionSolutionFile()
	use_case.Execute(solutionId)
}

func NewStoreQuestionSolutionFileHandler() StoreQuestionSolutionFileHandler {
	return StoreQuestionSolutionFileHandler{
		utils.GetApplicationLogger(),
	}
}
