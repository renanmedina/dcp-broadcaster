package daily_questions

import (
	"log"
	"time"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

type QuestionsWorker struct {
	questionsService QuestionsService
	logger           *utils.ApplicationLogger
}

func (r *QuestionsWorker) Work(every time.Duration) {
	use_case, err := NewFetchNewQuestions()

	if err != nil {
		log.Fatal(err.Error())
	}

	ticker := time.NewTicker(every)

	for {
		select {
		case <-ticker.C:
			use_case.Execute()
		}
	}
}

func NewQuestionsReceiver() (QuestionsWorker, error) {
	service, err := NewQuestionsService()

	if err != nil {
		return QuestionsWorker{}, err
	}

	return QuestionsWorker{service, utils.GetApplicationLogger()}, nil
}
