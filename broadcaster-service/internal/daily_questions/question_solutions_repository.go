package daily_questions

import "github.com/renanmedina/dcp-broadcaster/utils"

type QuestionSolutionsRepository struct {
	db     *utils.DatabaseAdapdater
	logger *utils.ApplicationLogger
}

const QUESTIONS_SOLUTION_TABLE_NAME = "daily_questions_solutions"

func (r *QuestionSolutionsRepository) Save(solution QuestionSolution) (*QuestionSolution, error) {
	if !solution.Persisted {
		_, err := r.db.Insert(QUESTIONS_SOLUTION_TABLE_NAME, solution.ToDbMap())

		if err != nil {
			return nil, err
		}
	} else {
		_, err := r.db.UpdateById(QUESTIONS_SOLUTION_TABLE_NAME, solution.Id.String(), solution.ToDbMap())

		if err != nil {
			return nil, err
		}
	}

	solution.Persisted = true
	return &solution, nil
}

func NewQuestionSolutionsRepository() QuestionSolutionsRepository {
	return QuestionSolutionsRepository{
		utils.GetDatabase(),
		utils.GetApplicationLogger(),
	}
}
