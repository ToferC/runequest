package runequest

// Homeland represents a homeland and cultural learnings
type Homeland struct {
	Name        string
	Description string
	Modifiers   []*Modifier
}

// ChooseHomeland modifies a character's skills by homeland
func (c *Character) ChooseHomeland() {

	// Homelands is a map of possible homelands in Runequest
	var Homelands = map[string]Homeland{
		"Sartar": Homeland{
			Name: "Sartar",
			Modifiers: []*Modifier{
				&Modifier{
					Object: c.Skills["Ride"],
					Value:  5,
					Set:    false,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Dance"],
					Value:  5,
					Set:    false,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Sing"],
					Value:  5,
					Set:    false,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Speak Own Language"],
					Value:  50,
					Base:   true,
					Set:    true,
					Modify: false,
				},
				&Modifier{
					Object:  c.Skills["Customs"],
					Value:   50,
					Base:    true,
					Set:     true,
					Modify:  false,
					Subject: "Heortling",
				},
				&Modifier{
					Object: c.Skills["Farm"],
					Value:  20,
					Set:    false,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Herd"],
					Value:  10,
					Set:    false,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Spirit Combat"],
					Value:  15,
					Set:    false,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Dagger"],
					Value:  10,
					Set:    false,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["1H Axe"],
					Value:  10,
					Set:    false,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["1H Spear"],
					Value:  10,
					Set:    false,
					Modify: true,
				},
			},
		},
		// Esrolia
	}

	c.Homeland = Homelands["Sartar"]

	for _, m := range c.Homeland.Modifiers {

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
