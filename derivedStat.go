package runequest

import "fmt"

// DerivedStat is a Character element that is based off other elements
type DerivedStat struct {
	Name     string
	MaxValue int
	Value    int
	Base     int
	Total    int
	Max      int
	Min      int
}

// UpdateDerivedStats totals base & value for DerivedStat
func (c *Character) UpdateDerivedStats() {
	for _, v := range c.DerivedStats {
		v.Total = v.Base + v.Value
	}
}

func (d *DerivedStat) String() string {

	d.Total = d.Base + d.Value

	text := fmt.Sprintf("%s: %d", d.Name, d.Total)
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
