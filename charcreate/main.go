package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {
	c := runequest.NewCharacter("Harthor")

	m := runequest.Modifier{
		Object: c.Statistics["STR"],
		Value:  3,
	}

	m.ModifyValue()

	c.ChooseRunes()
	c.UpdateCharacter()

	c.ChooseHomeland()
	c.ChooseOccupation()
	c.ChooseCult()

	mSkill := runequest.Modifier{
		Object: c.Skills["Broadsword"],
		Value:  25,
	}

	mSkill.ModifyValue()

	mPassion := runequest.Modifier{
		Object: c.Abilities["Loyalty"],
		Value:  10,
	}

	mPassion.ModifyValue()

	fmt.Println("")
	fmt.Println(c)
}
