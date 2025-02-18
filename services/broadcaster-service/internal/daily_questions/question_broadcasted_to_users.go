package daily_questions

const QUESTION_BROADCASTED_EVENT_NAME = "QuestionBroadcastedToUsers"

type QuestionBroadcastedToUsers struct {
	Question Question
}

func newQuestionBroadcastedToUsers(question Question) QuestionBroadcastedToUsers {
	return QuestionBroadcastedToUsers{question}
}

func (evt QuestionBroadcastedToUsers) Name() string {
	return QUESTION_BROADCASTED_EVENT_NAME
}

func (evt QuestionBroadcastedToUsers) ObjectId() string {
	return evt.Question.Id
}

func (evt QuestionBroadcastedToUsers) ObjectType() string {
	return "DailyQuestion"
}

func (evt QuestionBroadcastedToUsers) EventData() map[string]interface{} {
	return map[string]interface{}{
		"question": evt.Question.ToDbMap(),
	}
}
