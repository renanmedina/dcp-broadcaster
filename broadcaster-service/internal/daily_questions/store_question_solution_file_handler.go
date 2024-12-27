package daily_questions

import (
	"errors"
	"fmt"
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
	err := use_case.Execute(solutionId)

	if errors.As(err, &QuestionSolutionNotFound{}) {
		handler.logger.Info(fmt.Sprintf("Handler failed to execute, waiting %s seconds to retry", STORE_QUESTION_RETRY_TIME))
		time.Sleep(STORE_QUESTION_RETRY_TIME)
		handler.Handle(evt)
	}
}

func NewStoreQuestionSolutionFileHandler() StoreQuestionSolutionFileHandler {
	return StoreQuestionSolutionFileHandler{
		utils.GetApplicationLogger(),
	}
}
