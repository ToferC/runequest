package runequest

// Homeland represents a homeland and cultural learnings
type Homeland struct {
	Name        string
	Description string
	Skills      []Skill
	// Base Skill List
	SkillChoices []SkillChoice
	// Options for skills
	Abilities      []Ability
	AbilityChoices []AbilityChoice
	AbilityList    []Ability
	PassionList    []Ability
	Passions       []Ability
}

// ChooseHomeland modifies a character's skills by homeland
func (c *Character) ChooseHomeland(hl *Homeland) {

	if c.Homeland == nil {
		// First homeland so apply all modifiers
		c.Homeland = hl
		c.ApplyHomeland()
	} else {
		c.RemoveHomeland()
		c.Homeland = hl
		c.ApplyHomeland()
		// Already has a homeland & need to remove previous homeland skills
	}
}

// ApplyHomeland applies a Homeland Template to a character
func (c *Character) ApplyHomeland() {

	for _, s := range c.Homeland.Skills {
		s.generateName()
		c.ModifySkill(s)
	}

	for _, choice := range c.Homeland.SkillChoices {
		// Find number of skills
		l := len(choice.Skills)

		// Choose random index
		r := ChooseRandom(l)

		// Select index from choice.Skills
		selected := choice.Skills[r]
		c.Homeland.Skills = append(c.Homeland.Skills, selected)

		// Modify or add skill
		c.ModifySkill(selected)
	}

	passions := c.Homeland.PassionList

	// Homelands grant 3 base passions
	// Find number of abilities

	for _, selected := range passions {
		c.Homeland.Passions = append(c.Homeland.Passions, selected)
		c.ModifyAbility(selected)
	}

	// Homeland grants a bonus to a rune affinity

	for _, ability := range c.Homeland.AbilityList {

		c.Homeland.Abilities = append(c.Homeland.Abilities, ability)

		// Modify or add ability
		c.ModifyAbility(ability)
	}
}

// RemoveHomeland removes all Homeland Modifers from a character
func (c *Character) RemoveHomeland() {

	for _, s := range c.Homeland.Skills {
		s.HomelandValue = 0

		if s.Base > 0 {
			s.Base += s.Base * -1
		}
		c.ModifySkill(s)
	}

	for _, p := range c.Homeland.Passions {
		p.HomelandValue = 0

		if p.Base > 0 {
			p.Base += p.Base * -1
		}
		c.ModifyAbility(p)
	}

	for _, a := range c.Homeland.Abilities {
		a.HomelandValue = 0

		if a.Base > 0 {
			a.Base += a.Base * -1
		}
		c.ModifyAbility(a)
	}
}
