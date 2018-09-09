package runequest

import "fmt"

// Statistic is a core element for all characters in an RPG system
type Statistic struct {
	Name  string
	Value int
}

func (s *Statistic) String() string {
	text := fmt.Sprintf("%s: %d", s.Name, s.Value)
	return text
}

// RuneQuestStats is the base stats for RuneQuest
var RuneQuestStats = map[string]*Statistic{
	"STR": &Statistic{
		Name:  "Strength",
		Value: 10,
	},
	"DEX": &Statistic{
		Name:  "Dexterity",
		Value: 10,
	},
	"INT": &Statistic{
		Name:  "Intelligence",
		Value: 12,
	},
	"POW": &Statistic{
		Name:  "Power",
		Value: 10,
	},
	"SIZ": &Statistic{
		Name:  "Size",
		Value: 12,
	},
	"CHA": &Statistic{
		Name:  "Charisma",
		Value: 10,
	},
}
