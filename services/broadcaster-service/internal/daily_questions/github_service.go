package daily_questions

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/internal/exceptions"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type GithubFileStorageService struct {
	client utils.ApiClient[GithubRepositoryApi]
	logger *utils.ApplicationLogger
}

type GithubRepositoryApi struct {
	Content GithubRepositoryApiResponseContent `json:"content"`
	Commit  GithubRepositoryApiResponseCommit  `json:"commit"`
}

type GithubRepositoryApiResponseContent struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Url         string `json:"url"`
	DownloadUrl string `json:"download_url"`
	Size        int    `json:"size"`
}

type GithubRepositoryApiResponseCommit struct {
	Sha      string   `json:"sha"`
	NodeId   string   `json:"node_id"`
	Message  string   `json:"message"`
	Url      string   `json:"url"`
	Author   Commiter `json:"author"`
	Commiter Commiter `json:"committer"`
}

type Commiter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ErrGithubFileAlreadyExists = errors.New("File already exists")

func (s GithubFileStorageService) SaveFile(filename string, content string, author Commiter, commitMessage string) error {
	if commitMessage == "" {
		commitMessage = fmt.Sprintf("Adding file at %s", filename)
	}

	path := fmt.Sprintf("/contents/%s", filename)
	encoder := base64.StdEncoding
	params := map[string]interface{}{
		"message":   commitMessage,
		"committer": author,
		"content":   encoder.EncodeToString([]byte(content)),
	}

	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
	}

	_, err := s.client.Put(path, params, headers)

	if err != nil {
		var unprocessableError exceptions.HttpUnprocessableEntityError
		if errors.As(err, &unprocessableError) {
			return exceptions.GithubFileAlreadyExistsError{}
		}

		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func NewGithubFileStorageService() GithubFileStorageService {
	configs := utils.GetConfigs()

	return GithubFileStorageService{
		utils.NewApiClient[GithubRepositoryApi](utils.ApiConfig{
			ApiUrl:     configs.GITHUB_REPO_API_URL,
			AuthToken:  configs.GITHUB_API_TOKEN,
			LogEnabled: true,
		}),
		utils.GetApplicationLogger(),
	}
}

func NewGithubCommiter(name string, email string) Commiter {
	return Commiter{name, email}
}
