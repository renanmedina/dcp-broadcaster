package daily_questions

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions/questions_solver"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

const TypeSolveQuestionTask = "daily_questions:solve_question"

type SolveQuestionTaskPayload struct {
	QuestionId          string
	ProgrammingLanguage string
}

func NewSolveQuestionTask(questionId string, programmingLanguage string) (*asynq.Task, error) {
	payload, err := json.Marshal(SolveQuestionTaskPayload{questionId, programmingLanguage})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeSolveQuestionTask, payload), nil
}

func SolveQuestionTaskProcessor(ctx context.Context, t *asynq.Task) error {
	var p SolveQuestionTaskPayload
	logger := utils.GetApplicationLogger()
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		logger.Error(err.Error())
		return err
	}

	use_case := NewSolveQuestion(questions_solver.NewOllamaService())
	use_case.Execute(p.QuestionId, p.ProgrammingLanguage)
	return nil
}
