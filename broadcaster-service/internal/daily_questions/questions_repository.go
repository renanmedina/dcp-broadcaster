package daily_questions

import (
	"github.com/renanmedina/dcp-broadcaster/utils"
	"gorm.io/gorm"
)

const (
	QUESTIONS_TABLE_NAME    = "daily_questions"
	UNIQUE_CONSTRAINT_ERROR = "unique_violation"
)

type QuestionsRepository struct {
	db *gorm.DB
}

func (r *QuestionsRepository) GetLatest() *Question {
	var question Question
	result := r.db.Table(QUESTIONS_TABLE_NAME).Limit(1).Order("received_at desc").Find(&question)

	if result.Error != nil {
		return nil
	}

	return &question
}

func (r *QuestionsRepository) GetByOriginalId(id string) (*Question, error) {
	var question Question
	result := r.db.Table(QUESTIONS_TABLE_NAME).First(&question, "original_id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &question, nil
}

func (r *QuestionsRepository) GetById(id string) (*Question, error) {
	var question Question
	result := r.db.Table(QUESTIONS_TABLE_NAME).First(&question, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &question, nil
}

func (r *QuestionsRepository) Save(question Question) (*Question, error) {
	result := r.db.Table(QUESTIONS_TABLE_NAME).Save(&question)

	if result.Error != nil {
		return nil, result.Error
	}

	return &question, nil
}

func NewQuestionsRepository() QuestionsRepository {
	return QuestionsRepository{
		db: utils.GetDatabaseConnection(),
	}
}
