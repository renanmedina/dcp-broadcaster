package questions_solver

import (
	"strings"
	"time"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

const OLLAMA_MODEL_NAME = "llama3.2"

type OllamaService struct {
	client utils.ApiClient[OllamaApiResponse]
	logger *utils.ApplicationLogger
}

type OllamaApiResponse struct {
	ModelName       string    `json:"model"`
	CreatedAt       time.Time `json:"created_at"`
	ResponseContent string    `json:"response"`
	Done            bool      `json:"done"`
	DoneReason      string    `json:"done_reason"`
	ContextInfo     []int     `json:"context"`
}

func (s OllamaService) SolveByText(prompt string) (SolutionResponse, error) {
	preparedPrompt := preparePromptString(prompt)
	s.logger.Info(preparedPrompt)

	response, err := s.client.Post("/generate", map[string]string{
		"model":  OLLAMA_MODEL_NAME,
		"prompt": preparedPrompt,
		"stream": "false",
	}, make(map[string]string))

	if err != nil {
		return SolutionResponse{}, err
	}

	return SolutionResponse{response.ResponseContent}, nil
}

func (s OllamaService) SolveFor(question QuestionSolutionRequest) (SolutionResponse, error) {
	byTextResponse, err := s.SolveByText(question.Prompt())

	if err != nil {
		return SolutionResponse{}, err
	}

	solutionCode, err := extractSolutionCodeFrom(byTextResponse.content, question.programmingLanguge)

	if err != nil {
		return SolutionResponse{}, err
	}

	return SolutionResponse{solutionCode}, nil
}

func NewOllamaService() OllamaService {
	configs := utils.GetConfigs()

	return OllamaService{
		utils.NewApiClient[OllamaApiResponse](utils.ApiConfig{
			ApiUrl:     configs.OLLAMA_SERVICE_API_URL,
			AuthToken:  "",
			LogEnabled: true,
		}),
		utils.GetApplicationLogger(),
	}
}

func preparePromptString(prompt string) string {
	prepared := strings.ReplaceAll(prompt, "\n", " ")
	prepared = strings.ReplaceAll(prepared, "\r", "")
	return strings.ReplaceAll(prepared, "'", "")
}
