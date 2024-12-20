package daily_questions

import (
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type QuestionSolutionsRepository struct {
	db     *utils.DatabaseAdapdater
	logger *utils.ApplicationLogger
}

const (
	QUESTIONS_SOLUTION_TABLE_NAME = "daily_questions_solutions"
	QUESTIONS_SOLUTION_FIELDS     = "id, daily_question_id, programming_language, solution_code, created_at, updated_at"
)

func (r *QuestionSolutionsRepository) GetById(solutionId string) (*QuestionSolution, error) {
	scanner := r.db.SelectOne(QUESTIONS_SOLUTION_FIELDS, QUESTIONS_SOLUTION_TABLE_NAME, map[string]interface{}{
		"id": solutionId,
	})

	solution, err := buildFromDb(*scanner)

	if err != nil {
		return nil, err
	}

	return &solution, err
}

func (r *QuestionSolutionsRepository) GetAll() ([]QuestionSolution, error) {
	rows, err := r.db.Select(QUESTIONS_SOLUTION_FIELDS, QUESTIONS_SOLUTION_TABLE_NAME, nil)
	defer rows.Close()

	if err != nil {
		return make([]QuestionSolution, 0), err
	}

	var solutions []QuestionSolution

	for rows.Next() {
		var solution QuestionSolution
		err = rows.Scan(
			&solution.Id,
			&solution.DailyQuestionId,
			&solution.ProgrammingLanguage,
			&solution.SolutionCode,
			&solution.CreatedAt,
			&solution.UpdatedAt,
		)

		if err != nil {
			panic(err.Error())
		}

		solutions = append(solutions, solution)
	}

	return solutions, nil
}

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

func buildFromDb(dbRow squirrel.RowScanner) (QuestionSolution, error) {
	var solution QuestionSolution
	err := dbRow.Scan(
		&solution.Id,
		&solution.DailyQuestionId,
		&solution.ProgrammingLanguage,
		&solution.SolutionCode,
		&solution.CreatedAt,
		&solution.UpdatedAt,
	)

	if err != nil {
		return QuestionSolution{}, err
	}

	if solution.Id.String() == "00000000-0000-0000-0000-000000000000" {
		return QuestionSolution{}, errors.New("can't find QuestionSolution")
	}

	solution.Persisted = true
	return solution, nil
}
