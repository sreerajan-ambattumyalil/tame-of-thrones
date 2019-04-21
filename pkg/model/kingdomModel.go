package model

type Kingdoms struct {
	Kingdoms []*Kingdom
}

type Kingdom struct {
	Name       KingdomName
	Emblem     Emblem
	Allegiance KingdomName
}

type KingdomName string

const (
	None  KingdomName = "None"
	Shan  KingdomName = "shan"
	Land  KingdomName = "land"
	Water KingdomName = "water"
	Ice   KingdomName = "ice"
	Air   KingdomName = "air"
	Fire  KingdomName = "fire"
)

type Emblem string

const (
	Panda   Emblem = "panda"
	Octopus Emblem = "octopus"
	Mammoth Emblem = "mammoth"
	Owl     Emblem = "owl"
	Dragon  Emblem = "dragon"
)

func (e Emblem) rule() map[rune]int {
	rules := make(map[rune]int)
	for _, char := range e {
		if val, ok := rules[char]; ok {
			rules[char] = val + 1
		} else {
			rules[char] = 1
		}
	}
	return rules
}

func (k Kingdoms) Ruler() string {
	stats := make(map[KingdomName]int)
	for _, kingdom := range k.Kingdoms {
		if _, ok := stats[kingdom.Allegiance]; ok && kingdom.Allegiance != "" {
			stats[kingdom.Allegiance] = stats[kingdom.Allegiance] + 1
			if stats[kingdom.Allegiance] > ((len(k.Kingdoms) - 1) / 2) {
				return string(kingdom.Allegiance)
			}
		} else {
			stats[kingdom.Allegiance] = 1
		}
	}
	return "None"
}

func (k Kingdoms) Allies() []KingdomName {
	ruler := k.Ruler()
	allies := make([]KingdomName, 0)
	for _, kingdom := range k.Kingdoms {
		if string(kingdom.Allegiance) == ruler {
			allies = append(allies, kingdom.Name)
		}
	}
	return allies
}
