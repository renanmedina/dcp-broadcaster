package questions_solver

import (
	"fmt"
	"strings"
	"time"

	"github.com/renanmedina/dcp-broadcaster/utils"
)

const OLLAMA_MODEL_NAME = "llama3.2"

type OllamaService struct {
	client utils.ApiClient[OllamaApi]
	logger *utils.ApplicationLogger
}

type OllamaApi struct {
	ModelName       string    `json:"model"`
	CreatedAt       time.Time `json:"created_at"`
	ResponseContent string    `json:"response"`
	Done            bool      `json:"done"`
	DoneReason      string    `json:"done_reason"`
	ContextInfo     []int     `json:"context"`
}

func (s OllamaService) SolveByText(prompt string) (SolutionResponse, error) {
	preparedPrompt := preparePromptString(prompt)
	s.logger.Info(fmt.Sprintf("Sending prompt to Ollama AI %s", prompt))
	response, err := s.client.Post("/generate", map[string]interface{}{
		"model":  OLLAMA_MODEL_NAME,
		"prompt": preparedPrompt,
		"stream": false,
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

	solutionCode := byTextResponse.content
	s.logger.Info(fmt.Sprintf("Ollama Service successfully solved the question with %s", question.programmingLanguge))
	return SolutionResponse{solutionCode}, nil
}

func NewOllamaService() OllamaService {
	configs := utils.GetConfigs()

	return OllamaService{
		utils.NewApiClient[OllamaApi](utils.ApiConfig{
			ApiUrl:     configs.OLLAMA_SERVICE_API_URL,
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
