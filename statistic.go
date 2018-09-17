package runequest

import "fmt"

// Statistic is a core element for all characters in an RPG system
type Statistic struct {
	Name            string
	Value           int
	Base            int
	Max             int
	Min             int
	ExperienceCheck bool
}

func (s *Statistic) String() string {
	text := fmt.Sprintf("%s: %d", s.Name, s.Value)
	return text
}

// RuneQuestStats is the base stats for RuneQuest
var RuneQuestStats = map[string]*Statistic{
	"STR": &Statistic{
		Name:  "Strength",
		Value: RollDice(6, 1, 0, 3),
	},
	"DEX": &Statistic{
		Name:  "Dexterity",
		Value: RollDice(6, 1, 0, 3),
	},
	"INT": &Statistic{
		Name:  "Intelligence",
		Value: RollDice(6, 1, 6, 2),
	},
	"CON": &Statistic{
		Name:  "Constitution",
		Value: RollDice(6, 1, 6, 2),
	},
	"POW": &Statistic{
		Name:  "Power",
		Value: RollDice(6, 1, 0, 3),
	},
	"SIZ": &Statistic{
		Name:  "Size",
		Value: RollDice(6, 1, 6, 2),
	},
	"CHA": &Statistic{
		Name:  "Charisma",
		Value: RollDice(6, 1, 0, 3),
	},
}
