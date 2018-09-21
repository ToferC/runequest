package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {
	c := runequest.NewCharacter("Harthor")

	c.Statistics["STR"].Value += 3

	c.ChooseRunes()

	c.AddRuneModifiers()

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
			Name:  "Loyalty (Kallyr Starbrow)",
			Base:  60,
			Value: 25,
			Type:  "Passion",
		})

	c.ModifyAbility(
		runequest.Ability{
			Name:  "Loyalty (Kallyr Starbrow)",
			Base:  60,
			Value: 10,
			Type:  "Passion",
		})

	mp := c.DetermineMagicPoints()
	hp := c.DetermineHitPoints()
	hr := c.DetermineHealingRate()
	db := c.DetermineDamageBonus()
	sd := c.DetermineSpiritDamage()
	dx := c.DetermineDexStrikeRank()
	sz := c.DetermineSizStrikeRank()

	c.DerivedStats = map[string]*runequest.DerivedStat{
		"MP":    mp,
		"HP":    hp,
		"HR":    hr,
		"DB":    db,
		"SD":    sd,
		"DEXSR": dx,
		"SIZSR": sz,
	}

	c.UpdateCharacter()

	fmt.Println("")
	fmt.Println(c)
}
