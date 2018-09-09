package runequest

import "fmt"

// Character represents a generic RPG character
type Character struct {
	Name           string
	Setting        string
	Description    string
	Race           *Race
	Culture        *Culture
	Statistics     map[string]*Statistic
	StatMap        []string
	DerivedStats   map[string]*DerivedStat
	DerivedMap     []string
	Skills         map[string]*Skill
	SkillMap       []string
	Advantages     map[string]*Advantage
	AdvantageMap   []string
	Spells         map[string]*Spell
	Powers         map[string]*Power
	HitLocations   map[string]*HitLocation
	HitLocationMap []string
	Gear           string
	PointCost      int
	InPlay         bool
	Updates        []*Update
}

// Update tracks live changes to Character
type Update struct {
	Date       string
	ChangeFrom string
	ChangeTo   string
	Cost       int
}

func (c Character) String() string {
	text := c.Name

	for _, stat := range c.Statistics {
		text += fmt.Sprintf("\n%s", stat)
	}

	return text
}
