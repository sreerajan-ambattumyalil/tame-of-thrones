package repository

import (
	"github.com/sreerajan-ambattumyalil/tame-of-thrones/pkg/model"
)

type IKingdomRepository interface {
	GetKingdoms() *model.Kingdoms
	SaveAllegiance(model.Kingdom, model.Kingdom) *model.Kingdoms
}

type KingdomRepository struct {
	Kingdoms *model.Kingdoms
}

var kingdomEmblems = map[model.KingdomName]model.Emblem{
	model.Land:  model.Panda,
	model.Water: model.Octopus,
	model.Ice:   model.Mammoth,
	model.Air:   model.Owl,
	model.Fire:  model.Dragon,
}

func getKingdoms() []*model.Kingdom {
	k := make([]*model.Kingdom, 0)
	for key, val := range kingdomEmblems {
		k = append(k, &model.Kingdom{Name: key, Emblem: val})
	}
	return k
}

func NewKingdomRepository() IKingdomRepository {
	kingdoms := model.Kingdoms{Kingdoms: getKingdoms()}
	return KingdomRepository{Kingdoms: &kingdoms}
}

func (k KingdomRepository) GetKingdoms() *model.Kingdoms {
	return k.Kingdoms
}

func (k KingdomRepository) SaveAllegiance(from model.Kingdom, to model.Kingdom) *model.Kingdoms {
	for _, k := range k.Kingdoms.Kingdoms {
		if k.Name == from.Name {
			k.Allegiance = to.Name
		}
	}
	return k.Kingdoms
}
