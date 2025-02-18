package daily_questions

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Id              string `gorm:"primaryKey"`
	OriginalId      string
	DifficultyLevel string
	Title           string
	EmailBody       string `gorm:"column:question_email_body"`
	Text            string `gorm:"column:question_text"`
	CompanyName     string
	ReceivedAt      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Solutions       []QuestionSolution `gorm:"foreignKey:DailyQuestionId"`
}

// gorm before create hook
func (q *Question) BeforeCreate(tx *gorm.DB) (err error) {
	q.Id = uuid.New().String()
	return nil
}

func (Question) TableName() string {
	return "daily_questions"
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
		OriginalId:      metadata.MessageId,
		Title:           metadata.Title,
		EmailBody:       metadata.BodyHtml,
		Text:            metadata.BodyText,
		ReceivedAt:      metadata.Date,
		DifficultyLevel: metadata.Difficulty,
		CompanyName:     metadata.CompanyName,
	}
}
