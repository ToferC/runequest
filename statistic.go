package runequest

import "fmt"

// Statistic is a core element for all characters in an RPG system
type Statistic struct {
	Name            string
	Value           int
	Base            int
	Total           int
	Max             int
	Min             int
	ExperienceCheck bool
}

func (s *Statistic) String() string {

	s.Total = s.Base + s.Value

	text := fmt.Sprintf("%s: %d", s.Name, s.Total)
	return text
}

// TotalStatistics updates values for stats after being modified
func (c *Character) TotalStatistics() {

	for _, s := range c.Statistics {
		s.Total = s.Base + s.Value
	}
}

// RuneQuestStats is the base stats for RuneQuest
var RuneQuestStats = map[string]*Statistic{
	"STR": &Statistic{
		Name: "Strength",
		Base: RollDice(6, 1, 0, 3),
	},
	"DEX": &Statistic{
		Name: "Dexterity",
		Base: RollDice(6, 1, 0, 3),
	},
	"INT": &Statistic{
		Name: "Intelligence",
		Base: RollDice(6, 1, 6, 2),
	},
	"CON": &Statistic{
		Name: "Constitution",
		Base: RollDice(6, 1, 6, 2),
	},
	"POW": &Statistic{
		Name: "Power",
		Base: RollDice(6, 1, 0, 3),
	},
	"SIZ": &Statistic{
		Name: "Size",
		Base: RollDice(6, 1, 6, 2),
	},
	"CHA": &Statistic{
		Name: "Charisma",
		Base: RollDice(6, 1, 0, 3),
	},
}
