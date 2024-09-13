package daily_questions

import (
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type FetchNewQuestions struct {
	service             QuestionsService
	questionsRepository QuestionsRepository
	logger              *utils.ApplicationLogger
	publisher           *event_store.EventPublisher
}

func (uc *FetchNewQuestions) Execute() {
	questions, err := uc.service.GetNewQuestions(1)

	if err != nil {
		uc.logger.Error(err.Error())
	}

	uc.processQuestions(questions)
}

func (uc *FetchNewQuestions) processQuestions(questions []Question) {
	for _, question := range questions {
		uc.logger.Info("Processing message received from questions service", "question", question.ToLogMap())
		_, err := uc.questionsRepository.Save(question)
		if err != nil {
			errMsg := fmt.Sprintf("Failed processing message received from questions service: %s", err.Error())
			uc.logger.Error(errMsg, "question", question.ToLogMap())
			continue
		}

		uc.logger.Info("Processed message received from questions service", "question", question.ToLogMap())
		uc.publisher.Publish(newQuestionCreated(question))
	}
}

func NewFetchNewQuestions() (*FetchNewQuestions, error) {
	svc, err := NewQuestionsService()

	if err != nil {
		return nil, err
	}

	return &FetchNewQuestions{
		svc,
		NewQuestionsRepository(),
		utils.GetApplicationLogger(),
		newEventPublisher(),
	}, nil
}
