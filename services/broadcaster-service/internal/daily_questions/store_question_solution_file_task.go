package daily_questions

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

const TypeStoreQuestionSolutionFileTask = "daily_questions:store_solution_file"

type StoreQuestionSolutionFileTaskPayload struct {
	SolutionId string
}

func NewStoreQuestionSolutionFileTask(solutionId string) (*asynq.Task, error) {
	payload, err := json.Marshal(StoreQuestionSolutionFileTaskPayload{solutionId})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeStoreQuestionSolutionFileTask, payload), nil
}

func StoreQuestionSolutionFileTaskProcessor(ctx context.Context, t *asynq.Task) error {
	var p StoreQuestionSolutionFileTaskPayload
	logger := utils.GetApplicationLogger()
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		logger.Error(err.Error())
		return err
	}

	use_case := NewStoreQuestionSolutionFile()
	use_case.Execute(p.SolutionId)
	return nil
}
