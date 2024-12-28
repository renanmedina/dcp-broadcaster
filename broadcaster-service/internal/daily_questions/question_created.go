package daily_questions

const QUESTION_CREATED_EVENT_NAME = "QuestionCreated"

type QuestionCreated struct {
	Question Question
}

func NewQuestionCreatedEvent(question Question) QuestionCreated {
	return QuestionCreated{question}
}

func (evt QuestionCreated) Name() string {
	return QUESTION_CREATED_EVENT_NAME
}

func (evt QuestionCreated) ObjectId() string {
	return evt.Question.Id
}

func (evt QuestionCreated) ObjectType() string {
	return "DailyQuestion"
}

func (evt QuestionCreated) EventData() map[string]interface{} {
	return evt.Question.ToDbMap()
}
