package repository

type Question int

const (
	Ruler   Question = 0
	Allies  Question = 1
	Unknown Question = 2
)

type IQuestionsRepository interface {
	GetQuestion(text string) Question
}

type QuestionsRepository struct {
}

func NewQuestionsRepository() IQuestionsRepository {
	return QuestionsRepository{}
}

func getMap() map[string]Question {

	questionMap := make(map[string]Question)
	questionMap["who is the ruler of southeros?"] = Ruler
	questionMap["allies of ruler?"] = Allies

	return questionMap
}

func (r QuestionsRepository) GetQuestion(text string) Question {
	questions := getMap()
	val, exists := questions[text]
	if exists {
		return val
	}
	return Unknown
}
