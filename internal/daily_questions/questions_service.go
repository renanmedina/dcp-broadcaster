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
	client *client.Client
}

func (s *QuestionsService) GetNewMessages() ([]Question, error) {
	mailbox, err := s.client.Select("INBOX", true)

	if err != nil {
		return nil, err
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddRange(1, mailbox.Messages)

	messages := make(chan *imap.Message, 10)
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope}

	go func() {
		if err := s.client.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

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

	return newMessages, nil
}

func (s *QuestionsService) GetMessagesFromAfter(threshold time.Time) ([]Question, error) {
	return nil, nil
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

	return QuestionsService{imapClient}, nil
}
