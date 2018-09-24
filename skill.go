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
// TODO Combine ModifySkill and SkillSpecialization
func (c *Character) ModifySkill(s Skill) {

	var response, skillName string

	if s.UserChoice {

		// Show slice of existing skills with identical CoreString

		q := fmt.Sprintf("Enter a specialization for %s or hit Enter to use (%s): ",
			s.CoreString, s.UserString)

		response = UserQuery(q)

		if response == "" {
			response = s.UserString
		}

		skillName = fmt.Sprintf("%s (%s)", s.CoreString, response)
	} else {
		skillName = s.Name
	}

	if c.Skills[skillName] == nil {
		// Create new Skill in map
		c.Skills[skillName] = &Skill{
			Name:            skillName,
			Category:        s.Category,
			CoreString:      s.CoreString,
			UserChoice:      false,
			UserString:      response,
			Base:            s.Base,
			CategoryValue:   s.CategoryValue,
			Value:           s.Value,
			Total:           s.Total,
			Min:             s.Min,
			Max:             s.Max,
			ExperienceCheck: s.ExperienceCheck,
		}
		// Remove base skill from map
		delete(c.Skills, s.CoreString)
	} else {
		// Modify existing skill
		if c.Skills[skillName].Base < s.Base {
			// Change Skill.Base if advantageous
			c.Skills[skillName].Base = s.Base
		}
		// Add or subtract s.Value from skill
		c.Skills[skillName].Value += s.Value
	}
}
