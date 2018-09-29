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
	// Find number of abilities
	l := len(passions)

	// Choose random index
	r := ChooseRandom(l)

	// Select index from Passions
	selected := passions[r]
	c.Homeland.Passions = append(c.Homeland.Passions, selected)

	// Modify or add ability
	c.ModifyAbility(selected)

	// Same for abilities

	for _, choice := range c.Homeland.AbilityChoices {
		// Find number of skills
		l = len(choice.Abilities)

		// Choose random index
		r := ChooseRandom(l)

		// Select index from choice.Abilities
		selected := choice.Abilities[r]
		c.Homeland.Abilities = append(c.Homeland.Abilities, selected)

		// Modify or add skill
		c.ModifyAbility(selected)
	}
}

// RemoveHomeland removes all Homeland Modifers from a character
func (c *Character) RemoveHomeland() {

	for _, s := range c.Homeland.Skills {
		s.Value *= -1
		c.ModifySkill(s)
	}

	for _, p := range c.Homeland.Passions {
		p.Value *= -1
		c.ModifyAbility(p)
	}

	for _, a := range c.Homeland.Abilities {
		a.Value *= -1
		c.ModifyAbility(a)
	}
}
