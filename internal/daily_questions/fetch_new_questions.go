package daily_questions

import (
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

type FetchNewQuestions struct {
	service             QuestionsService
	questionsRepository QuestionsRepository
	logger              *utils.ApplicationLogger
}

func (uc *FetchNewQuestions) Execute() {
	questions, err := uc.service.GetNewQuestions()

	if err != nil {
		uc.logger.Error(err.Error())
	}

	for _, question := range questions {
		fmt.Println("ID:", question.Id)
		fmt.Println("Title:", question.Title)
		fmt.Println("Difficulty:", question.DifficultyLevel)
		fmt.Println("CompanyName:", question.CompanyName)
		fmt.Println("Date:", question.ReceivedAt)
		fmt.Println("------------------------------------------------------------------------------------------------------")
		_, err := uc.questionsRepository.Save(question)
		if err != nil {
			uc.logger.Fatal(err.Error())
		}
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
	}, nil
}
