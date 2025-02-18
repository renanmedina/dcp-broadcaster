package daily_questions

import (
	"fmt"

	"github.com/renanmedina/dcp-broadcaster/utils"
	"gorm.io/gorm"
)

const (
	UNIQUE_CONSTRAINT_ERROR = "unique_violation"
)

type QuestionsRepository struct {
	db     *gorm.DB
	logger *utils.ApplicationLogger
}

type QuestionNotFound struct {
	msg string
}

func (er QuestionNotFound) Error() string {
	return er.msg
}

func NewQuestionNotFound(msg string) QuestionNotFound {
	return QuestionNotFound{msg}
}

func (r *QuestionsRepository) GetLatest() *Question {
	var question Question
	result := r.db.WithContext(r.logger.GetCurrentContext()).Limit(1).Order("received_at desc").Find(&question)

	if result.Error != nil {
		return nil
	}

	return &question
}

func (r *QuestionsRepository) GetAll() ([]Question, error) {
	var questions []Question
	result := r.db.Model(&Question{}).Preload("Solutions").WithContext(r.logger.GetCurrentContext()).Find(&questions)

	if result.Error != nil {
		return make([]Question, 0), result.Error
	}

	return questions, nil
}

func (r *QuestionsRepository) GetByOriginalId(id string) (*Question, error) {
	var question Question
	result := r.db.Model(&Question{}).Preload("Solutions").WithContext(r.logger.GetCurrentContext()).First(&question, "original_id = ?", id)

	if result.Error != nil {
		return nil, NewQuestionNotFound(fmt.Sprintf("Question with original id %s not found", id))
	}

	return &question, nil
}

func (r *QuestionsRepository) GetById(id string) (*Question, error) {
	var question Question
	result := r.db.Model(&Question{}).Preload("Solutions").WithContext(r.logger.GetCurrentContext()).First(&question, "id = ?", id)

	if result.Error != nil {
		return nil, NewQuestionNotFound(fmt.Sprintf("Question %s not found", id))
	}

	return &question, nil
}

func (r *QuestionsRepository) Save(question Question) (*Question, error) {
	result := r.db.WithContext(r.logger.GetCurrentContext()).Save(&question)

	if result.Error != nil {
		return nil, result.Error
	}

	return &question, nil
}

func NewQuestionsRepository() QuestionsRepository {
	return QuestionsRepository{
		db:     utils.GetDatabaseConnection(),
		logger: utils.GetApplicationLogger(),
	}
}
