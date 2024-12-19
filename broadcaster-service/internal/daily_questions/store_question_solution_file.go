package daily_questions

import (
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/internal/event_store"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type StoreQuestionSolutionFile struct {
	solutionsRepository QuestionSolutionsRepository
	githubService       GithubFileStorageService
	logger              *utils.ApplicationLogger
	eventPublisher      *event_store.EventPublisher
}

func (uc StoreQuestionSolutionFile) Execute(solutionId string) {
	solution, err := uc.solutionsRepository.GetById(solutionId)

	if err != nil {
		uc.logger.Info(fmt.Sprintf("Solution %s not found", solutionId))
		return
	}

	uc.ExecuteFor(solution)
}

func (uc StoreQuestionSolutionFile) ExecuteFor(solution *QuestionSolution) error {
	filename := fmt.Sprintf("dcp-solutions/%s/%s", solution.DailyQuestionId, solution.Filename())
	commiter := NewGithubCommiter("dcp-broadcaster", "dcp-broadcaster@silvamedina.com.br")
	err := uc.githubService.SaveFile(filename, solution.FileContent(), commiter)

	if err != nil {
		uc.logger.Info(fmt.Sprintf("Solution %s not found", solution.Id))
		return err
	}

	return nil
}

func NewStoreQuestionSolutionFile() StoreQuestionSolutionFile {
	return StoreQuestionSolutionFile{
		NewQuestionSolutionsRepository(),
		NewGithubFileStorageService(),
		utils.GetApplicationLogger(),
		newEventPublisher(),
	}
}
