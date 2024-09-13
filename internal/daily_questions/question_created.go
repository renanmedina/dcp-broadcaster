package daily_questions

const QUESTION_CREATED_EVENT_NAME = "QuestionCreated"

type QuestionCreated struct {
	question Question
}

func newQuestionCreated(question Question) QuestionCreated {
	return QuestionCreated{question}
}

func (evt QuestionCreated) Name() string {
	return QUESTION_CREATED_EVENT_NAME
}

func (evt QuestionCreated) ObjectId() string {
	return evt.question.Id.String()
}

func (evt QuestionCreated) ObjectType() string {
	return "DailyQuestion"
}

func (evt QuestionCreated) EventData() map[string]interface{} {
	return evt.question.ToDbMap()
}
