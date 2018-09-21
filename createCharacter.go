package runequest

// NewCharacter generates a random starting character in Runequest
func NewCharacter(name string) *Character {
	c := Character{
		Name:            name,
		Statistics:      RuneQuestStats,
		Abilities:       Abilities,
		SkillCategories: SkillCategories,
		// Skills is a map of regular skills in Runequest
	}

	c.Skills = map[string]*Skill{
		// Agility
		"Boat": &Skill{
			Name:     "Boat",
			Base:     5,
			Category: "Agility",
		},
		"Climb": &Skill{
			Name:     "Climb",
			Base:     40,
			Category: "Agility",
		},
		"Dodge": &Skill{
			Name:     "Dodge",
			Base:     c.Statistics["DEX"].Value * 2,
			Category: "Agility",
		},
		"Drive Chariot": &Skill{
			Name:     "Drive Chariot",
			Base:     5,
			Category: "Agility",
		},
		"Jump": &Skill{
			Name:     "Jump",
			Base:     c.Statistics["DEX"].Value * 3,
			Category: "Agility",
		},
		"Ride": &Skill{
			Name:     "Ride Horse",
			Base:     5,
			Category: "Agility",
		},
		"Swim": &Skill{
			Name:     "Swim",
			Base:     15,
			Category: "Agility",
		},

		// Communication
		"Act": &Skill{
			Name:     "Act",
			Base:     5,
			Category: "Communication",
		},
		"Art": &Skill{
			Name:     "Art",
			Base:     5,
			Category: "Communication",
		},
		"Bargain": &Skill{
			Name:     "Bargain",
			Base:     5,
			Category: "Communication",
		},
		"Charm": &Skill{
			Name:     "Charm",
			Base:     15,
			Category: "Communication",
		},
		"Dance": &Skill{
			Name:     "Dance",
			Base:     10,
			Category: "Communication",
		},
		"Disguise": &Skill{
			Name:     "Disguise",
			Base:     5,
			Category: "Communication",
		},
		"Fast Talk": &Skill{
			Name:     "Fast Talk",
			Base:     5,
			Category: "Communication",
		},
		"Intimidate": &Skill{
			Name:     "Intimidate",
			Base:     15,
			Category: "Communication",
		},
		"Intrigue": &Skill{
			Name:     "Intrigue",
			Base:     5,
			Category: "Communication",
		},
		"Orate": &Skill{
			Name:     "Orate",
			Base:     10,
			Category: "Communication",
		},
		"Sing": &Skill{
			Name:     "Sing",
			Base:     10,
			Category: "Communication",
		},
		"Speak (Heortling)": &Skill{
			Name:     "Speak (Heortling)",
			Base:     50,
			Category: "Communication",
		},
		// Knowledge
		"Alchemy": &Skill{
			Name:     "Alchemy",
			Base:     0,
			Category: "Knowledge",
		},
		"Animal Lore": &Skill{
			Name:     "Animal Lore",
			Base:     5,
			Category: "Knowledge",
		},
		"Battle": &Skill{
			Name:     "Battle",
			Base:     10,
			Category: "Knowledge",
		},
		"Bureacracy": &Skill{
			Name:     "Bureacracy",
			Base:     0,
			Category: "Knowledge",
		},
		"Celestial Lore": &Skill{
			Name:     "Celestial Lore",
			Base:     5,
			Category: "Knowledge",
		},
		"Cult Lore (Orlanth)": &Skill{
			Name:     "Cult Lore (Orlanth)",
			Base:     0,
			Category: "Knowledge",
		},
		"Customs (Heortling)": &Skill{
			Name:     "Customs (Heortling)",
			Base:     25,
			Category: "Knowledge",
		},
		"Elf Lore": &Skill{
			Name:     "Elf Lore",
			Base:     5,
			Category: "Knowledge",
		},
		"Evaluate": &Skill{
			Name:     "Evaluate",
			Base:     10,
			Category: "Knowledge",
		},
		"Farm": &Skill{
			Name:     "Farm",
			Base:     10,
			Category: "Knowledge",
		},
		"First Aid": &Skill{
			Name:     "First Aid",
			Base:     10,
			Category: "Knowledge",
		},
		"Game": &Skill{
			Name:     "Game",
			Base:     15,
			Category: "Knowledge",
		},
		"Herd": &Skill{
			Name:     "Herd",
			Base:     5,
			Category: "Knowledge",
		},
		"Homeland Lore (local)": &Skill{
			Name:     "Homeland Lore (local)",
			Base:     30,
			Category: "Knowledge",
		},
		"Homeland Lore (Spirits)": &Skill{
			Name:     "Homeland Lore (Spirits)",
			Base:     0,
			Category: "Knowledge",
		},
		"Library Use": &Skill{
			Name:     "Library Use",
			Base:     0,
			Category: "Knowledge",
		},
		"Manage Household": &Skill{
			Name:     "Manage Household",
			Base:     10,
			Category: "Knowledge",
		},
		"Mineral Lore": &Skill{
			Name:     "Mineral Lore",
			Base:     5,
			Category: "Knowledge",
		},
		"Peaceful Cut": &Skill{
			Name:     "Peaceful Cut",
			Base:     10,
			Category: "Knowledge",
		},
		"Plant Lore": &Skill{
			Name:     "Plant Lore",
			Base:     5,
			Category: "Knowledge",
		},
		"Read/Write (Tarsh)": &Skill{
			Name:     "Read/Write (Tarsh)",
			Base:     0,
			Category: "Knowledge",
		},
		"Shiphandling": &Skill{
			Name:     "Shiphandling",
			Base:     0,
			Category: "Knowledge",
		},
		"Survival": &Skill{
			Name:     "Survival",
			Base:     15,
			Category: "Knowledge",
		},
		"Treat Disease": &Skill{
			Name:     "Treat Disease",
			Base:     5,
			Category: "Knowledge",
		},
		"Treat Poison": &Skill{
			Name:     "Treat Poison",
			Base:     5,
			Category: "Knowledge",
		},
		// Magic
		"Meditate": &Skill{
			Name:     "Meditate",
			Base:     0,
			Category: "Magic",
		},
		"Prepare Corpse": &Skill{
			Name:     "Prepare Corpse",
			Base:     10,
			Category: "Magic",
		},
		"Sense Assassin": &Skill{
			Name:     "Sense Assassin",
			Base:     0,
			Category: "Magic",
		},
		"Sense Chaos": &Skill{
			Name:     "Sense Chaos",
			Base:     0,
			Category: "Magic",
		},
		"Spirit Combat": &Skill{
			Name:     "Spirit Combat",
			Base:     20,
			Category: "Magic",
		},
		"Spirit Dance": &Skill{
			Name:     "Spirit Dance",
			Base:     0,
			Category: "Magic",
		},
		"Spirit Lore": &Skill{
			Name:     "Spirit Lore",
			Base:     0,
			Category: "Magic",
		},
		"Spirit Travel": &Skill{
			Name:     "Spirit Travel",
			Base:     0,
			Category: "Magic",
		},
		"Understand Herd Beast": &Skill{
			Name:     "Understand Herd Beast",
			Base:     0,
			Category: "Magic",
		},
		"Worship (Orlanth)": &Skill{
			Name:     "Worship (Orlanth)",
			Base:     0,
			Category: "Magic",
		},

		// Manipulation
		"Conceal": &Skill{
			Name:     "Conceal",
			Base:     5,
			Category: "Manipulation",
		},
		"Craft (Weapons)": &Skill{
			Name:     "Craft (Weapons)",
			Base:     10,
			Category: "Manipulation",
		},
		"Devise": &Skill{
			Name:     "Devise",
			Base:     5,
			Category: "Manipulation",
		},
		"Play Instrument": &Skill{
			Name:     "Play Instrument",
			Base:     5,
			Category: "Manipulation",
		},
		"Sleight": &Skill{
			Name:     "Sleight",
			Base:     10,
			Category: "Manipulation",
		},

		// Melee Weapons
		"1H Axe": &Skill{
			Name:     "1H Axe",
			Base:     10,
			Category: "Manipulation",
		},
		"2H Axe": &Skill{
			Name:     "2H Axe",
			Base:     5,
			Category: "Manipulation",
		},
		"Broadsword": &Skill{
			Name:     "Broadsword",
			Base:     10,
			Category: "Manipulation",
		},
		"Dagger": &Skill{
			Name:     "Dagger",
			Base:     15,
			Category: "Manipulation",
		},
		"Fist": &Skill{
			Name:     "Fist",
			Base:     25,
			Category: "Manipulation",
		},
		"Grapple": &Skill{
			Name:     "Grapple",
			Base:     25,
			Category: "Manipulation",
		},
		"1H Hammer": &Skill{
			Name:     "1H Hammer",
			Base:     10,
			Category: "Manipulation",
		},
		"2H Hammer": &Skill{
			Name:     "2H Hammer",
			Base:     5,
			Category: "Manipulation",
		},
		"Kick": &Skill{
			Name:     "Kick",
			Base:     15,
			Category: "Manipulation",
		},
		"Kopis": &Skill{
			Name:     "Kopis",
			Base:     10,
			Category: "Manipulation",
		},
		"1H Mace": &Skill{
			Name:     "1H Mace",
			Base:     15,
			Category: "Manipulation",
		},
		"2H Mace": &Skill{
			Name:     "2H Mace",
			Base:     10,
			Category: "Manipulation",
		},
		"Pike": &Skill{
			Name:     "Pike",
			Base:     15,
			Category: "Manipulation",
		},
		"Quarterstaff": &Skill{
			Name:     "Quarterstaff",
			Base:     15,
			Category: "Manipulation",
		},
		"Rapier": &Skill{
			Name:     "Rapier",
			Base:     5,
			Category: "Manipulation",
		},
		"Shortsword": &Skill{
			Name:     "Shortsword",
			Base:     10,
			Category: "Manipulation",
		},
		"1H Spear": &Skill{
			Name:     "1H Spear",
			Base:     05,
			Category: "Manipulation",
		},
		"2H Spear": &Skill{
			Name:     "2H Spear",
			Base:     15,
			Category: "Manipulation",
		},

		// Missile Weapons
		"Arbalest": &Skill{
			Name:     "Arbalest",
			Base:     10,
			Category: "Manipulation",
		},
		"Axe, Throwing": &Skill{
			Name:     "Axe, Throwing",
			Base:     10,
			Category: "Manipulation",
		},
		"Composite Bow": &Skill{
			Name:     "Composite Bow",
			Base:     5,
			Category: "Manipulation",
		},
		"Crossbows": &Skill{
			Name:     "Crossbows",
			Base:     25,
			Category: "Manipulation",
		},
		"Dagger, Throwing": &Skill{
			Name:     "Dagger, Throwing",
			Base:     5,
			Category: "Manipulation",
		},
		"Elf Bow": &Skill{
			Name:     "Elf Bow",
			Base:     5,
			Category: "Manipulation",
		},
		"Javelin": &Skill{
			Name:     "Javelin",
			Base:     10,
			Category: "Manipulation",
		},
		"Pole Lasso": &Skill{
			Name:     "Pole Lasso",
			Base:     5,
			Category: "Manipulation",
		},
		"Rock": &Skill{
			Name:     "Rock",
			Base:     15,
			Category: "Manipulation",
		},
		"Self Bow": &Skill{
			Name:     "Self Bow",
			Base:     5,
			Category: "Manipulation",
		},
		"Sling": &Skill{
			Name:     "Sling",
			Base:     5,
			Category: "Manipulation",
		},
		"Staff Sling": &Skill{
			Name:     "Staff Sling",
			Base:     10,
			Category: "Manipulation",
		},
		"Thrown Axe": &Skill{
			Name:     "Thrown Axe",
			Base:     10,
			Category: "Manipulation",
		},
		"Throwing Dagger": &Skill{
			Name:     "Throwing Dagger",
			Base:     10,
			Category: "Manipulation",
		},

		// Shields
		"Large Shield": &Skill{
			Name:     "Large Shield",
			Base:     15,
			Category: "Manipulation",
		},
		"Medium Shield": &Skill{
			Name:     "Medium Shield",
			Base:     15,
			Category: "Manipulation",
		},
		"Small Shield": &Skill{
			Name:     "Small Shield",
			Base:     15,
			Category: "Manipulation",
		},

		// Perception
		"Insight (species)": &Skill{
			Name:     "Insight (species)",
			Base:     20,
			Category: "Perception",
		},
		"Listen": &Skill{
			Name:     "Listen",
			Base:     25,
			Category: "Perception",
		},
		"Scan": &Skill{
			Name:     "Scan",
			Base:     25,
			Category: "Perception",
		},
		"Search": &Skill{
			Name:     "Search",
			Base:     25,
			Category: "Perception",
		},
		"Track": &Skill{
			Name:     "Track",
			Base:     5,
			Category: "Perception",
		},

		// Stealth
		"Hide": &Skill{
			Name:     "Hide",
			Base:     10,
			Category: "Stealth",
		},
		"Move Quietly": &Skill{
			Name:     "Move Quietly",
			Base:     10,
			Category: "Stealth",
		},
	}
	return &c
}
