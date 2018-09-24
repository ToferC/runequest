package runequest

// Occupation represents a profession in Runequest
type Occupation struct {
	Name             string
	Description      string
	SkillList        []Skill
	SkillChoices     []SkillChoice
	StandardOfLiving string
	Income           int
	Cults            []Cult
	Passions         []string
	Ransom           int
	Equipment        []string
}

// ChooseOccupation modifies a character's skills by homeland
func (c *Character) ChooseOccupation() {

	// Homelands is a map of possible homelands in Runequest
	var Occupations = map[string]Occupation{
		"Farmer": Occupation{
			Name: "Farmer",
			SkillList: []Skill{
				Skill{
					Name:     "Homeland Lore (Local)",
					Value:    15,
					Category: "Knowledge",
				},
				Skill{
					Name:  "Farm",
					Value: 30,
				},
				Skill{
					Name:       "Craft",
					UserChoice: true,
					CoreString: "Craft",
					UserString: "Arms",
					Base:       10,
					Value:      15,
					Category:   "Manipulation",
				},
				Skill{
					Name:  "First Aid",
					Value: 10,
				},
				Skill{
					Name:  "Scan",
					Value: 10,
				},
				Skill{
					Name:  "Herd",
					Value: 15,
				},
				Skill{
					Name:  "Manage Household",
					Value: 30,
				},
				Skill{
					Name:  "Medium Shield",
					Value: 15,
				},
				Skill{
					Name:  "Broadsword",
					Value: 15,
				},
			},
			// Skill Choices
			SkillChoices: []SkillChoice{
				// Choice of 2 skills
				SkillChoice{
					Skills: []Skill{
						// Skill 1
						Skill{
							Name:  "Jump",
							Value: 10,
						},
						// Skill 2
						Skill{
							Name:  "Climb",
							Value: 10,
						},
					},
				},
			},
		},
		// Esrolia
	}

	c.Occupation = Occupations["Farmer"]

	for _, s := range c.Occupation.SkillList {
		c.ModifySkill(s)
	}
}
