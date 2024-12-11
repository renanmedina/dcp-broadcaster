package questions_solver

import "fmt"

type QuestionSolverService interface {
	// SolveByText(questionText string) (SolutionResponse, error)
	SolveFor(request QuestionSolutionRequest) (SolutionResponse, error)
}

type QuestionSolutionRequest struct {
	questionContent    string
	programmingLanguge string
}

func (r QuestionSolutionRequest) Prompt() string {
	return fmt.Sprintf(
		"%s can you solve using %s language?",
		r.questionContent,
		r.programmingLanguge,
	)
}

type SolutionResponse struct {
	content string
}
