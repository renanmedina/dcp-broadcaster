package daily_questions

import (
	"github.com/renanmedina/dcp-broadcaster/internal/broadcasting"
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type BroadcastQuestionToUsers struct {
	questionsRepository QuestionsRepository
	messagingService    broadcasting.MessageService
	logger              *utils.ApplicationLogger
	publisher           *event_store.EventPublisher
}

func (uc *BroadcastQuestionToUsers) Execute(questionId string) {
	question, err := uc.questionsRepository.GetById(questionId)

	if err != nil {
		uc.logger.Error(err.Error())
	}

	err = uc.messagingService.Broadcast(question.Text)

	if err != nil {
		uc.logger.Error(err.Error())
		return
	}

	uc.publisher.Publish(newQuestionBroadcastedToUsers(*question))
}

func newBroadcastQuestionToUsers() BroadcastQuestionToUsers {
	return BroadcastQuestionToUsers{
		NewQuestionsRepository(),
		broadcasting.NewWhatsappService(),
		utils.GetApplicationLogger(),
		newEventPublisher(),
	}
}
