package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {
	c := runequest.Character{
		Name:            "Bigby",
		Statistics:      runequest.RuneQuestStats,
		Abilities:       runequest.Abilities,
		Skills:          runequest.Skills,
		SkillCategories: runequest.SkillCategories,
	}

	fmt.Println(c)

	m := runequest.Modifier{
		Object: c.Statistics["STR"],
		Value:  3,
	}

	m.ModifyValue()

	mAbility := runequest.Modifier{
		Object: c.Abilities["Air"],
		Value:  60,
	}

	mAbility.ModifyValue()

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
