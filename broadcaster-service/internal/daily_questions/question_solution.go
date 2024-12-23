package daily_questions

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions/questions_solver"
	"gorm.io/gorm"
)

type QuestionSolution struct {
	gorm.Model
	Id                  uuid.UUID `gorm:"primaryKey"`
	DailyQuestionId     uuid.UUID
	ProgrammingLanguage string
	SolutionCode        string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
}

// gorm before create hook
func (q *QuestionSolution) BeforeCreate(tx *gorm.DB) (err error) {
	q.Id = uuid.New()
	return nil
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
	extension := s.ProgrammingLanguage
	languageInfo, exists := questions_solver.SolvingLanguages[s.ProgrammingLanguage]

	if exists {
		extension = languageInfo.FileExtension
	}

	return fmt.Sprintf("solution.%s", extension)
}

func (s QuestionSolution) FileContent() string {
	return s.SolutionCode
}

func newQuestionSolution(questionId uuid.UUID, programmingLanguage string, solutionCode string) QuestionSolution {
	return QuestionSolution{
		DailyQuestionId:     questionId,
		ProgrammingLanguage: programmingLanguage,
		SolutionCode:        solutionCode,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}
