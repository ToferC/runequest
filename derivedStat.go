package runequest

import "fmt"

// DerivedStat is a Character element that is based off other elements
type DerivedStat struct {
	Name     string
	MaxValue int
	Value    int
	Base     int
	Total    int
	Dice     int
	Max      int
	Min      int
	Text     string
}

// UpdateDerivedStats totals base & value for DerivedStat
func (c *Character) UpdateDerivedStats() {
	for _, v := range c.DerivedStats {
		v.Total = v.Base + v.Value
	}
}

func (d *DerivedStat) String() string {

	d.Total = d.Base + d.Value

	var text string

	if d.Text == "" {
		text = fmt.Sprintf("%s: %d", d.Name, d.Total)
	} else {
		text = fmt.Sprintf("%s: %s", d.Name, d.Text)
	}
	return text
}

// DetermineMagicPoints calculates magic points for a Character
func (c *Character) DetermineMagicPoints() *DerivedStat {
	mp := &DerivedStat{
		Name:     "Magic Points",
		MaxValue: 21,
	}

	p := c.Statistics["POW"]
	p.Total = p.Base + p.Value

	pow := p.Total

	mp.Base = pow

	return mp
}

// DetermineHitPoints calculates hit points for a Character
func (c *Character) DetermineHitPoints() *DerivedStat {

	hp := &DerivedStat{
		Name:     "Hit Points",
		MaxValue: 21,
	}

	s := c.Statistics["SIZ"]
	s.Total = s.Base + s.Value

	siz := s.Total

	p := c.Statistics["POW"]
	p.Total = p.Base + p.Value

	pow := p.Total

	fmt.Println("SIZ ", siz)
	fmt.Println("POW ", pow)

	con := c.Statistics["CON"]
	con.Total = con.Base + con.Value

	fmt.Println("CON ", con.Total)

	baseHP := con.Total

	switch {
	case siz < 5:
		baseHP -= 2
		fmt.Println("-2")
	case siz < 9:
		baseHP--
		fmt.Println("-1")
	case siz < 13:
		fmt.Println("No change")
	case siz < 17:
		baseHP++
		fmt.Println("+1")
	case siz < 21:
		baseHP += 2
		fmt.Println("+2")
	case siz < 25:
		baseHP += 3
		fmt.Println("+3")
	case siz > 24:
		baseHP += ((siz - 24) / 4) + 4
	}

	switch {
	case pow < 5:
		baseHP--
	case pow < 9:
		fmt.Println("No change")
	case pow < 13:
		fmt.Println("No change")
	case pow < 17:
		fmt.Println("No change")
	case pow < 21:
		baseHP++
	case pow < 25:
		baseHP += 2
	case pow > 24:
		baseHP += ((pow - 24) / 4) + 3
	}

	hp.Base = baseHP

	return hp
}

// DetermineHealingRate determines the Character's healingrate based on Con
func (c *Character) DetermineHealingRate() *DerivedStat {

	healingRate := &DerivedStat{
		Name: "Healing Rate",
		Max:  21,
	}

	con := c.Statistics["CON"]

	con.Total = con.Base + con.Value
	tCon := con.Total

	switch {
	case tCon < 7:
		healingRate.Base = 1
	case tCon < 13:
		healingRate.Base = 2
	case tCon < 19:
		healingRate.Base = 3
	case tCon > 18:
		healingRate.Base = ((tCon - 18) / 6) + 3
	}
	healingRate.Total = healingRate.Base + healingRate.Value
	return healingRate
}

// DetermineDamageBonus determines the Character's healingrate based on Con
func (c *Character) DetermineDamageBonus() *DerivedStat {

	damageBonus := &DerivedStat{
		Name: "Damage Bonus",
		Max:  21,
		Dice: 1,
	}

	str := c.Statistics["STR"]
	siz := c.Statistics["SIZ"]

	str.Total = str.Base + str.Value
	siz.Total = siz.Base + siz.Value

	db := siz.Total + str.Total

	switch {
	case db < 13:
		damageBonus.Base = -4
		damageBonus.Text = "-1D4"
	case db < 25:
		damageBonus.Base = 0
		damageBonus.Text = "-"
	case db < 33:
		damageBonus.Base = 4
		damageBonus.Text = "+1D4"
	case db < 41:
		damageBonus.Base = 6
		damageBonus.Text = "+1D6"
	case db < 57:
		damageBonus.Base = 6
		damageBonus.Dice = 2
		damageBonus.Text = "+2D6"
	case db > 56:
		damageBonus.Base = 6
		damageBonus.Dice = ((db - 56) / 16) + 2
		damageBonus.Text = fmt.Sprintf("+%dD%d",
			damageBonus.Dice,
			damageBonus.Base,
		)
	}

	return damageBonus
}

// DetermineSpiritDamage determines the Character's healingrate based on Con
func (c *Character) DetermineSpiritDamage() *DerivedStat {

	damage := &DerivedStat{
		Name: "Spirit Damage",
		Max:  21,
		Dice: 1,
	}

	pow := c.Statistics["POW"]
	cha := c.Statistics["CHA"]

	pow.Total = pow.Base + pow.Value
	cha.Total = cha.Base + cha.Value

	db := pow.Total + cha.Total

	switch {
	case db < 13:
		damage.Base = 3
		damage.Text = "1D3"
	case db < 25:
		damage.Base = 6
		damage.Text = "1D6"
	case db < 33:
		damage.Base = 6
		damage.Value = 1
		damage.Text = "1D6+1"
	case db < 41:
		damage.Base = 6
		damage.Value = 3
		damage.Text = "1D6+3"
	case db < 57:
		damage.Base = 6
		damage.Dice = 2
		damage.Value = 3
		damage.Text = "2D6+3"
	case db > 56:
		damage.Base = 6
		damage.Dice = ((db - 56) / 16) + 2
		damage.Value = ((db - 56) / 16) + 3
		damage.Text = fmt.Sprintf("%dD%d+%d",
			damage.Dice,
			damage.Base,
			damage.Value,
		)
	}
	return damage
}

// DetermineDexStrikeRank determines the Character's healingrate based on Con
func (c *Character) DetermineDexStrikeRank() *DerivedStat {

	dexSR := &DerivedStat{
		Name: "DEX Strike Rank",
		Max:  5,
	}

	dex := c.Statistics["DEX"]

	dex.Total = dex.Base + dex.Value

	switch {
	case dex.Total < 6:
		dexSR.Base = 5
	case dex.Total < 9:
		dexSR.Base = 4
	case dex.Total < 13:
		dexSR.Base = 3
	case dex.Total < 16:
		dexSR.Base = 2
	case dex.Total < 19:
		dexSR.Base = 1
	case dex.Total > 18:
		dexSR.Base = 0
	}
	return dexSR
}

// DetermineSizStrikeRank determines the Character's healingrate based on Con
func (c *Character) DetermineSizStrikeRank() *DerivedStat {

	sizSR := &DerivedStat{
		Name: "SIZ Strike Rank",
		Max:  5,
	}

	siz := c.Statistics["SIZ"]

	siz.Total = siz.Base + siz.Value

	switch {
	case siz.Total < 7:
		sizSR.Base = 3
	case siz.Total < 15:
		sizSR.Base = 2
	case siz.Total < 22:
		sizSR.Base = 1
	case siz.Total > 21:
		sizSR.Base = 0
	}
	return sizSR
}
