package runequest

// Cult represents a Religion in Runequest
type Cult struct {
	Name        string
	Description string
	SkillList   []Skill
	SpellList   []*Spell
}

// ChooseCult modifies a character's skills by homeland
func (c *Character) ChooseCult() {

	// Homelands is a map of possible homelands in Runequest
	var Cults = map[string]Cult{
		"Orlanth": Cult{
			Name: "Orlanth",
			SkillList: []Skill{
				Skill{
					Name:     "Cult Lore (Orlanth)",
					Value:    15,
					Category: "Knowledge",
				},
				Skill{
					Name:     "Worship (Orlanth)",
					Value:    15,
					Category: "Magic",
				},
				Skill{
					Name:  "Meditate",
					Value: 25,
				},
				Skill{
					Name:  "Orate",
					Value: 30,
				},
				Skill{
					Name:  "Sing",
					Value: 10,
				},
				Skill{
					Name:     "Speak (Stormspeech)",
					Value:    10,
					Category: "Communication",
				},
			},
		},
		// Esrolia
	}

	c.Cult = Cults["Orlanth"]

	for _, s := range c.Cult.SkillList {
		c.ModifySkill(s)
	}
}
