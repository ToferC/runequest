package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {
	c := runequest.NewCharacter("Bob")

	c.Description = "Man"

	grandparentHomeland := "Esrolia"
	grandparentOccupation := "Hunter"

	for _, v := range runequest.PersonalHistoryEvents {

		hlBonus := v.HomelandModifiers[grandparentHomeland]
		ocBonus := v.OccupationModifiers[grandparentOccupation]

		roll := runequest.RollDice(20, 1, hlBonus+ocBonus, 1)

		if roll > 20 {
			roll = 20
		}

		fmt.Println(roll)

		fmt.Println(c.Skills["Dodge"])
		fmt.Println(c.Skills["Bargain"])

		for _, r := range v.Results {
			if runequest.IsInIntArray(r.Range, roll) {
				fmt.Println(r.Description)

				c.Lunars += r.Lunars
				c.Abilities["Reputation"].CreationBonusValue += r.Reputation

				text := r.Description

				for _, s := range r.Skills {
					// Update skills
					text += "\nSkills:\n"

					c.Skills[s.Name].Updates = append(c.Skills[s.Name].Updates, s.Updates[0])
					text += fmt.Sprintf("\n%s +%d\n", s.Name, s.Updates[0].Value)
				}

				for _, p := range r.Passions {
					// Update skills
					text += "\nPassions:\n"

					_, ok := c.Abilities[p.Name]

					if !ok {
						c.Abilities[p.Name] = &runequest.Ability{
							CoreString: p.CoreString,
							UserString: p.UserString,
							Base:       60,
							Updates:    []*runequest.Update{},
						}
					}

					c.Abilities[p.Name].Updates = append(c.Abilities[p.Name].Updates, p.Updates[0])
					text += fmt.Sprintf("\n%s +%d\n", p.Name, p.Updates[0].Value)
				}

				c.History[fmt.Sprintf("%d - %s:", v.Year, v.Name)] = text
			}
		}
	}

	// Character done
	for k, v := range c.History {
		fmt.Println(k, v)
	}
}
