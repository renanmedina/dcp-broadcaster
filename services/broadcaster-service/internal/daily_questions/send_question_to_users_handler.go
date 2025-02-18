package daily_questions

import (
	"github.com/renanmedina/dcp-broadcaster/internal/accounts"
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
)

type SendQuestionToUsersHandler struct {
	usersRepository accounts.UsersRepository
}

func (handler SendQuestionToUsersHandler) Handle(evt event_store.PublishableEvent) {
	questionId := evt.ObjectId()
	use_case := newBroadcastQuestionToUsers()
	go use_case.Execute(questionId)
}

func NewSendQuestionToUsersHandler() SendQuestionToUsersHandler {
	return SendQuestionToUsersHandler{
		accounts.NewUsersRepository(),
	}
}
