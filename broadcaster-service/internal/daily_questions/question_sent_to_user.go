package daily_questions

import "github.com/renanmedina/dcp-broadcaster/internal/accounts"

const QUESTION_SENT_EVENT_NAME = "QuestionSentToUser"

type QuestionSentToUser struct {
	Question Question
	User     accounts.User
}

func newQuestionSentToUser(question Question, user accounts.User) QuestionSentToUser {
	return QuestionSentToUser{question, user}
}

func (evt QuestionSentToUser) Name() string {
	return QUESTION_SENT_EVENT_NAME
}

func (evt QuestionSentToUser) ObjectId() string {
	return evt.Question.Id.String()
}

func (evt QuestionSentToUser) ObjectType() string {
	return "DailyQuestion"
}

func (evt QuestionSentToUser) EventData() map[string]interface{} {
	return map[string]interface{}{
		"question": evt.Question.ToDbMap(),
		"user":     evt.User.ToDbMap(),
	}
}
