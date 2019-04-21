package repository

import (
	"testing"

	"github.com/sreerajan-ambattumyalil/tame-of-thrones/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAllKingdoms(t *testing.T) {
	repo := NewKingdomRepository()
	res := repo.GetKingdoms().Kingdoms

	assert.Equal(t, 5, len(res))
}

func TestShouldSaveAllegiance(t *testing.T) {
	repo := NewKingdomRepository()
	from := model.Kingdom{Name: model.Land}
	to := model.Kingdom{Name: model.Ice}
	res := repo.SaveAllegiance(from, to)

	for _, k := range res.Kingdoms {
		if k.Name == model.Land {
			assert.Equal(t, model.Ice, k.Allegiance)
		}
	}
}
