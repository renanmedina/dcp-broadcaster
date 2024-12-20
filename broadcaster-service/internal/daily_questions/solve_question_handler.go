package daily_questions

import (
	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions/questions_solver"
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
)

type SolveQuestionHandler struct{}

func (handler SolveQuestionHandler) Handle(evt event_store.PublishableEvent) {
	questionId := evt.ObjectId()
	use_case := NewSolveQuestion(questions_solver.NewOllamaService())

	for _, language := range questions_solver.SolvingLanguages {
		go use_case.Execute(questionId, language.LanguageName)
	}
}

func NewSolveQuestionHandler() SolveQuestionHandler {
	return SolveQuestionHandler{}
}
