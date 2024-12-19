package daily_questions

const QUESTION_SOLUTION_CREATED_EVENT_NAME = "QuestionSolutionCreated"

type QuestionSolutionCreated struct {
	QuestionSolution QuestionSolution
}

func newQuestionSolutionCreated(solution QuestionSolution) QuestionSolutionCreated {
	return QuestionSolutionCreated{solution}
}

func (evt QuestionSolutionCreated) Name() string {
	return QUESTION_SOLUTION_CREATED_EVENT_NAME
}

func (evt QuestionSolutionCreated) ObjectId() string {
	return evt.QuestionSolution.Id.String()
}

func (evt QuestionSolutionCreated) ObjectType() string {
	return "DailyQuestionSolution"
}

func (evt QuestionSolutionCreated) EventData() map[string]interface{} {
	return evt.QuestionSolution.ToDbMap()
}
