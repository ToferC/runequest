package runequest

// Homelands is a map of possible homelands in Runequest
var Homelands = map[string]*Homeland{
	// Sartar
	"Sartar": &Homeland{
		Name: "Sartar",
		Skills: []Skill{
			Skill{
				Name:          "Ride",
				CoreString:    "Ride",
				UserChoice:    true,
				UserString:    "Horse",
				Base:          5,
				HomelandValue: 5,
				Category:      "Agility",
			},
			Skill{
				Name:          "Dance",
				HomelandValue: 5,
			},
			Skill{
				Name:          "Sing",
				HomelandValue: 10,
			},
			Skill{
				Name:     "Speak (Heortling)",
				Base:     50,
				Category: "Communication",
			},
			Skill{
				Name:          "Speak (Tradetalk)",
				HomelandValue: 10,
				Category:      "Communication",
			},
			Skill{
				Name:     "Customs (Heortling)",
				Base:     25,
				Category: "Knowledge",
			},
			Skill{
				Name:          "Farm",
				HomelandValue: 20,
			},
			Skill{
				Name:          "Herd",
				HomelandValue: 10,
			},
			Skill{
				Name:          "Spirit Combat",
				HomelandValue: 15,
			},
			Skill{
				Name:          "Dagger",
				HomelandValue: 10,
			},
			Skill{
				Name:          "Broadsword",
				HomelandValue: 15,
			},
			Skill{
				Name:          "Battle Axe",
				HomelandValue: 10,
			},
			Skill{
				Name:          "1H Spear",
				HomelandValue: 10,
			},
			Skill{
				Name:          "Javelin",
				HomelandValue: 10,
			},
			Skill{
				Name:          "Medium Shield",
				HomelandValue: 15,
			},
			Skill{
				Name:          "Large Shield",
				HomelandValue: 10,
			},
		},
		// Skill Choices
		SkillChoices: []SkillChoice{
			// Choice of 2 skills
			SkillChoice{
				Skills: []Skill{
					// Skill 1
					Skill{
						Name:          "Composite Bow",
						Base:          5,
						HomelandValue: 10,
					},
					// Skill 2
					Skill{
						Name:          "Sling",
						HomelandValue: 10,
					},
				},
			},
		},
		// Rune Affinities
		AbilityList: []Ability{
			// Ability 1
			Ability{
				Name:          "Air",
				CoreString:    "Air",
				Type:          "Elemental Rune",
				HomelandValue: 10,
			},
		},
		// Passions
		PassionList: []Ability{
			// Ability 1
			Ability{
				Name:          "Loyalty (Sartar)",
				CoreString:    "Loyalty",
				UserString:    "Sartar",
				UserChoice:    true,
				Type:          "Passion",
				Base:          60,
				HomelandValue: 10,
			},
			// Ability 2
			Ability{
				Name:          "Loyalty (clan)",
				CoreString:    "Loyalty",
				UserString:    "clan",
				Type:          "Passion",
				UserChoice:    true,
				Base:          60,
				HomelandValue: 10,
			},
			// Ability 3
			Ability{
				Name:          "Loyalty (tribe)",
				CoreString:    "Loyalty",
				UserString:    "tribe",
				Type:          "Passion",
				UserChoice:    true,
				Base:          60,
				HomelandValue: 10,
			},
		},
	},
	// Esrolia
}

// Occupations is a map of possible Occupations in Runequest
var Occupations = map[string]*Occupation{
	"Farmer": &Occupation{
		Name: "Farmer",
		Skills: []Skill{
			Skill{
				Name:            "Occupation Lore (Local)",
				OccupationValue: 15,
				Category:        "Knowledge",
			},
			Skill{
				Name:            "Farm",
				OccupationValue: 30,
			},
			Skill{
				Name:            "Craft",
				UserChoice:      true,
				CoreString:      "Craft",
				UserString:      "Arms",
				Base:            10,
				OccupationValue: 15,
				Category:        "Manipulation",
			},
			Skill{
				Name:            "First Aid",
				OccupationValue: 10,
			},
			Skill{
				Name:            "Scan",
				OccupationValue: 10,
			},
			Skill{
				Name:            "Herd",
				OccupationValue: 15,
			},
			Skill{
				Name:            "Manage Household",
				OccupationValue: 30,
			},
			Skill{
				Name:            "Medium Shield",
				OccupationValue: 15,
			},
			Skill{
				Name:            "Broadsword",
				OccupationValue: 15,
			},
		},
		// Skill Choices
		SkillChoices: []SkillChoice{
			// Choice of 2 skills
			SkillChoice{
				Skills: []Skill{
					// Skill 1
					Skill{
						Name:            "Jump",
						OccupationValue: 10,
					},
					// Skill 2
					Skill{
						Name:            "Climb",
						OccupationValue: 10,
					},
				},
			},
		},
		// Passions
		PassionList: []Ability{
			// Ability 1
			Ability{
				Name:            "Love (family)",
				CoreString:      "Love",
				UserString:      "family",
				UserChoice:      true,
				Type:            "Passion",
				Base:            60,
				OccupationValue: 10,
			},
			// Ability 2
			Ability{
				Name:            "Loyalty (clan)",
				CoreString:      "Loyalty",
				UserString:      "clan",
				Type:            "Passion",
				UserChoice:      true,
				Base:            60,
				OccupationValue: 10,
			},
			// Ability 3
			Ability{
				Name:            "Loyalty (tribe)",
				CoreString:      "Loyalty",
				UserString:      "tribe",
				Type:            "Passion",
				UserChoice:      true,
				Base:            60,
				OccupationValue: 10,
			},
		},
	},
	// Esrolia
}

// Cults is a map of possible Cults in Runequest
var Cults = map[string]*Cult{
	"Orlanth": &Cult{
		Name: "Orlanth",
		Skills: []Skill{
			Skill{
				Name:      "Cult Lore (Orlanth)",
				CultValue: 15,
				Category:  "Knowledge",
			},
			Skill{
				Name:      "Worship (Orlanth)",
				CultValue: 15,
				Category:  "Magic",
			},
			Skill{
				Name:      "Meditate",
				CultValue: 25,
			},
			Skill{
				Name:      "Orate",
				CultValue: 30,
			},
			Skill{
				Name:      "Sing",
				CultValue: 10,
			},
			Skill{
				Name:      "Speak (Stormspeech)",
				CultValue: 10,
				Category:  "Communication",
			},
		},
		// Passions
		PassionList: []Ability{
			// Ability 1
			Ability{
				Name:       "Devotion (Orlanth)",
				CoreString: "Devotion",
				UserString: "Orlanth",
				UserChoice: true,
				Type:       "Passion",
				Base:       60,
				CultValue:  10,
			},
			// Ability 2
			Ability{
				Name:       "Hate (Chaos)",
				CoreString: "Hate",
				UserString: "Chaos",
				Type:       "Passion",
				UserChoice: true,
				Base:       60,
				CultValue:  10,
			},
			// Ability 3
			Ability{
				Name:       "Loyalty (temple)",
				CoreString: "Loyalty",
				UserString: "temple",
				Type:       "Passion",
				UserChoice: true,
				Base:       60,
				CultValue:  10,
			},
			// Ability 4
			Ability{
				Name:       "Honor",
				CoreString: "Honor",
				Type:       "Passion",
				Base:       60,
				CultValue:  10,
			},
		},
	},
	// Esrolia
}
