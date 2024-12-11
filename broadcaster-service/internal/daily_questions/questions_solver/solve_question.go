package questions_solver

import (
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type SolveQuestion struct {
	questionsRepository daily_questions.QuestionsRepository
	solutionsRepository QuestionSolutionsRepository
	logger              *utils.ApplicationLogger
	solver              QuestionSolverService
}

func (uc SolveQuestion) Execute(questionId string, programmingLanguage string) {
	question, err := uc.questionsRepository.GetById(questionId)

	if err != nil {
		uc.logger.Info(fmt.Sprintf("Question %s not found", questionId))
		return
	}

	solution, err := uc.solver.SolveFor(QuestionSolutionRequest{
		question.Text,
		programmingLanguage,
	})

	if err != nil {
		uc.logger.Error(err.Error())
		return
	}

	fmt.Println(solution.content)
	// uc.solutionsRepository.Save(solution.content)
}

func NewSolveQuestion(solver QuestionSolverService) SolveQuestion {
	return SolveQuestion{
		daily_questions.NewQuestionsRepository(),
		NewQuestionSolutionsRepository(),
		utils.GetApplicationLogger(),
		solver,
	}
}
