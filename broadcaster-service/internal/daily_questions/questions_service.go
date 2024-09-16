package daily_questions

import (
	"log"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

const (
	SENDER_LOOKUP_EMAIL = "founders@dailycodingproblem.com"
)

type QuestionsService struct {
	client    *client.Client
	inboxName string
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
	if quantity == 0 {
		quantity = 1
	}

	mailbox, err := s.client.Select(s.inboxName, true)

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
		if err := s.client.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	return messages, nil
}

func NewQuestionsService() (QuestionsService, error) {
	config := utils.GetImapConfigs()
	imapClient, err := client.DialTLS(config.Address(), nil)

	if err != nil {
		return QuestionsService{}, err
	}

	if err := imapClient.Login(config.Username, config.Password); err != nil {
		return QuestionsService{}, err
	}

	return QuestionsService{imapClient, "INBOX"}, nil
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
