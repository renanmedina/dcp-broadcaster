package questions_solver

import "github.com/renanmedina/dcp-broadcaster/utils"

type QuestionSolutionsRepository struct {
	logger *utils.ApplicationLogger
}

func NewQuestionSolutionsRepository() QuestionSolutionsRepository {
	return QuestionSolutionsRepository{
		utils.GetApplicationLogger(),
	}
}
