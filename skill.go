package runequest

import "fmt"

// Skill is a learned ability of an RPG Character
type Skill struct {
	Name               string
	Category           string
	UserChoice         bool
	CoreString         string
	UserString         string
	Base               int
	CategoryValue      int
	HomelandValue      int
	OccupationValue    int
	CultValue          int
	CreationBonusValue int
	Updates            []*Update
	Value              int
	InPlayXPValue      int
	Total              int
	Min                int
	Max                int
	ExperienceCheck    bool
}

// SkillChoice is a choice between 2 or more skills
type SkillChoice struct {
	Skills []Skill
}

// UpdateSkill totals skill values based on input
func (s *Skill) UpdateSkill() {

	s.generateName()

	updates := 0

	for _, u := range s.Updates {
		updates += u.Value
	}

	s.Total = s.Base + s.CategoryValue + s.HomelandValue + s.OccupationValue + s.CultValue + s.CreationBonusValue + s.InPlayXPValue + s.Value + updates
}

// generateName sets the skill map name
func (s *Skill) generateName() {

	var n string

	if s.UserString != "" {
		n = fmt.Sprintf("%s (%s)", s.CoreString, s.UserString)
	} else {
		n = s.CoreString
	}
	s.Name = n
}

func (s *Skill) String() string {

	s.UpdateSkill()

	text := ""

	text += fmt.Sprintf("%s %d%%", s.Name, s.Total)

	return text
}

// ModifySkill adds or modifies a Skill value
func (c *Character) ModifySkill(s Skill) {

	/*
		var response string

		if s.UserChoice {

			// Show slice of existing skills with identical CoreString

			q := fmt.Sprintf("Enter a specialization for %s or hit Enter to use (%s): ",
				s.CoreString, s.UserString)

			response = UserQuery(q)

			if response == "" {
				response = s.UserString
			}
			s.UserString = response
		}
	*/

	s.generateName()

	if c.Skills[s.Name] == nil {
		// Create new Skill in map
		c.Skills[s.Name] = &Skill{
			Name:               s.Name,
			Category:           s.Category,
			CoreString:         s.CoreString,
			UserChoice:         false,
			UserString:         s.UserString,
			Base:               s.Base,
			CategoryValue:      s.CategoryValue,
			HomelandValue:      s.HomelandValue,
			OccupationValue:    s.OccupationValue,
			CultValue:          s.CultValue,
			CreationBonusValue: s.CreationBonusValue,
			InPlayXPValue:      s.InPlayXPValue,
			Value:              s.Value,
			Total:              s.Total,
			Min:                s.Min,
			Max:                s.Max,
			ExperienceCheck:    s.ExperienceCheck,
		}
		// Remove base skill from map
		delete(c.Skills, s.CoreString)
	} else {
		// Modify existing skill
		sk := c.Skills[s.Name]
		if sk.Base < s.Base {
			// Change Skill.Base if advantageous
			sk.Base = s.Base
		}
		// Add or subtract s.XXValue from skill
		// This doesn't work
		sk.Value = s.Value
		sk.HomelandValue = s.HomelandValue
		sk.OccupationValue = s.OccupationValue
		sk.CultValue = s.CultValue
		sk.CreationBonusValue = s.CreationBonusValue
		sk.InPlayXPValue = s.InPlayXPValue
	}
}

// ApplySkillChoice executes the skill choice on a character
func (c *Character) ApplySkillChoice(sc SkillChoice, r int) {

	// Select index from choice.Skills
	selected := sc.Skills[r]

	// Modify or add skill
	c.ModifySkill(selected)
}

// MapSkills determines the best way to update a Character's SkillMap
func (c *Character) MapSkills(s *Skill, targetString string) {

	if c.Skills[targetString] != nil {
		// Skill exists in Character, modify it via pointer
		c.Skills[targetString].HomelandValue = s.HomelandValue

		if s.Base > c.Skills[targetString].Base {
			c.Skills[targetString].Base = s.Base
		}

		fmt.Println(c.Skills[targetString])

	} else {
		// We need to find the base skill from the Master list or create it
		if Skills[s.CoreString] == nil {
			fmt.Println("Skill is new: " + targetString)

			// New Skill
			baseSkill := s
			// Update our new skill
			sc := c.SkillCategories[baseSkill.Category]

			baseSkill.CategoryValue = sc.Value

			baseSkill.UpdateSkill()

			fmt.Println("Add Skill to character: " + baseSkill.Name)
			c.Skills[baseSkill.Name] = baseSkill
		} else {
			// Skill exists in master list
			fmt.Println("Skill in master list: " + targetString)

			baseSkill := Skills[s.CoreString]
			fmt.Println(baseSkill)
			baseSkill.HomelandValue = s.HomelandValue

			if s.Base > baseSkill.Base {
				baseSkill.Base = s.Base
			}
			if s.UserString != "" {
				baseSkill.UserString = s.UserString
			}
			// Add Skill to Character
			fmt.Println("Add Skill to character: " + baseSkill.Name)
			c.Skills[baseSkill.Name] = baseSkill
			c.Skills[baseSkill.Name].UpdateSkill()
		}
	}
}
