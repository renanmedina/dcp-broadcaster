package daily_questions

import (
	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions/questions_solver"
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type SolveQuestionHandler struct {
	logger *utils.ApplicationLogger
}

func (handler SolveQuestionHandler) Handle(evt event_store.PublishableEvent) {
	questionId := evt.ObjectId()
	use_case := NewSolveQuestion(questions_solver.NewOllamaService())
	// languageRepository := questions_solver.NewSolveLanguagesRepository()
	// solvingLanguages, err := languageRepository.GetAllEnabled()

	// if err != nil {
	// 	handler.logger.Error(err.Error())
	// 	return
	// }

	for _, language := range questions_solver.SolvingLanguages {
		go use_case.Execute(questionId, language.LanguageName)
	}
}

func NewSolveQuestionHandler() SolveQuestionHandler {
	return SolveQuestionHandler{
		utils.GetApplicationLogger(),
	}
}
