package runequest

import "fmt"

// Homeland represents a homeland and cultural learnings
type Homeland struct {
	Name        string
	Description string
	SkillList   []Skill
	AbilityList []Ability
}

// ChooseHomeland modifies a character's skills by homeland
func (c *Character) ChooseHomeland() {

	// Homelands is a map of possible homelands in Runequest
	var Homelands = map[string]Homeland{
		"Sartar": Homeland{
			Name: "Sartar",
			SkillList: []Skill{
				Skill{
					Name:  "Ride",
					Value: 5,
				},
				Skill{
					Name:  "Dance",
					Value: 5,
				},
				Skill{
					Name:  "Sing",
					Value: 5,
				},
				Skill{
					Name:     "Speak (Heortling)",
					Base:     50,
					Category: "Communication",
				},
				Skill{
					Name:     "Customs (Heortling)",
					Base:     50,
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
					Name:  "1H Axe",
					Value: 10,
				},
				Skill{
					Name:  "1H Spear",
					Value: 10,
				},
			},
		},
		// Esrolia
	}

	c.Homeland = Homelands["Sartar"]

	for _, s := range c.Homeland.SkillList {
		fmt.Println(s.Name)
		c.ModifySkill(s)
	}
}
