package runequest

// NewCharacter generates a random starting character in Runequest
func NewCharacter(name string) *Character {
	c := Character{
		Name:            name,
		Statistics:      RuneQuestStats,
		Abilities:       Abilities,
		SkillCategories: SkillCategories,
	}

	// Skills is a map of regular skills in Runequest
	c.Skills = Skills

	c.Skills["Dodge"].Base = c.Statistics["DEX"].Value * 2
	c.Skills["Jump"].Base = c.Statistics["DEX"].Value * 3

	return &c
}
