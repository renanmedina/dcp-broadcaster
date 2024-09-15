package daily_questions

import (
	"github.com/renanmedina/dcp-broadcaster/internal/accounts"
	"github.com/renanmedina/dcp-broadcaster/internal/broadcasting"
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type SendQuestionToUser struct {
	userRepository      accounts.UsersRepository
	questionsRepository QuestionsRepository
	messagingService    broadcasting.MessageService
	logger              *utils.ApplicationLogger
	publisher           *event_store.EventPublisher
}

func (uc *SendQuestionToUser) Execute(questionId string, userId string) {
	question, err := uc.questionsRepository.GetById(questionId)

	if err != nil {
		uc.logger.Error(err.Error())
	}

	user, err := uc.userRepository.GetById(userId)

	if err != nil {
		uc.logger.Error(err.Error())
	}

	err = uc.messagingService.Send(question.Text, user)

	if err != nil {
		uc.logger.Error(err.Error())
		return
	}

	uc.publisher.Publish(newQuestionSentToUser(*question, user))
}

func NewSendQuestionToUser() SendQuestionToUser {
	return SendQuestionToUser{
		accounts.NewUsersRepository(),
		NewQuestionsRepository(),
		broadcasting.NewWhatsappService(),
		utils.GetApplicationLogger(),
		newEventPublisher(),
	}
}
