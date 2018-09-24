package runequest

import (
	"fmt"
)

// Ability represents any non-Ability % ability in Runequest
type Ability struct {
	Name            string
	Type            string
	OpposedAbility  string
	UserChoice      bool
	CoreString      string
	UserString      string
	Base            int
	Value           int
	Total           int
	Max             int
	Min             int
	ExperienceCheck bool
}

// AbilityChoice is a choice between 2 or more Abilities
type AbilityChoice struct {
	Abilities []Ability
}

// AbilityTypes is a slice of potential types for abilities
var AbilityTypes = []string{
	"Base",
	"Elemental Rune",
	"Power Rune",
	"Passion",
}

// ElementalRuneOrder sets the order for Elemental Runes
var ElementalRuneOrder = []string{
	"Air", "Darkness", "Earth", "Fire/Sky", "Moon", "Water",
}

// PowerRuneOrder sets the order for Power Runes
var PowerRuneOrder = []string{
	"Man", "Beast",
	"Fertility", "Death",
	"Harmony", "Disorder",
	"Truth", "Illusion",
	"Movement", "Stasis",
}

func (a *Ability) String() string {

	var text = ""

	a.Total = a.Base + a.Value

	text += fmt.Sprintf("%s %d%%", a.Name, a.Total)

	return text
}

// ModifyAbility adds or modifies a Ability value
func (c *Character) ModifyAbility(s Ability) {

	var response, AbilityName string

	if s.UserChoice {

		// Show slice of existing Abilities with identical CoreString

		q := fmt.Sprintf("Enter a specialization for %s or hit Enter to use (%s): ",
			s.CoreString, s.UserString)

		response = UserQuery(q)

		if response == "" {
			response = s.UserString
		}

		AbilityName = fmt.Sprintf("%s (%s)", s.CoreString, response)
	} else {
		AbilityName = s.Name
	}

	if c.Abilities[AbilityName] == nil {
		// Create new Ability in map
		c.Abilities[AbilityName] = &Ability{
			Name:            AbilityName,
			CoreString:      s.CoreString,
			OpposedAbility:  s.OpposedAbility,
			UserChoice:      false,
			UserString:      response,
			Base:            s.Base,
			Value:           s.Value,
			Total:           s.Total,
			Min:             s.Min,
			Max:             s.Max,
			ExperienceCheck: s.ExperienceCheck,
		}
		// Remove base Ability from map
		delete(c.Abilities, s.CoreString)
	} else {
		// Modify existing Ability
		if c.Abilities[AbilityName].Base < s.Base {
			// Change Ability.Base if advantageous
			c.Abilities[AbilityName].Base = s.Base
		}
		// Add or subtract s.Value from Ability
		c.Abilities[AbilityName].Value += s.Value
	}

	ability := c.Abilities[AbilityName]

	// Modify Opposing Rune if required
	if ability.OpposedAbility != "" {
		opposed := c.Abilities[ability.OpposedAbility]

		ability.Total = ability.Base + ability.Value

		// Maximum of 99 in a Power Rune
		if ability.Total > 99 {
			ability.Total = 99
		}

		opposed.Total = opposed.Base + opposed.Value

		diff := ability.Total + opposed.Total

		if diff > 100 {
			opposed.Base -= diff - 100
		}
	}
}

// ChooseRunes selects Runes for a Character
func (c *Character) ChooseRunes() {

	// Choose Rune Values
	c.ModifyAbility(Ability{
		Name: "Air",
		Base: 60,
	})

	c.ModifyAbility(Ability{
		Name: "Earth",
		Base: 40,
	})

	c.ModifyAbility(Ability{
		Name: "Fire/Sky",
		Base: 20,
	})

	c.ModifyAbility(Ability{
		Name:  "Movement",
		Value: 25,
	})

	c.ModifyAbility(Ability{
		Name:  "Man",
		Value: 25,
	})

	// Add 50 more points
	c.ModifyAbility(Ability{
		Name:  "Movement",
		Value: 10,
	})

	c.ModifyAbility(Ability{
		Name:  "Air",
		Value: 30,
	})

	c.ModifyAbility(Ability{
		Name:  "Earth",
		Value: 10,
	})
}

// Abilities is a map of the basic abilities in Runequest
var Abilities = map[string]*Ability{
	"Reputation": &Ability{
		Name:  "Reputation",
		Type:  "Base",
		Value: 5,
	},
	// Elemental Runes
	"Fire": &Ability{
		Name:  "Fire",
		Type:  "Elemental Rune",
		Value: 0,
	},
	"Air": &Ability{
		Name:  "Air",
		Type:  "Elemental Rune",
		Value: 0,
	},
	"Water": &Ability{
		Name:  "Water",
		Type:  "Elemental Rune",
		Value: 0,
	},
	"Earth": &Ability{
		Name:  "Earth",
		Type:  "Elemental Rune",
		Value: 0,
	},
	"Fire/Sky": &Ability{
		Name:  "Fire/Sky",
		Type:  "Elemental Rune",
		Value: 0,
	},
	"Moon": &Ability{
		Name:  "Moon",
		Type:  "Elemental Rune",
		Value: 0,
	},
	// Power Runes
	"Man": &Ability{
		Name:           "Man",
		Type:           "Power Rune",
		OpposedAbility: "Beast",
		Base:           50,
	},
	"Beast": &Ability{
		Name:           "Beast",
		Type:           "Power Rune",
		OpposedAbility: "Man",
		Base:           50,
	},
	"Fertility": &Ability{
		Name:           "Fertility",
		Type:           "Power Rune",
		OpposedAbility: "Death",
		Base:           50,
	},
	"Death": &Ability{
		Name:           "Death",
		Type:           "Power Rune",
		OpposedAbility: "Fertility",
		Base:           50,
	},
	"Harmony": &Ability{
		Name:           "Harmony",
		Type:           "Power Rune",
		OpposedAbility: "Disorder",
		Base:           50,
	},
	"Disorder": &Ability{
		Name:           "Disorder",
		Type:           "Power Rune",
		OpposedAbility: "Harmony",
		Base:           50,
	},
	"Truth": &Ability{
		Name:           "Truth",
		Type:           "Power Rune",
		OpposedAbility: "Illusion",
		Base:           50,
	},
	"Illusion": &Ability{
		Name:           "Illusion",
		Type:           "Power Rune",
		OpposedAbility: "Truth",
		Base:           50,
	},
	"Stasis": &Ability{
		Name:           "Stasis",
		Type:           "Power Rune",
		OpposedAbility: "Movement",
		Base:           50,
	},
	"Movement": &Ability{
		Name:           "Movement",
		Type:           "Power Rune",
		OpposedAbility: "Stasis",
		Base:           50,
	},
	// Passions
	"Loyalty": &Ability{
		Name:  "Loyalty (Clan)",
		Type:  "Passion",
		Base:  60,
		Value: 0,
	},
	"Devotion": &Ability{
		Name:  "Loyalty (Orlanth)",
		Type:  "Passion",
		Base:  60,
		Value: 0,
	},
}
