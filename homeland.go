package runequest

import "fmt"

// Homeland represents a homeland and cultural learnings
type Homeland struct {
	Name           string
	Description    string
	SkillList      []Skill
	SkillChoices   []SkillChoice
	AbilityList    []Ability
	AbilityChoices []AbilityChoice
}

// ChooseHomeland modifies a character's skills by homeland
func (c *Character) ChooseHomeland() {

	// Homelands is a map of possible homelands in Runequest
	var Homelands = map[string]Homeland{
		// Sartar
		"Sartar": Homeland{
			Name: "Sartar",
			SkillList: []Skill{
				Skill{
					Name:       "Ride",
					CoreString: "Ride",
					UserChoice: true,
					UserString: "Horse",
					Base:       5,
					Value:      5,
					Category:   "Agility",
				},
				Skill{
					Name:  "Dance",
					Value: 5,
				},
				Skill{
					Name:  "Sing",
					Value: 10,
				},
				Skill{
					Name:     "Speak (Heortling)",
					Base:     50,
					Category: "Communication",
				},
				Skill{
					Name:     "Speak (Tradetalk)",
					Value:    10,
					Category: "Communication",
				},
				Skill{
					Name:     "Customs (Heortling)",
					Base:     25,
					Category: "Knowledge",
				},
				Skill{
					Name:  "Farm",
					Value: 20,
				},
				Skill{
					Name:  "Herd",
					Value: 10,
				},
				Skill{
					Name:  "Spirit Combat",
					Value: 15,
				},
				Skill{
					Name:  "Dagger",
					Value: 10,
				},
				Skill{
					Name:  "Broadsword",
					Value: 15,
				},
				Skill{
					Name:  "Battle Axe",
					Value: 10,
				},
				Skill{
					Name:  "1H Spear",
					Value: 10,
				},
				Skill{
					Name:  "Javelin",
					Value: 10,
				},
				Skill{
					Name:  "Medium Shield",
					Value: 15,
				},
				Skill{
					Name:  "Large Shield",
					Value: 10,
				},
			},
			// Skill Choices
			SkillChoices: []SkillChoice{
				// Choice of 2 skills
				SkillChoice{
					Skills: []Skill{
						// Skill 1
						Skill{
							Name:  "Composite Bow",
							Base:  5,
							Value: 10,
						},
						// Skill 2
						Skill{
							Name:  "Sling",
							Value: 10,
						},
					},
				},
			},
		},
		// Esrolia
	}

	c.Homeland = Homelands["Sartar"]

	for _, s := range c.Homeland.SkillList {
		fmt.Println(s)
		c.ModifySkill(s)
	}

	for _, choice := range c.Homeland.SkillChoices {
		// Find number of skills
		l := len(choice.Skills)

		// Choose random index
		r := ChooseRandom(l)

		// Select index from choice.Skills
		selected := choice.Skills[r]

		// Modify or add skill
		c.ModifySkill(selected)
	}
}
