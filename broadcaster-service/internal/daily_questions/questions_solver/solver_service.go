package questions_solver

import "fmt"

type QuestionSolverService interface {
	// SolveByText(questionText string) (SolutionResponse, error)
	SolveFor(request SolveQuestionRequest) (SolutionResponse, error)
}

type SolveQuestionRequest struct {
	QuestionContent    string
	ProgrammingLanguge string
}

func (r SolveQuestionRequest) Prompt() string {
	return fmt.Sprintf(
		"%s can you solve using %s language? return ONLY the solution enclosed in markdown for that language without any example use case",
		r.QuestionContent,
		r.ProgrammingLanguge,
	)
}

type SolutionResponse struct {
	Content string
}
