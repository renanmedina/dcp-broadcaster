package daily_questions

import (
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/internal/daily_questions/questions_solver"
	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type SolveQuestion struct {
	questionsRepository QuestionsRepository
	solutionsRepository QuestionSolutionsRepository
	logger              *utils.ApplicationLogger
	solver              questions_solver.QuestionSolverService
	eventPublisher      *event_store.EventPublisher
}

func (uc SolveQuestion) Execute(questionId string, programmingLanguage string) {
	question, err := uc.questionsRepository.GetById(questionId)

	if err != nil {
		uc.logger.Info(fmt.Sprintf("Question %s not found", questionId))
		return
	}

	solvedQuestion, err := uc.solver.SolveFor(questions_solver.QuestionSolutionRequest{question.Text, programmingLanguage})

	if err != nil {
		uc.logger.Error(err.Error())
		return
	}

	questionSolution := newQuestionSolution(question.Id, programmingLanguage, solvedQuestion.Content)
	_, err = uc.solutionsRepository.Save(questionSolution)

	if err != nil {
		uc.logger.Error(err.Error())
		return
	}

	event := newQuestionSolved(*question, questionSolution)
	uc.eventPublisher.Publish(event)
}

func NewSolveQuestion(solver questions_solver.QuestionSolverService) SolveQuestion {
	return SolveQuestion{
		NewQuestionsRepository(),
		NewQuestionSolutionsRepository(),
		utils.GetApplicationLogger(),
		solver,
		newEventPublisher(),
	}
}
