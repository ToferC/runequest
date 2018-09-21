package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {
	c := runequest.NewCharacter("Harthor")

	c.Statistics["STR"].Value += 3

	c.ChooseRunes()
	c.UpdateCharacter()

	c.ChooseHomeland()
	c.ChooseOccupation()
	c.ChooseCult()

	c.ModifySkill(
		runequest.Skill{
			Name:  "Broadsword",
			Value: 25,
		})

	c.ModifyAbility(
		runequest.Ability{
			Name:  "Loyalty (Clan)",
			Value: 25,
		})

	c.UpdateCharacter()
	fmt.Println("")
	fmt.Println(c)
}
