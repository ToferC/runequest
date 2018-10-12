package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {
	c := runequest.NewCharacter("Harthor")

	c.Statistics["STR"].Value += 3

	c.ChooseRunes()

	// Add Modifiers to Stats
	c.AddRuneModifiers()

	// Calculate Derived Stats
	c.SetAttributes()

	for _, s := range c.Skills {

		if s.UserChoice {
			c.ModifySkill(runequest.Skill{
				Name:       s.Name,
				Category:   s.Category,
				CoreString: s.CoreString,
				UserChoice: s.UserChoice,
				UserString: s.UserString,
				Base:       s.Base,
				Value:      s.Value,
			})
		}
	}

	c.UpdateCharacter()

	fmt.Println("Homeland")
	c.ChooseHomeland(runequest.Homelands["Sartar"])
	fmt.Println("Occupation")
	c.ChooseOccupation(runequest.Occupations["Farmer"])
	fmt.Println("Cult")
	c.ChooseCult(runequest.Cults["Orlanth"])

	c.ModifySkill(
		runequest.Skill{
			CoreString: "Broadsword",
			Value:      25,
		})

	c.ModifyAbility(
		runequest.Ability{
			Name:       "Loyalty",
			CoreString: "Loyalty",
			UserString: "Kallyr Starbrow",
			Base:       60,
			Value:      25,
			Type:       "Passion",
		})

	c.ModifyAbility(
		runequest.Ability{
			Name:       "Loyalty",
			CoreString: "Loyalty",
			UserString: "Kallyr Starbrow",
			Base:       60,
			Value:      10,
			Type:       "Passion",
		})

	fmt.Println("")
	fmt.Println(c)
}
