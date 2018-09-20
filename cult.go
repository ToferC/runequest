package runequest

// Cult represents a Religion in Runequest
type Cult struct {
	Name        string
	Description string
	Modifiers   []*Modifier
	SpellList   []*Spell
}

// ChooseCult modifies a character's skills by homeland
func (c *Character) ChooseCult() {

	// Homelands is a map of possible homelands in Runequest
	var Cults = map[string]Cult{
		"Orlanth": Cult{
			Name: "Orlanth",
			Modifiers: []*Modifier{
				&Modifier{
					Object:  c.Skills["Cult Lore"],
					Value:   15,
					Modify:  true,
					Subject: "Orlanth",
				},
				&Modifier{
					Object: c.Skills["Worship"],
					Value:  15,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Meditate"],
					Value:  25,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Orate"],
					Value:  30,
					Base:   false,
					Modify: true,
				},
				&Modifier{
					Object: c.Skills["Sing"],
					Value:  10,
					Modify: true,
				},
				&Modifier{
					Object:  c.Skills["Speak Other Language"],
					Value:   10,
					Modify:  true,
					Subject: "Stormspeech",
				},
			},
		},
		// Esrolia
	}

	c.Cult = Cults["Orlanth"]

	for _, m := range c.Cult.Modifiers {

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
