package broadcasting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/renanmedina/dcp-broadcaster/internal/accounts"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

const BROADCAST_COMMUNITY_GROUP_CHAT_ID = "120363316303547295@g.us"

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
	phone_number := fmt.Sprintf("%s@c.us", user.PhoneNumber)
	s.logger.Info(fmt.Sprintf("Sending received daily question via whatsapp to phone number %s", phone_number), "user", user.ToLogMap(), "message", message)

	err := s.sendMessageRequest(message, phone_number)

	if err != nil {
		errMsg := fmt.Sprintf("Failed send daily question via whatsapp to phone number %s", phone_number)
		s.logger.Error(errMsg, "user", user.ToLogMap(), "message", message)
		return err
	}

	return nil
}

func (s WhatsappService) Broadcast(message string) error {
	s.logger.Info("Sending received daily question via whatsapp to community group", "message", message)
	err := s.sendMessageRequest(message, BROADCAST_COMMUNITY_GROUP_CHAT_ID)

	if err != nil {
		s.logger.Error("Failed send daily question via whatsapp to community group", "message", message)
		return err
	}

	return nil
}

func (s WhatsappService) sendMessageRequest(message string, chatId string) error {
	err := s.restartSession()
	if err != nil {
		s.logger.Error("Failed to restart session", "error", err.Error())
		return err
	}

	time.Sleep(time.Second * 45)

	url := fmt.Sprintf("%s/client/sendMessage/%s", s.configs.apiUrl, s.configs.sessionId)
	bodyParams, err := json.Marshal(map[string]string{
		"chatId":      chatId,
		"contentType": "string",
		"content":     message,
	})

	if err != nil {
		s.logger.Error("Failed marshaling params to send whatsapp message")
		return err
	}

	requestParams := bytes.NewBuffer(bodyParams)
	request, err := http.NewRequest("POST", url, requestParams)

	request.Header.Add("Accept", "*/*")
	request.Header.Add("x-api-key", s.configs.apiToken)
	request.Header.Add("Content-Type", "application/json")

	if err != nil {
		s.logger.Error("Failed creating request to send whatsapp message")
		return err
	}

	response, err := s.client.Do(request)

	if err != nil {
		return err
	}

	_, err = io.ReadAll(response.Body)

	if err != nil {
		s.logger.Error("Failed to read response from whatsapp service")
		return err
	}

	return nil
}

func (s WhatsappService) restartSession() error {
	url := fmt.Sprintf("%s/session/restart/%s", s.configs.apiUrl, s.configs.sessionId)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	request.Header.Add("Accept", "*/*")
	request.Header.Add("x-api-key", s.configs.apiToken)
	request.Header.Add("Content-Type", "application/json")

	_, err = s.client.Do(request)

	if err != nil {
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
