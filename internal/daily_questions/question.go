package daily_questions

import (
	"time"

	"github.com/emersion/go-imap"
)

type Question struct {
	Id         int
	Title      string
	EmailBody  string
	Text       string
	ReceivedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewQuestionFromEmailMessage(msg *imap.Message) Question {
	return Question{
		Id:         int(msg.Uid),
		Title:      msg.Envelope.Subject,
		EmailBody:  "",
		Text:       parseQuestionFromHtml(""),
		ReceivedAt: msg.Envelope.Date,
	}
}

func parseQuestionFromHtml(bodyHtml string) string {
	return ""
}
