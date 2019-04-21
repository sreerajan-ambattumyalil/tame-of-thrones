package parser

import (
	"strings"

	"github.com/sreerajan-ambattumyalil/tame-of-thrones/pkg/repository"
)

type Parser interface {
	Parse()
}

type InputParser struct {
	questionsRepo    repository.IQuestionsRepository
	questionsChannel chan repository.Question
}

func NewInputParser(questionsRepo repository.IQuestionsRepository, questionsChannel chan repository.Question) InputParser {
	return InputParser{questionsRepo: questionsRepo, questionsChannel: questionsChannel}
}

func (p InputParser) Parse(input string) {
	input = strings.ToLower(input)
	question := p.questionsRepo.GetQuestion(input)
	if question != repository.Unknown {
		p.questionsChannel <- question
	}
	p.questionsChannel <- repository.Unknown
}
