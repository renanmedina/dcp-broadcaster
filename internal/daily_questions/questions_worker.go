package daily_questions

import (
	"fmt"
	"log"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

type QuestionsWorker struct {
	questionsService QuestionsService
	logger           *utils.ApplicationLogger
}

func (r *QuestionsWorker) Work() {
	questions, err := r.questionsService.GetNewMessages()

	if err != nil {
		r.logger.Error(err.Error())
	}

	repo := NewQuestionsRepository()

	for _, question := range questions {
		fmt.Println("ID:", question.Id)
		fmt.Println("Title:", question.Title)
		fmt.Println("Difficulty:", question.DifficultyLevel)
		fmt.Println("CompanyName:", question.CompanyName)
		fmt.Println("Date:", question.ReceivedAt)
		fmt.Println("------------------------------------------------------------------------------------------------------")
		_, err := repo.Save(question)
		if err != nil {
			log.Fatal(err.Error())
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
