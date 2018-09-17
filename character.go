package runequest

import "fmt"

// Character represents a generic RPG character
type Character struct {
	Name            string
	Setting         string
	Description     string
	Race            *Race
	Culture         *Culture
	Cult            *Cult
	Abilities       map[string]*Ability
	Statistics      map[string]*Statistic
	StatMap         []string
	DerivedStats    map[string]*DerivedStat
	DerivedMap      []string
	Skills          map[string]*Skill
	SkillMap        []string
	SkillCategories map[string]*SkillCategory
	Advantages      map[string]*Advantage
	AdvantageMap    []string
	Spells          map[string]*Spell
	Powers          map[string]*Power
	HitLocations    map[string]*HitLocation
	HitLocationMap  []string
	Gear            string
	PointCost       int
	InPlay          bool
	Updates         []*Update
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

	text += "\n\nStats:\n"
	for _, stat := range c.Statistics {
		text += fmt.Sprintf("%s\n", stat)
	}

	text += "\n\nAbilities:"

	for _, at := range AbilityTypes {

		text += fmt.Sprintf("\n\n**%s**", at)

		for _, ability := range c.Abilities {

			if ability.Type == at {
				text += fmt.Sprintf("\n%s", ability)
			}
		}
	}

	text += "\n\nSkills:"

	for _, co := range CategoryOrder {

		sc := c.SkillCategories[co]

		text += fmt.Sprintf("%s", sc)
		for _, skill := range c.Skills {

			if skill.category == sc.Name {
				text += fmt.Sprintf("\n%s", skill)
			}
		}

	}
	return text
}
