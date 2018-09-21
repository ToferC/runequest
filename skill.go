package runequest

import "fmt"

// Skill is a learned ability of an RPG Character
type Skill struct {
	Name            string
	Category        string
	UserChoice      bool
	CoreString      string
	UserString      string
	Base            int
	CategoryValue   int
	Value           int
	Total           int
	Min             int
	Max             int
	ExperienceCheck bool
}

// SkillChoice is a choice between 2 or more skills
type SkillChoice struct {
	Skills []Skill
}

func (s *Skill) String() string {

	s.Total = s.Base + s.Value + s.CategoryValue

	text := ""

	text += fmt.Sprintf("%s %d%%", s.Name, s.Total)

	return text
}

// ModifySkill adds or modifies a Skill value
func (c *Character) ModifySkill(s Skill) {
	if c.Skills[s.Name] == nil {
		// New Skill
		c.Skills[s.Name] = &Skill{
			Name:     s.Name,
			Base:     s.Base,
			Value:    s.Value,
			Category: s.Category,
		}
	} else {
		// Modify existing skill
		if c.Skills[s.Name].Base != s.Base {
			// Change Skill.Base if needed
			c.Skills[s.Name].Base = s.Base
		}
		// Add or subtract s.Value from skill
		c.Skills[s.Name].Value += s.Value
	}
}
