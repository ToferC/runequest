package runequest

import (
	"fmt"
)

// Ability represents any non-Ability % ability in Runequest
type Ability struct {
	Name               string
	Type               string
	OpposedAbility     string
	UserChoice         bool
	CoreString         string
	UserString         string
	Base               int
	HomelandValue      int
	OccupationValue    int
	CultValue          int
	CreationBonusValue int
	Updates            []*Update
	Value              int
	InPlayXPValue      int
	Total              int
	Max                int
	Min                int
	ExperienceCheck    bool
}

// AbilityChoice is a choice between 2 or more skills
type AbilityChoice struct {
	Abilities []Ability
}

// UpdateAbility totals skill values based on input
func (a *Ability) UpdateAbility() {

	a.generateName()

	updates := 0

	for _, u := range a.Updates {
		updates += u.Value
	}

	a.Total = a.Base + a.HomelandValue + a.OccupationValue + a.CultValue + a.CreationBonusValue + a.InPlayXPValue + a.Value + updates
}

// generateName sets the ability map name
func (a *Ability) generateName() {

	var n string

	if a.UserString != "" {
		n = fmt.Sprintf("%s (%s)", a.CoreString, a.UserString)
	} else {
		n = a.CoreString
	}
	a.Name = n
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

// PassionTypes represents different passions in Glorantha
var PassionTypes = []string{
	"Devotion", "Fear", "Hate", "Honor",
	"Loyalty", "Love",
}

func (a *Ability) String() string {

	a.UpdateAbility()

	var text = ""

	text += fmt.Sprintf("%s %d%%", a.Name, a.Total)

	return text
}

// ModifyAbility adds or modifies a Ability value
func (c *Character) ModifyAbility(a Ability) {

	a.generateName()

	if c.Abilities[a.Name] == nil {
		// New Ability
		c.Abilities[a.Name] = &Ability{
			Name:       a.Name,
			CoreString: a.CoreString,
			UserString: a.UserString,
			Base:       a.Base,
			Type:       a.Type,
		}
	} else {
		// Modify existing Ability
		ab := c.Abilities[a.Name]

		if ab.Base == 0 {
			// Change Ability.Base if needed
			ab.Base = a.Base
		} else {
			// Add or subtract s.Value from Ability
			ab.Value += a.Value
			ab.HomelandValue += a.HomelandValue
			ab.OccupationValue += a.OccupationValue
			ab.CultValue += a.CultValue
			ab.CreationBonusValue += a.CreationBonusValue
			ab.InPlayXPValue += a.InPlayXPValue
		}
	}

	ability := c.Abilities[a.Name]

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
		CoreString: "Air",
		Base:       60,
	})

	c.ModifyAbility(Ability{
		CoreString: "Earth",
		Base:       40,
	})

	c.ModifyAbility(Ability{
		CoreString: "Fire/Sky",
		Base:       20,
	})

	c.ModifyAbility(Ability{
		CoreString: "Movement",
		Value:      25,
	})

	c.ModifyAbility(Ability{
		CoreString: "Man",
		Value:      25,
	})

	// Add 50 more points
	c.ModifyAbility(Ability{
		CoreString: "Movement",
		Value:      10,
	})

	c.ModifyAbility(Ability{
		CoreString: "Air",
		Value:      30,
	})

	c.ModifyAbility(Ability{
		CoreString: "Earth",
		Value:      10,
	})
}

// Abilities is a map of the basic abilities in Runequest
var Abilities = map[string]*Ability{
	"Reputation": &Ability{
		CoreString: "Reputation",
		Type:       "Base",
		Value:      5,
	},
}

// ElementalRunes is a map of Rune abilities
var ElementalRunes = map[string]*Ability{
	// Elemental Runes
	"Fire": &Ability{
		CoreString: "Fire",
		Type:       "Elemental Rune",
		Value:      0,
	},
	"Air": &Ability{
		CoreString: "Air",
		Type:       "Elemental Rune",
		Value:      0,
	},
	"Water": &Ability{
		CoreString: "Water",
		Type:       "Elemental Rune",
		Value:      0,
	},
	"Earth": &Ability{
		CoreString: "Earth",
		Type:       "Elemental Rune",
		Value:      0,
	},
	"Fire/Sky": &Ability{
		CoreString: "Fire/Sky",
		Type:       "Elemental Rune",
		Value:      0,
	},
	"Moon": &Ability{
		CoreString: "Moon",
		Type:       "Elemental Rune",
		Value:      0,
	},
}

// PowerRunes is a map of Power Runes
var PowerRunes = map[string]*Ability{
	// Power Runes
	"Man": &Ability{
		CoreString:     "Man",
		Type:           "Power Rune",
		OpposedAbility: "Beast",
		Base:           50,
	},
	"Beast": &Ability{
		CoreString:     "Beast",
		Type:           "Power Rune",
		OpposedAbility: "Man",
		Base:           50,
	},
	"Fertility": &Ability{
		CoreString:     "Fertility",
		Type:           "Power Rune",
		OpposedAbility: "Death",
		Base:           50,
	},
	"Death": &Ability{
		CoreString:     "Death",
		Type:           "Power Rune",
		OpposedAbility: "Fertility",
		Base:           50,
	},
	"Harmony": &Ability{
		CoreString:     "Harmony",
		Type:           "Power Rune",
		OpposedAbility: "Disorder",
		Base:           50,
	},
	"Disorder": &Ability{
		CoreString:     "Disorder",
		Type:           "Power Rune",
		OpposedAbility: "Harmony",
		Base:           50,
	},
	"Truth": &Ability{
		CoreString:     "Truth",
		Type:           "Power Rune",
		OpposedAbility: "Illusion",
		Base:           50,
	},
	"Illusion": &Ability{
		CoreString:     "Illusion",
		Type:           "Power Rune",
		OpposedAbility: "Truth",
		Base:           50,
	},
	"Stasis": &Ability{
		CoreString:     "Stasis",
		Type:           "Power Rune",
		OpposedAbility: "Movement",
		Base:           50,
	},
	"Movement": &Ability{
		CoreString:     "Movement",
		Type:           "Power Rune",
		OpposedAbility: "Stasis",
		Base:           50,
	},
}
