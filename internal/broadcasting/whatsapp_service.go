package broadcasting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/renanmedina/dcp-broadcaster/internal/accounts"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

type WhatsappConfigs struct {
	apiUrl    string
	apiToken  string
	sessionId string
}

type WhatsappService struct {
	configs WhatsappConfigs
	client  http.Client
	logger  *utils.ApplicationLogger
}

func (s WhatsappService) Send(message string, user accounts.User) error {
	phone_number := user.PhoneNumber
	s.logger.Info(fmt.Sprintf("Sending received daily question via whatsapp to phone number %s", phone_number), "user", user.ToLogMap(), "message", message)

	url := fmt.Sprintf("%s/client/sendMessage/%s", s.configs.apiUrl, s.configs.sessionId)
	bodyParams, err := json.Marshal(map[string]string{
		"chatId":      fmt.Sprintf("%s@c.us", phone_number),
		"contentType": "string",
		"content":     message,
	})

	if err != nil {
		s.logger.Error("Failed marshaling params to send whatsapp message")
		return err
	}

	params := bytes.NewBuffer(bodyParams)

	request, err := http.NewRequest("POST", url, params)
	request.Header.Add("Accept", "*/*")
	request.Header.Add("x-api-key", s.configs.apiToken)
	request.Header.Add("Content-Type", "application/json")

	if err != nil {
		s.logger.Error("Failed creating request to send whatsapp message")
		return err
	}

	response, err := s.client.Do(request)

	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed send daily question via whatsapp to phone number %s", phone_number), "user", user.ToLogMap(), "message", message)
		return err
	}

	_, err = io.ReadAll(response.Body)

	if err != nil {
		s.logger.Error("Failed to read response from whatsapp service")
		return err
	}

	return nil
}

func NewWhatsappService() WhatsappService {
	configs := utils.GetConfigs()

	return WhatsappService{
		WhatsappConfigs{
			configs.WHATSAPP_API_URL,
			configs.WHATSAPP_API_TOKEN,
			configs.WHATSAPP_SESSION_ID,
		},
		*http.DefaultClient,
		utils.GetApplicationLogger(),
	}
}
