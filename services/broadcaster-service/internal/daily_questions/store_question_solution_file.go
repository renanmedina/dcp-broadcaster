package daily_questions

import (
	"errors"
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/internal/exceptions"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type StoreQuestionSolutionFile struct {
	solutionsRepository QuestionSolutionsRepository
	questionsRepository QuestionsRepository
	githubService       GithubFileStorageService
	logger              *utils.ApplicationLogger
}

func (uc StoreQuestionSolutionFile) Execute(solutionId string) error {
	solution, err := uc.solutionsRepository.GetById(solutionId)

	if err != nil {
		uc.logger.Info(err.Error())
		return err
	}

	return uc.ExecuteFor(solution)
}

func (uc StoreQuestionSolutionFile) ExecuteFor(solution *QuestionSolution) error {
	question, err := uc.questionsRepository.GetById(solution.DailyQuestionId)
	if err != nil {
		uc.logger.Error(err.Error())
		return err
	}

	commiter := NewGithubCommiter("dcp-solver", "dcp-solver@silvamedina.com.br")
	questionDateFormatted := question.ReceivedAt.Format("2006-01-02")

	question_filename := fmt.Sprintf("dcp-solutions/%s/%s", questionDateFormatted, "README.md")
	uc.storeFile(question_filename, question.Text, commiter)

	solution_filename := fmt.Sprintf("dcp-solutions/%s/%s", questionDateFormatted, solution.Filename())
	return uc.storeFile(solution_filename, solution.FileContent(), commiter)
}

func (uc StoreQuestionSolutionFile) storeFile(filepath string, content string, commiter Commiter) error {
	err := uc.githubService.SaveFile(filepath, content, commiter, "")

	if err != nil {
		var fileExits exceptions.GithubFileAlreadyExistsError
		if errors.As(err, &fileExits) {
			return nil
		}

		uc.logger.Error(err.Error())
		return err
	}

	return nil
}

func NewStoreQuestionSolutionFile() StoreQuestionSolutionFile {
	return StoreQuestionSolutionFile{
		NewQuestionSolutionsRepository(),
		NewQuestionsRepository(),
		NewGithubFileStorageService(),
		utils.GetApplicationLogger(),
	}
}
