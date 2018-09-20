package runequest

import "fmt"

// Homeland represents a homeland and cultural learnings
type Homeland struct {
	Name        string
	Description string
	Modifiers   []*Modifier
}

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
					Object: c.Skills["Customs (Heortling)"],
					Value:  50,
					Base:   true,
					Set:    true,
					Modify: false,
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
					Object: c.Skills["Battle Axe"],
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

	for _, m := range Homelands["Sartar"].Modifiers {

		switch {
		case m.Set && m.Base:
			m.SetBase()
			fmt.Println("set base" + string(m.Value))
		case m.Set && !m.Base:
			m.SetValue()
			fmt.Println("set value" + string(m.Value))
		case m.Modify:
			m.ModifyValue()
			fmt.Println("set value" + string(m.Value))
		}
	}
}
