package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {
	c := runequest.NewCharacter("Bigby")

	m := runequest.Modifier{
		Object: c.Statistics["STR"],
		Value:  3,
	}

	m.ModifyValue()

	// Choose Rune Values
	mAbility1 := runequest.Modifier{
		Object: c.Abilities["Air"],
		Value:  60,
	}

	mAbility1.ModifyValue()

	mAbility2 := runequest.Modifier{
		Object: c.Abilities["Earth"],
		Value:  40,
	}

	mAbility2.ModifyValue()

	mAbility3 := runequest.Modifier{
		Object: c.Abilities["Fire/Sky"],
		Value:  20,
	}

	mAbility3.ModifyValue()

	c.AddRuneModifiers()

	c.DetermineSkillCategoryValues()

	mSkill := runequest.Modifier{
		Object: c.Skills["Broadsword"],
		Value:  25,
	}

	mSkill.ModifyValue()

	mPassion := runequest.Modifier{
		Object: c.Abilities["Loyalty (Clan)"],
		Value:  10,
	}

	mPassion.ModifyValue()

	fmt.Println("")
	fmt.Println(c)
}
