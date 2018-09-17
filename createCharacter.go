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
			category: "Agility",
		},
		"Climb": &Skill{
			Name:     "Climb",
			Base:     40,
			category: "Agility",
		},
		"Dodge": &Skill{
			Name:     "Dodge",
			Base:     c.Statistics["DEX"].Value * 2,
			category: "Agility",
		},
		"Drive Chariot": &Skill{
			Name:     "Drive Chariot",
			Base:     5,
			category: "Agility",
		},
		"Jump": &Skill{
			Name:     "Jump",
			Base:     c.Statistics["DEX"].Value * 3,
			category: "Agility",
		},
		"Ride": &Skill{
			Name:         "Ride",
			TakesSubject: true,
			Subject:      "Horse",
			Base:         5,
			category:     "Agility",
		},
		"Swim": &Skill{
			Name:     "Swim",
			Base:     15,
			category: "Agility",
		},

		// Communication
		"Act": &Skill{
			Name:     "Act",
			Base:     5,
			category: "Communication",
		},
		"Art": &Skill{
			Name:     "Art",
			Base:     5,
			category: "Communication",
		},
		"Bargain": &Skill{
			Name:     "Bargain",
			Base:     5,
			category: "Communication",
		},
		"Charm": &Skill{
			Name:     "Charm",
			Base:     15,
			category: "Communication",
		},
		"Dance": &Skill{
			Name:     "Dance",
			Base:     10,
			category: "Communication",
		},
		"Disguise": &Skill{
			Name:     "Disguise",
			Base:     5,
			category: "Communication",
		},
		"Fast Talk": &Skill{
			Name:     "Fast Talk",
			Base:     5,
			category: "Communication",
		},
		"Intimidate": &Skill{
			Name:     "Intimidate",
			Base:     15,
			category: "Communication",
		},
		"Intrigue": &Skill{
			Name:     "Intrigue",
			Base:     5,
			category: "Communication",
		},
		"Orate": &Skill{
			Name:     "Orate",
			Base:     10,
			category: "Communication",
		},
		"Sing": &Skill{
			Name:     "Sing",
			Base:     10,
			category: "Communication",
		},
		"Speak Own Language": &Skill{
			Name:     "Speak Own Language",
			Base:     50,
			category: "Communication",
		},
		"Speak Other Language": &Skill{
			Name:         "Speak Other Language",
			TakesSubject: true,
			Subject:      "Praxian",
			Base:         0,
			category:     "Communication",
		},
		// Knowledge
		"Alchemy": &Skill{
			Name:     "Alchemy",
			Base:     0,
			category: "Knowledge",
		},
		"Animal Lore": &Skill{
			Name:     "Animal Lore",
			Base:     5,
			category: "Knowledge",
		},
		"Battle": &Skill{
			Name:     "Battle",
			Base:     10,
			category: "Knowledge",
		},
		"Bureacracy": &Skill{
			Name:     "Bureacracy",
			Base:     0,
			category: "Knowledge",
		},
		"Celestial Lore": &Skill{
			Name:     "Celestial Lore",
			Base:     5,
			category: "Knowledge",
		},
		"Cult Lore": &Skill{
			Name:         "Cult Lore",
			TakesSubject: true,
			Subject:      "Orlanth",
			Base:         0,
			category:     "Knowledge",
		},
		"Customs (Local)": &Skill{
			Name:     "Customs (Local)",
			Base:     25,
			category: "Knowledge",
		},
		"Customs": &Skill{
			Name:         "Customs",
			TakesSubject: true,
			Subject:      "Praxian",
			Base:         0,
			category:     "Knowledge",
		},
		"Elder Race Lore": &Skill{
			Name:         "Elder Race Lore",
			TakesSubject: true,
			Subject:      "Elves",
			Base:         5,
			category:     "Knowledge",
		},
		"Evaluate": &Skill{
			Name:     "Evaluate",
			Base:     10,
			category: "Knowledge",
		},
		"Farm": &Skill{
			Name:     "Farm",
			Base:     10,
			category: "Knowledge",
		},
		"First Aid": &Skill{
			Name:     "First Aid",
			Base:     10,
			category: "Knowledge",
		},
		"Game": &Skill{
			Name:     "Game",
			Base:     15,
			category: "Knowledge",
		},
		"Herd": &Skill{
			Name:     "Herd",
			Base:     5,
			category: "Knowledge",
		},
		"Homeland Lore (local)": &Skill{
			Name:     "Homeland Lore (local)",
			Base:     30,
			category: "Knowledge",
		},
		"Homeland Lore": &Skill{
			Name:         "Homeland Lore",
			TakesSubject: true,
			Subject:      "Spirits",
			Base:         0,
			category:     "Knowledge",
		},
		"Library Use": &Skill{
			Name:     "Library Use",
			Base:     0,
			category: "Knowledge",
		},
		"Manage Household": &Skill{
			Name:     "Manage Household",
			Base:     10,
			category: "Knowledge",
		},
		"Mineral Lore": &Skill{
			Name:     "Mineral Lore",
			Base:     5,
			category: "Knowledge",
		},
		"Peaceful Cut": &Skill{
			Name:     "Peaceful Cut",
			Base:     10,
			category: "Knowledge",
		},
		"Plant Lore": &Skill{
			Name:     "Plant Lore",
			Base:     5,
			category: "Knowledge",
		},
		"Read/Write": &Skill{
			Name:         "Read/Write",
			TakesSubject: true,
			Subject:      "Old Tarsh",
			Base:         0,
			category:     "Knowledge",
		},
		"Shiphandling": &Skill{
			Name:     "Shiphandling",
			Base:     0,
			category: "Knowledge",
		},
		"Survival": &Skill{
			Name:     "Survival",
			Base:     15,
			category: "Knowledge",
		},
		"Treat Disease": &Skill{
			Name:     "Treat Disease",
			Base:     5,
			category: "Knowledge",
		},
		"Treat Poison": &Skill{
			Name:     "Treat Poison",
			Base:     5,
			category: "Knowledge",
		},
		// Magic
		"Meditate": &Skill{
			Name:     "Meditate",
			Base:     0,
			category: "Magic",
		},
		"Prepare Corpse": &Skill{
			Name:     "Prepare Corpse",
			Base:     10,
			category: "Magic",
		},
		"Sense Assassin": &Skill{
			Name:     "Sense Assassin",
			Base:     0,
			category: "Magic",
		},
		"Sense Chaos": &Skill{
			Name:     "Sense Chaos",
			Base:     0,
			category: "Magic",
		},
		"Spirit Combat": &Skill{
			Name:     "Spirit Combat",
			Base:     20,
			category: "Magic",
		},
		"Spirit Dance": &Skill{
			Name:     "Spirit Dance",
			Base:     0,
			category: "Magic",
		},
		"Spirit Lore": &Skill{
			Name:     "Spirit Lore",
			Base:     0,
			category: "Magic",
		},
		"Spirit Travel": &Skill{
			Name:     "Spirit Travel",
			Base:     0,
			category: "Magic",
		},
		"Understand Herd Beast": &Skill{
			Name:     "Understand Herd Beast",
			Base:     0,
			category: "Magic",
		},
		"Worship": &Skill{
			Name:         "Meditate",
			TakesSubject: true,
			Subject:      "Orlanth",
			Base:         0,
			category:     "Magic",
		},

		// Manipulation
		"Conceal": &Skill{
			Name:     "Conceal",
			Base:     5,
			category: "Manipulation",
		},
		"Craft": &Skill{
			Name:         "Craft",
			TakesSubject: true,
			Subject:      "Armory",
			Base:         10,
			category:     "Manipulation",
		},
		"Devise": &Skill{
			Name:     "Devise",
			Base:     5,
			category: "Manipulation",
		},
		"Play Instrument": &Skill{
			Name:     "Play Instrument",
			Base:     5,
			category: "Manipulation",
		},
		"Sleight": &Skill{
			Name:     "Sleight",
			Base:     10,
			category: "Manipulation",
		},

		// Melee Weapons
		"1H Axe": &Skill{
			Name:     "1H Axe",
			Base:     10,
			category: "Manipulation",
		},
		"2H Axe": &Skill{
			Name:     "2H Axe",
			Base:     5,
			category: "Manipulation",
		},
		"Broadsword": &Skill{
			Name:     "Broadsword",
			Base:     10,
			category: "Manipulation",
		},
		"Dagger": &Skill{
			Name:     "Dagger",
			Base:     15,
			category: "Manipulation",
		},
		"Fist": &Skill{
			Name:     "Fist",
			Base:     25,
			category: "Manipulation",
		},
		"Grapple": &Skill{
			Name:     "Grapple",
			Base:     25,
			category: "Manipulation",
		},
		"1H Hammer": &Skill{
			Name:     "1H Hammer",
			Base:     10,
			category: "Manipulation",
		},
		"2H Hammer": &Skill{
			Name:     "2H Hammer",
			Base:     5,
			category: "Manipulation",
		},
		"Kick": &Skill{
			Name:     "Kick",
			Base:     15,
			category: "Manipulation",
		},
		"Kopis": &Skill{
			Name:     "Kopis",
			Base:     10,
			category: "Manipulation",
		},
		"1H Mace": &Skill{
			Name:     "1H Mace",
			Base:     15,
			category: "Manipulation",
		},
		"2H Mace": &Skill{
			Name:     "2H Mace",
			Base:     10,
			category: "Manipulation",
		},
		"Pike": &Skill{
			Name:     "Pike",
			Base:     15,
			category: "Manipulation",
		},
		"Quarterstaff": &Skill{
			Name:     "Quarterstaff",
			Base:     15,
			category: "Manipulation",
		},
		"Rapier": &Skill{
			Name:     "Rapier",
			Base:     5,
			category: "Manipulation",
		},
		"Shortsword": &Skill{
			Name:     "Shortsword",
			Base:     10,
			category: "Manipulation",
		},
		"1H Spear": &Skill{
			Name:     "1H Spear",
			Base:     05,
			category: "Manipulation",
		},
		"2H Spear": &Skill{
			Name:     "2H Spear",
			Base:     15,
			category: "Manipulation",
		},

		// Missile Weapons
		"Arbalest": &Skill{
			Name:     "Arbalest",
			Base:     10,
			category: "Manipulation",
		},
		"Axe, Throwing": &Skill{
			Name:     "Axe, Throwing",
			Base:     10,
			category: "Manipulation",
		},
		"Composite Bow": &Skill{
			Name:     "Composite Bow",
			Base:     5,
			category: "Manipulation",
		},
		"Crossbows": &Skill{
			Name:     "Crossbows",
			Base:     25,
			category: "Manipulation",
		},
		"Dagger, Throwing": &Skill{
			Name:     "Dagger, Throwing",
			Base:     5,
			category: "Manipulation",
		},
		"Elf Bow": &Skill{
			Name:     "Elf Bow",
			Base:     5,
			category: "Manipulation",
		},
		"Javelin": &Skill{
			Name:     "Javelin",
			Base:     10,
			category: "Manipulation",
		},
		"Pole Lasso": &Skill{
			Name:     "Pole Lasso",
			Base:     5,
			category: "Manipulation",
		},
		"Rock": &Skill{
			Name:     "Rock",
			Base:     15,
			category: "Manipulation",
		},
		"Self Bow": &Skill{
			Name:     "Self Bow",
			Base:     5,
			category: "Manipulation",
		},
		"Sling": &Skill{
			Name:     "Sling",
			Base:     5,
			category: "Manipulation",
		},
		"Staff Sling": &Skill{
			Name:     "Staff Sling",
			Base:     10,
			category: "Manipulation",
		},
		"Thrown Axe": &Skill{
			Name:     "Thrown Axe",
			Base:     10,
			category: "Manipulation",
		},
		"Throwing Dagger": &Skill{
			Name:     "Throwing Dagger",
			Base:     10,
			category: "Manipulation",
		},

		// Shields
		"Large Shield": &Skill{
			Name:     "Large Shield",
			Base:     15,
			category: "Manipulation",
		},
		"Medium Shield": &Skill{
			Name:     "Medium Shield",
			Base:     15,
			category: "Manipulation",
		},
		"Small Shield": &Skill{
			Name:     "Small Shield",
			Base:     15,
			category: "Manipulation",
		},

		// Perception
		"Insight": &Skill{
			Name:         "Insight",
			TakesSubject: true,
			Subject:      "Troll",
			Base:         00,
			category:     "Perception",
		},
		"Insight (species)": &Skill{
			Name:     "Insight (species)",
			Base:     20,
			category: "Perception",
		},
		"Listen": &Skill{
			Name:     "Listen",
			Base:     25,
			category: "Perception",
		},
		"Scan": &Skill{
			Name:     "Scan",
			Base:     25,
			category: "Perception",
		},
		"Search": &Skill{
			Name:     "Search",
			Base:     25,
			category: "Perception",
		},
		"Track": &Skill{
			Name:     "Track",
			Base:     5,
			category: "Perception",
		},

		// Stealth
		"Hide": &Skill{
			Name:     "Hide",
			Base:     10,
			category: "Stealth",
		},
		"Move Quietly": &Skill{
			Name:     "Move Quietly",
			Base:     10,
			category: "Stealth",
		},
	}
	return &c
}
