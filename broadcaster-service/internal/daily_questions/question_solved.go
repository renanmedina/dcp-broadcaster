package daily_questions

const QUESTION_SOLVED_EVENT_NAME = "QuestionSolved"

type QuestionSolved struct {
	Question         Question
	QuestionSolution QuestionSolution
}

func newQuestionSolved(question Question, solution QuestionSolution) QuestionSolved {
	return QuestionSolved{question, solution}
}

func (evt QuestionSolved) Name() string {
	return QUESTION_SOLVED_EVENT_NAME
}

func (evt QuestionSolved) ObjectId() string {
	return evt.Question.Id.String()
}

func (evt QuestionSolved) ObjectType() string {
	return "DailyQuestion"
}

func (evt QuestionSolved) EventData() map[string]interface{} {
	return evt.Question.ToDbMap()
}
