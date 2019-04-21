package parser

import (
	"testing"

	"github.com/sreerajan-ambattumyalil/tame-of-thrones/pkg/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type QuestionsRepositoryMock struct {
	mock.Mock
}

func (m *QuestionsRepositoryMock) GetQuestion(text string) repository.Question {

	args := m.Called(text)
	return args.Get(0).(repository.Question)

}
func TestShouldRespondToQuestionsChannelForRuler(t *testing.T) {

	repoMock := new(QuestionsRepositoryMock)
	responseChannel := make(chan repository.Question)

	repoMock.On("GetQuestion", mock.Anything).Return(repository.Ruler)

	parser := NewInputParser(repoMock, responseChannel)

	go parser.Parse("Who is the ruler of Southeros?")

	resp := <-responseChannel

	assert.Equal(t, repository.Ruler, resp)
	repoMock.AssertNumberOfCalls(t, "GetQuestion", 1)
}

func TestShouldRespondToQuestionsChannelForAllies(t *testing.T) {

	repoMock := new(QuestionsRepositoryMock)
	responseChannel := make(chan repository.Question)

	repoMock.On("GetQuestion", mock.Anything).Return(repository.Allies)

	parser := NewInputParser(repoMock, responseChannel)

	go parser.Parse("Allies of ?")

	resp := <-responseChannel

	assert.Equal(t, repository.Allies, resp)
	repoMock.AssertNumberOfCalls(t, "GetQuestion", 1)
}
