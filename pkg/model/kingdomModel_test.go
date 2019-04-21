package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnRuleForPanda(t *testing.T) {
	rules := Panda.rule()

	expected := make(map[rune]int)
	expected['p'] = 1
	expected['a'] = 2
	expected['n'] = 1
	expected['d'] = 1
	assert.Equal(t, expected, rules)
}

func TestShouldReturnRuleForOctopus(t *testing.T) {
	rules := Octopus.rule()

	expected := make(map[rune]int)
	expected['o'] = 2
	expected['c'] = 1
	expected['t'] = 1
	expected['p'] = 1
	expected['u'] = 1
	expected['s'] = 1
	assert.Equal(t, expected, rules)
}

func TestShouldReturnRuleForMammoth(t *testing.T) {
	rules := Mammoth.rule()

	expected := make(map[rune]int)
	expected['m'] = 3
	expected['a'] = 1
	expected['o'] = 1
	expected['t'] = 1
	expected['h'] = 1
	assert.Equal(t, expected, rules)
}

func TestShouldReturnRuleForOwl(t *testing.T) {
	rules := Owl.rule()

	expected := make(map[rune]int)
	expected['o'] = 1
	expected['w'] = 1
	expected['l'] = 1
	assert.Equal(t, expected, rules)
}

func TestShouldReturnRuleForDragon(t *testing.T) {
	rules := Dragon.rule()

	expected := make(map[rune]int)
	expected['d'] = 1
	expected['r'] = 1
	expected['a'] = 1
	expected['g'] = 1
	expected['o'] = 1
	expected['n'] = 1
	assert.Equal(t, expected, rules)
}

func TestShouldReturnTheRulerAsShanIf3OutOf5Allegiance(t *testing.T) {
	kingdoms := make([]*Kingdom, 0)
	kingdoms = append(kingdoms, &Kingdom{Name: Land, Allegiance: Shan})
	kingdoms = append(kingdoms, &Kingdom{Name: Air, Allegiance: Shan})
	kingdoms = append(kingdoms, &Kingdom{Name: Ice, Allegiance: Shan})
	kingdoms = append(kingdoms, &Kingdom{Name: Water})
	kingdoms = append(kingdoms, &Kingdom{Name: Fire})

	res := Kingdoms{Kingdoms: kingdoms}.Ruler()

	assert.Equal(t, string(Shan), res)
}

func TestShouldReturnTheAlliesForShanIf3OutOf5Allegiance(t *testing.T) {
	kingdoms := make([]*Kingdom, 0)
	kingdoms = append(kingdoms, &Kingdom{Name: Land, Allegiance: Shan})
	kingdoms = append(kingdoms, &Kingdom{Name: Air, Allegiance: Shan})
	kingdoms = append(kingdoms, &Kingdom{Name: Ice, Allegiance: Shan})
	kingdoms = append(kingdoms, &Kingdom{Name: Water})
	kingdoms = append(kingdoms, &Kingdom{Name: Fire})

	res := Kingdoms{Kingdoms: kingdoms}.Allies()

	allies := make([]KingdomName, 0)
	allies = append(allies, Land)
	allies = append(allies, Air)
	allies = append(allies, Ice)
	assert.Equal(t, allies, res)
}

func TestShouldReturnTheRulerAsNoneIfAllegianceIsLessThan3(t *testing.T) {
	kingdoms := make([]*Kingdom, 0)
	kingdoms = append(kingdoms, &Kingdom{Name: Land, Allegiance: Shan})
	kingdoms = append(kingdoms, &Kingdom{Name: Air, Allegiance: Shan})
	kingdoms = append(kingdoms, &Kingdom{Name: Ice})
	kingdoms = append(kingdoms, &Kingdom{Name: Water})
	kingdoms = append(kingdoms, &Kingdom{Name: Fire})

	res := Kingdoms{Kingdoms: kingdoms}.Ruler()

	assert.Equal(t, string(None), res)
}
