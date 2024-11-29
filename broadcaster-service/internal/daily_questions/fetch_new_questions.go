package daily_questions

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/lib/pq"
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/monitoring"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type FetchNewQuestions struct {
	service             QuestionsService
	questionsRepository QuestionsRepository
	logger              *utils.ApplicationLogger
	publisher           *event_store.EventPublisher
}

func (uc *FetchNewQuestions) Execute() {
	trace := monitoring.NewTrace("FetchNewQuestions.Execute", context.Background())
	defer trace.End()

	latestQuestion := uc.questionsRepository.GetLatest()
	var fetchQuantity uint32 = 3

	if latestQuestion != nil {
		diff := time.Since(latestQuestion.ReceivedAt)
		fetchQuantity = uint32(math.Ceil(diff.Hours() / 24))
	}

	questions, err := uc.service.GetNewQuestions(fetchQuantity)

	if err != nil {
		uc.logger.Error(err.Error())
		monitoring.ReportErrorFor(trace, err)
	}

	uc.processQuestions(questions, trace)
}

func (uc *FetchNewQuestions) processQuestions(questions []Question, trace monitoring.TraceUnit) {
	for _, question := range questions {
		trace.NewChildSpan(fmt.Sprintf("FetchNewQuestions.processQuestions[%s]", question.Id.String()))

		_, err := uc.questionsRepository.Save(question)

		if err != nil {
			if err, ok := err.(*pq.Error); ok {
				// ignore duplicate inserts of original_id, doing this here do make usage of db constraint and not need to manually load to check if exists
				if err.Code.Name() == UNIQUE_CONSTRAINT_ERROR {
					continue
				}
			}

			errMsg := fmt.Sprintf("Failed processing message received from questions service: %s", err.Error())
			uc.logger.Error(errMsg, "question", question.ToLogMap())
			monitoring.ReportErrorFor(trace, err)
			continue
		}

		uc.logger.Info("Processed message received from questions service", "question", question.ToLogMap())
		trace.NewChildSpan(fmt.Sprintf("event_store.EventPublisher.publish[QuestionCreated][%s]", question.Id.String()))
		uc.publisher.Publish(newQuestionCreated(question))
	}
}

func NewFetchNewQuestions() (*FetchNewQuestions, error) {
	service := NewQuestionsService()

	return &FetchNewQuestions{
		service,
		NewQuestionsRepository(),
		utils.GetApplicationLogger(),
		newEventPublisher(),
	}, nil
}
