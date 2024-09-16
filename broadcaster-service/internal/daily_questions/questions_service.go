package daily_questions

import (
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

const (
	SENDER_LOOKUP_EMAIL = "founders@dailycodingproblem.com"
)

type InternalImapConnector struct {
	config utils.ImapConfigs
	client *client.Client
}

func (c *InternalImapConnector) connect() error {
	if c.client == nil {
		imapClient, err := client.DialTLS(c.config.Address(), nil)

		if err != nil {
			return err
		}

		if err := c.client.Login(c.config.Username, c.config.Password); err != nil {
			return err
		}

		c.client = imapClient
	}

	return nil
}

func (c *InternalImapConnector) disconnect() error {
	if c.client != nil {
		err := c.client.Close()
		c.client = nil

		if err != nil {
			return err
		}
	}

	return nil
}

type QuestionsService struct {
	connector InternalImapConnector
	inboxName string
	logger    *utils.ApplicationLogger
}

func (s *QuestionsService) Client() *client.Client {
	return s.connector.client
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

func (s *QuestionsService) fetchMessages(quantity uint32) (chan *imap.Message, error) {
	err := s.connector.connect()

	defer s.connector.disconnect()

	if err != nil {
		return nil, err
	}

	if quantity == 0 {
		quantity = 1
	}

	mailbox, err := s.Client().Select(s.inboxName, true)

	if err != nil {
		return nil, err
	}

	seqSet := new(imap.SeqSet)
	realQuantity := (quantity - 1)
	seqSet.AddRange(mailbox.Messages, mailbox.Messages-realQuantity)

	messages := make(chan *imap.Message, quantity)
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope}

	go func() {
		if err := s.Client().Fetch(seqSet, items, messages); err != nil {
			s.logger.Error("Failed to fetch messages from imap client", "error", err.Error())
		}
	}()

	return messages, nil
}

func NewQuestionsService() QuestionsService {
	config := utils.GetImapConfigs()
	return QuestionsService{
		connector: InternalImapConnector{
			config: *config,
		},
		inboxName: "INBOX",
		logger:    utils.GetApplicationLogger(),
	}
}

func buildQuestionsFromMessages(messages chan *imap.Message) []Question {
	var newMessages []Question

	for msg := range messages {
		if msg.Envelope.From[0].Address() == SENDER_LOOKUP_EMAIL {
			metadata := parseQuestionEmailMessage(msg)

			if metadata.Valid() {
				question := NewQuestionFromEmailMetadata(metadata)
				newMessages = append(newMessages, question)
			}
		}
	}

	return newMessages
}
