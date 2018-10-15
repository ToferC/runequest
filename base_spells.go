package runequest

// RuneSpells is the base rune magic list
var RuneSpells = []Spell{
	Spell{
		CoreString:  "Absorption",
		Description: "",
		Domain:      "Rune",
		Source:      "Earth Dark",
		Variable:    false,
		Cost:        1,
	},
	Spell{
		CoreString:  "Absorption",
		Description: "",
		Domain:      "Rune",
		Source:      "Earth Plant",
		Variable:    false,
		Cost:        1,
	},
	Spell{
		CoreString:  "Affix Darkness",
		Description: "",
		Domain:      "Rune",
		Source:      "Dark",
		Variable:    false,
		Cost:        1,
	},
	Spell{
		CoreString:  "Alter Creature",
		Description: "",
		Domain:      "Rune",
		Source:      "Man",
		Variable:    false,
		Cost:        1,
	},
	Spell{
		CoreString:  "Analyxe Magic",
		Description: "",
		Domain:      "Rune",
		Source:      "Sky/Wind",
		Variable:    false,
		Cost:        1,
	},
	Spell{
		CoreString:  "Arouse Passion",
		Description: "",
		Domain:      "Rune",
		Source:      "Earth",
		Variable:    false,
		Cost:        1,
	},
	Spell{
		CoreString:  "Axe Trance",
		Description: "",
		Domain:      "Rune",
		Source:      "Death",
		Variable:    false,
		Cost:        1,
	},
	Spell{
		CoreString:  "Axis Mundi",
		Description: "",
		Domain:      "Rune",
		Source:      "Magic",
		Variable:    false,
		Cost:        1,
	},
	Spell{
		CoreString:  "Ban",
		Description: "",
		Domain:      "Rune",
		Source:      "Magic",
		Variable:    false,
		Cost:        1,
	},
}

// SpiritMagicSpells represents spirit magic
var SpiritMagicSpells = []Spell{
	Spell{
		CoreString:  "Befuddle",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        2,
	},
	Spell{
		CoreString:  "Binding Enchantment",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
	},
	Spell{
		CoreString:  "Bladesharp",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Variable:    true,
	},
	Spell{
		CoreString:  "Bludgeon",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Variable:    true,
	},
	Spell{
		CoreString:  "Control",
		Description: "",
		UserChoice:  true,
		UserString:  "Fire Elemental",
		Domain:      "Spirit",
		Source:      "POW",
		Variable:    true,
	},
	Spell{
		CoreString:  "Coordination",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        2,
		Variable:    true,
	},
	Spell{
		CoreString:  "Countermagic",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Variable:    true,
	},
	Spell{
		CoreString:  "Darkwall",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        2,
		Variable:    false,
	},
	Spell{
		CoreString:  "Demoralize",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        2,
		Variable:    false,
	},
	Spell{
		CoreString:  "Detect Enemies",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Detect Life",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Detect Magic",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Detect Spirit",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Detect (Substance)",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Detect Trap",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Detect Undead",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Dispel Magic",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Disruption",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Distraction",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Dullblade",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Extinguish",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Fanaticism",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Farsee",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Firearrow",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        2,
		Variable:    false,
	},
	Spell{
		CoreString:  "Fireblade",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        4,
		Variable:    false,
	},
	Spell{
		CoreString:  "Glamour",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        2,
		Variable:    false,
	},
	Spell{
		CoreString:  "Glue",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Heal",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Ignite",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Ironhand",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Lantern",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Light",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Magic Point Enchantment",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Mobility",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Multimissile",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Parry",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Protection",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Repair",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Rivereyes",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Second Sight",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        3,
		Variable:    false,
	},
	Spell{
		CoreString:  "Shimmer",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Silence",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Sleep",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        3,
		Variable:    false,
	},
	Spell{
		CoreString:  "Slow",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Disruption",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Speedart",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Spell Matrix Enchantment",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Spirit Binding",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        1,
		Variable:    false,
	},
	Spell{
		CoreString:  "Spirit Screen",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Strength",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        2,
		Variable:    false,
	},
	Spell{
		CoreString:  "Summon",
		UserChoice:  true,
		UserString:  "Fire Elemental",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        0,
		Variable:    true,
	},
	Spell{
		CoreString:  "Vigor",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        2,
		Variable:    false,
	},
	Spell{
		CoreString:  "Visibility",
		Description: "",
		Domain:      "Spirit",
		Source:      "POW",
		Cost:        2,
		Variable:    false,
	},
}