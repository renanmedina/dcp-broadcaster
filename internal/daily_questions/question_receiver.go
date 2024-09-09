package daily_questions

import (
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

type QuestionsReceiver struct {
	questionsService QuestionsService
	logger           *utils.ApplicationLogger
}

func (r *QuestionsReceiver) Work() {
	questions, err := r.questionsService.GetNewMessages()

	if err != nil {
		r.logger.Error(err.Error())
	}

	for _, question := range questions {
		fmt.Println("ID:", question.Id)
		fmt.Println("Title:", question.Title)
		fmt.Println("Date:", question.ReceivedAt)
		// fmt.Println("Body:", msg.Body)
	}
}

func NewQuestionsReceiver() (QuestionsReceiver, error) {
	service, err := NewQuestionsService()

	if err != nil {
		return QuestionsReceiver{}, err
	}

	return QuestionsReceiver{service, utils.GetApplicationLogger()}, nil
}
