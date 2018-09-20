package runequest

// Occupation represents a profession in Runequest
type Occupation struct {
	Name             string
	Description      string
	Modifiers        []*Modifier
	StandardOfLiving string
	Income           int
	Cults            []*Cult
	Passions         []string
	Ransom           int
	Equipment        []string
}

// ChooseOccupation modifies a character's skills by homeland
func (c *Character) ChooseOccupation() {

	// Homelands is a map of possible homelands in Runequest
	var Occupations = map[string]Occupation{
		"Farmer": Occupation{
			Name: "Farmer",
			Modifiers: []*Modifier{
				&Modifier{
					Object: c.Skills["Homeland Lore (local)"],
					Value:  15,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Jump"],
					Value:  10,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Farm"],
					Value:  30,
					Modify: true,
				},
				&Modifier{
					Object:  c.Skills["Craft"],
					Value:   15,
					Base:    false,
					Modify:  true,
					Subject: "Armoury",
				},
				&Modifier{
					Object: c.Skills["First Aid"],
					Value:  10,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Scan"],
					Value:  10,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Herd"],
					Value:  15,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Manage Household"],
					Value:  30,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Medium Shield"],
					Value:  15,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Broadsword"],
					Value:  15,
					Modify: true,
				},
			},
		},
		// Esrolia
	}

	c.Occupation = Occupations["Farmer"]

	for _, m := range c.Occupation.Modifiers {

		switch {
		case m.Set && m.Base:
			m.SetBase()
		case m.Set && !m.Base:
			m.SetValue()
		case m.Modify:
			m.ModifyValue()
		}
	}
}
