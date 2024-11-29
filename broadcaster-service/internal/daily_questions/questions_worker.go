package daily_questions

import (
	"time"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

type QuestionsWorker struct {
	logger *utils.ApplicationLogger
}

func (worker *QuestionsWorker) Work(every time.Duration, runImmediately bool) {
	worker.logger.Info("Starting questions receiver worker")

	use_case, err := NewFetchNewQuestions()

	if err != nil {
		worker.logger.Fatal(err.Error())
	}

	if runImmediately {
		use_case.Execute() // imediately calls first run
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
	return QuestionsWorker{utils.GetApplicationLogger()}, nil
}

func StartWorker(every time.Duration) {
	receiver, err := NewQuestionsReceiver()

	if err != nil {
		receiver.logger.Fatal(err.Error())
	}

	receiver.Work(every, true)
}
