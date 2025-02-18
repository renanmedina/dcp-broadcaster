package daily_questions

import (
	"errors"
	"fmt"
	"log"
	"maps"
	"slices"
	"time"

	imap "github.com/BrianLeishman/go-imap"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

const (
	SENDER_LOOKUP_EMAIL = "founders@dailycodingproblem.com"
)

type InternalImapClient struct {
	config    utils.ImapConfigs
	inboxName string
	dialer    *imap.Dialer
}

func (c *InternalImapClient) connect() error {
	if c.dialer == nil {
		imap.Verbose = false
		imapClient, err := imap.New(c.config.Username, c.config.Password, c.config.ServerUrl, c.config.ServerPort)

		if err != nil {
			log.Fatal(err)
			return err
		}

		imapClient.SelectFolder(c.inboxName)
		c.dialer = imapClient
	}

	return nil
}

func (c *InternalImapClient) disconnect() error {
	if c.dialer != nil {
		err := c.dialer.Close()
		c.dialer = nil

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *InternalImapClient) GetEmails(quantity uint32) (map[int]*imap.Email, error) {
	err := c.connect()
	defer c.disconnect()

	emails := make(map[int]*imap.Email, 0)

	if err != nil {
		return emails, errors.Join(errors.New("failed to connect to imap server"), err)
	}

	totalMessages, err := c.dialer.GetTotalEmailCount()
	if err != nil {
		return emails, errors.Join(errors.New("failed to fetch total messages in selected imap folder"), err)
	}

	fetchOffset := totalMessages - int(quantity)
	uids, err := c.dialer.GetUIDs(fmt.Sprintf("%d:%d", fetchOffset, totalMessages))

	if err != nil {
		return emails, errors.Join(errors.New("no uuids found for search"), err)
	}

	emails, err = c.dialer.GetEmails(uids...)

	if err != nil {
		return emails, errors.Join(errors.New("no uuids found for search"), err)
	}

	return emails, nil
}

type QuestionsService struct {
	client InternalImapClient
	logger *utils.ApplicationLogger
}

func (s *QuestionsService) Client() InternalImapClient {
	return s.client
}

func (s *QuestionsService) GetNewQuestions(quantity uint32) ([]Question, error) {
	messages, err := s.fetchMessages(quantity)

	if err != nil {
		return nil, err
	}

	newMessages := buildQuestionsFromMessages(messages)
	return newMessages, nil
}

func (s *QuestionsService) GetQuestionsFromAfter(threshold time.Time) ([]Question, error) {
	messages, err := s.fetchMessages(0)

	if err != nil {
		return nil, err
	}

	newMessages := buildQuestionsFromMessages(messages)
	return newMessages, nil
}

func (s *QuestionsService) fetchMessages(quantity uint32) (map[int]*imap.Email, error) {
	if quantity == 0 {
		quantity = 1
	}

	emails, err := s.client.GetEmails(quantity)

	if err != nil {
		s.logger.Error(err.Error(), "error", err.Error())
		return make(map[int]*imap.Email), err
	}

	return emails, nil
}

func NewQuestionsService() QuestionsService {
	config := utils.GetImapConfigs()
	return QuestionsService{
		client: InternalImapClient{
			config:    *config,
			inboxName: "INBOX",
		},
		logger: utils.GetApplicationLogger(),
	}
}

func buildQuestionsFromMessages(messages map[int]*imap.Email) []Question {
	var newMessages []Question

	for _, msg := range messages {
		address := slices.Collect(maps.Keys(msg.From))

		if address[0] == SENDER_LOOKUP_EMAIL {
			metadata := parseQuestionEmailMessage(msg)

			if metadata.Valid() {
				question := NewQuestionFromEmailMetadata(metadata)
				newMessages = append(newMessages, question)
			}
		}
	}

	return newMessages
}
