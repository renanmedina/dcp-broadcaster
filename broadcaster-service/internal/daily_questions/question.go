package daily_questions

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	Id              uuid.UUID
	OriginalId      string
	DifficultyLevel string
	Title           string
	EmailBody       string
	Text            string
	CompanyName     string
	ReceivedAt      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Persisted       bool
}

func (q *Question) ToDbMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                  q.Id,
		"original_id":         q.OriginalId,
		"difficulty_level":    q.DifficultyLevel,
		"title":               q.Title,
		"question_email_body": q.EmailBody,
		"question_text":       q.Text,
		"company_name":        q.CompanyName,
		"received_at":         q.ReceivedAt,
	}
}

func (q *Question) ToLogMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               q.Id,
		"original_id":      q.OriginalId,
		"difficulty_level": q.DifficultyLevel,
		"title":            q.Title,
		"question_text":    q.Text,
		"company_name":     q.CompanyName,
		"received_at":      q.ReceivedAt,
	}
}

func NewQuestionFromEmailMetadata(metadata QuestionEmailMetadata) Question {
	return Question{
		Id:              uuid.New(),
		OriginalId:      metadata.MessageId,
		Title:           metadata.Title,
		EmailBody:       metadata.BodyHtml,
		Text:            metadata.BodyText,
		ReceivedAt:      metadata.Date,
		DifficultyLevel: metadata.Difficulty,
		CompanyName:     metadata.CompanyName,
	}
}
