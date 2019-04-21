package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnRulerQuestion(t *testing.T) {
	r := NewQuestionsRepository()
	result := r.GetQuestion("who is the ruler of southeros?")
	assert.Equal(t, Ruler, result)
}

func TestShouldReturnAlliesQuestion(t *testing.T) {
	r := NewQuestionsRepository()
	result := r.GetQuestion("allies of ruler?")
	assert.Equal(t, Allies, result)
}

func TestShouldReturnUnknownQuestion(t *testing.T) {
	r := NewQuestionsRepository()
	result := r.GetQuestion("allies of ruler")
	assert.Equal(t, Unknown, result)
}
