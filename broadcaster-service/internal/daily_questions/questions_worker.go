package daily_questions

import (
	"time"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

type QuestionsWorker struct {
	questionsService QuestionsService
	logger           *utils.ApplicationLogger
}

func (worker *QuestionsWorker) Work(every time.Duration) {
	worker.logger.Info("Starting questions receiver worker")

	use_case, err := NewFetchNewQuestions()

	if err != nil {
		worker.logger.Fatal(err.Error())
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
