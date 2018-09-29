package runequest

// Occupation represents a profession in Runequest
type Occupation struct {
	Name             string
	Description      string
	Skills           []Skill
	SkillChoices     []SkillChoice
	StandardOfLiving string
	Income           int
	CultChoices      []Cult
	PassionList      []Ability
	Passions         []Ability
	Abilities        []Ability
	AbilityChoices   []AbilityChoice
	Ransom           int
	Equipment        []string
}

// ChooseOccupation modifies a character's skills by Occupation
func (c *Character) ChooseOccupation(o *Occupation) {

	if c.Occupation == nil {
		// First Occupation so apply all modifiers
		c.Occupation = o
		c.ApplyOccupation()
	} else {
		c.RemoveOccupation()
		c.Occupation = o
		c.ApplyOccupation()
		// Already has a Occupation & need to remove previous Occupation skills
	}
}

// ApplyOccupation applies a Occupation Template to a character
func (c *Character) ApplyOccupation() {

	for _, s := range c.Occupation.Skills {
		c.ModifySkill(s)
	}

	for _, choice := range c.Occupation.SkillChoices {
		// Find number of skills
		l := len(choice.Skills)

		// Choose random index
		r := ChooseRandom(l)

		// Select index from choice.Skills
		selected := choice.Skills[r]
		c.Occupation.Skills = append(c.Occupation.Skills, selected)

		// Modify or add skill
		c.ModifySkill(selected)
	}

	passions := c.Occupation.PassionList
	// Find number of abilities
	l := len(passions)

	// Choose random index
	r := ChooseRandom(l)

	// Select index from Passions
	selected := passions[r]
	c.Occupation.Passions = append(c.Occupation.Passions, selected)

	// Modify or add ability
	c.ModifyAbility(selected)

	// Same for abilities

	for _, choice := range c.Occupation.AbilityChoices {
		// Find number of skills
		l = len(choice.Abilities)

		// Choose random index
		r := ChooseRandom(l)

		// Select index from choice.Abilities
		selected := choice.Abilities[r]
		c.Occupation.Abilities = append(c.Occupation.Abilities, selected)

		// Modify or add skill
		c.ModifyAbility(selected)
	}
}

// RemoveOccupation removes all Occupation Modifers from a character
func (c *Character) RemoveOccupation() {

	for _, s := range c.Occupation.Skills {
		s.Value *= -1
		c.ModifySkill(s)
	}

	for _, p := range c.Occupation.Passions {
		p.Value *= -1
		c.ModifyAbility(p)
	}

	for _, a := range c.Occupation.Abilities {
		a.Value *= -1
		c.ModifyAbility(a)
	}
}
