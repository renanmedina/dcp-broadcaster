package daily_questions

import (
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

func (r *QuestionSolutionsRepository) GetById(solutionId string) (*QuestionSolution, error) {
	var solution QuestionSolution
	result := r.db.Table(QUESTIONS_SOLUTION_TABLE_NAME).First(&solution, "id = ?", solutionId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &solution, nil
}

func (r *QuestionSolutionsRepository) GetAll() ([]QuestionSolution, error) {
	var solutions []QuestionSolution
	result := r.db.Find(&solutions)

	if result.Error != nil {
		return make([]QuestionSolution, 0), result.Error
	}

	return solutions, nil
}

func (r *QuestionSolutionsRepository) Save(solution QuestionSolution) (*QuestionSolution, error) {
	result := r.db.Save(&solution)

	if result.Error != nil {
		return nil, result.Error
	}

	return &solution, nil
}

func NewQuestionSolutionsRepository() QuestionSolutionsRepository {
	return QuestionSolutionsRepository{
		utils.GetDatabaseConnection(),
		utils.GetApplicationLogger(),
	}
}
