package daily_questions

import (
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/renanmedina/dcp-broadcaster/utils"
)

const (
	QUESTIONS_TABLE_NAME = "daily_questions"
)

type QuestionsRepository struct {
	db *utils.DatabaseAdapdater
}

func (r *QuestionsRepository) GetByOriginalId(id string) (*Question, error) {
	scanner := r.db.SelectOne(QUESTIONS_TABLE_NAME, map[string]interface{}{
		"original_id": id,
	})

	question, err := buildQuestionFromDb(*scanner)

	if err != nil {
		return nil, err
	}

	return &question, err
}

func (r *QuestionsRepository) Save(question Question) (*Question, error) {
	if !question.Persisted {
		_, err := r.db.Insert(QUESTIONS_TABLE_NAME, question.ToDbMap())

		if err != nil {
			return nil, err
		}

	} else {
		_, err := r.db.UpdateById(QUESTIONS_TABLE_NAME, question.Id.String(), question.ToDbMap())

		if err != nil {
			return nil, err
		}
	}

	question.Persisted = true
	return &question, nil
}

func NewQuestionsRepository() *QuestionsRepository {
	return &QuestionsRepository{
		db: utils.GetDatabase(),
	}
}

func buildQuestionFromDb(dbRow squirrel.RowScanner) (Question, error) {
	var question Question
	dbRow.Scan(
		&question.Id,
		&question.OriginalId,
		&question.DifficultyLevel,
		&question.ReceivedAt,
		&question.Title,
		&question.EmailBody,
		&question.Text,
		&question.CompanyName,
		&question.CreatedAt,
		&question.UpdatedAt,
	)

	if question.Id.String() == "" {
		return Question{}, errors.New("can't find User")
	}

	question.Persisted = true
	return question, nil
}
