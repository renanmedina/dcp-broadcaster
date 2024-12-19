package daily_questions

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type QuestionSolution struct {
	Id                  uuid.UUID
	DailyQuestionId     uuid.UUID
	ProgrammingLanguage string
	SolutionCode        string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
	Persisted           bool
}

func (s QuestionSolution) ToDbMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                   s.Id,
		"daily_question_id":    s.DailyQuestionId,
		"programming_language": s.ProgrammingLanguage,
		"solution_code":        s.SolutionCode,
	}
}

func (s QuestionSolution) Filename() string {
	extensions := map[string]string{
		"golang": "go",
		"ruby":   "rb",
		"php":    "php",
		"python": "py",
	}

	extension, exists := extensions[s.ProgrammingLanguage]

	if !exists {
		extension = s.ProgrammingLanguage
	}

	return fmt.Sprintf("solution.%s", extension)
}

func (s QuestionSolution) FileContent() string {
	return s.SolutionCode
}

func newQuestionSolution(questionId uuid.UUID, programmingLanguage string, solutionCode string) QuestionSolution {
	return QuestionSolution{
		Id:                  uuid.New(),
		DailyQuestionId:     questionId,
		ProgrammingLanguage: programmingLanguage,
		SolutionCode:        solutionCode,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}
