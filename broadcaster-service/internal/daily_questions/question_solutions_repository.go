package daily_questions

import (
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/utils"
	"gorm.io/gorm"
)

type QuestionSolutionsRepository struct {
	db     *gorm.DB
	logger *utils.ApplicationLogger
}

const (
	QUESTIONS_SOLUTION_TABLE_NAME = "daily_questions_solutions"
)

type QuestionSolutionNotFound struct {
	msg string
}

func (e QuestionSolutionNotFound) Error() string {
	return e.msg
}

func NewQuestionSolutionNotFound(msg string) QuestionSolutionNotFound {
	return QuestionSolutionNotFound{msg}
}

func (r *QuestionSolutionsRepository) GetById(solutionId string) (*QuestionSolution, error) {
	var solution QuestionSolution
	result := r.db.Table(QUESTIONS_SOLUTION_TABLE_NAME).First(&solution, "id = ?", solutionId)

	if result.Error != nil {
		return nil, NewQuestionSolutionNotFound(fmt.Sprintf("Solution %s not found", solutionId))
	}

	return &solution, nil
}

func (r *QuestionSolutionsRepository) GetAll() ([]QuestionSolution, error) {
	var solutions []QuestionSolution
	result := r.db.Table(QUESTIONS_SOLUTION_TABLE_NAME).Find(&solutions)

	if result.Error != nil {
		return make([]QuestionSolution, 0), result.Error
	}

	return solutions, nil
}

func (r *QuestionSolutionsRepository) Save(solution QuestionSolution) (QuestionSolution, error) {
	result := r.db.Table(QUESTIONS_SOLUTION_TABLE_NAME).Save(&solution)

	if result.Error != nil {
		return QuestionSolution{}, result.Error
	}

	return solution, nil
}

func NewQuestionSolutionsRepository() QuestionSolutionsRepository {
	return QuestionSolutionsRepository{
		utils.GetDatabaseConnection(),
		utils.GetApplicationLogger(),
	}
}
