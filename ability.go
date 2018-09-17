package runequest

import "fmt"

// Ability represents any non-Skill % ability in Runequest
type Ability struct {
	Name            string
	Type            string
	TakesSubject    bool
	Subject         string
	OpposedAbility  string
	Base            int
	Value           int
	Total           int
	Max             int
	Min             int
	ExperienceCheck bool
}

// AbilityTypes is a slice of potential types for abilities
var AbilityTypes = []string{
	"Base",
	"Elemental Rune",
	"Power Rune",
	"Passion",
}

func (a *Ability) String() string {

	var text = ""

	a.Total = a.Base + a.Value

	if a.TakesSubject {
		text += fmt.Sprintf("%s (%s) %d%%", a.Name, a.Subject, a.Total)
	} else {
		text += fmt.Sprintf("%s %d%%", a.Name, a.Total)
	}
	return text
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
	"Loyalty (Clan)": &Ability{
		Name:    "Loyalty",
		Type:    "Passion",
		Subject: "Clan",
		Base:    60,
		Value:   0,
	},
}
