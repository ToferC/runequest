package runequest

// Homelands is a map of possible homelands in Runequest
var Homelands = map[string]*Homeland{
	// Sartar
	"Sartar": &Homeland{
		Name: "Sartar",
		Skills: []Skill{
			Skill{
				Name:       "Ride",
				CoreString: "Ride",
				UserChoice: true,
				UserString: "Horse",
				Base:       5,
				Value:      5,
				Category:   "Agility",
			},
			Skill{
				Name:  "Dance",
				Value: 5,
			},
			Skill{
				Name:  "Sing",
				Value: 10,
			},
			Skill{
				Name:     "Speak (Heortling)",
				Base:     50,
				Category: "Communication",
			},
			Skill{
				Name:     "Speak (Tradetalk)",
				Value:    10,
				Category: "Communication",
			},
			Skill{
				Name:     "Customs (Heortling)",
				Base:     25,
				Category: "Knowledge",
			},
			Skill{
				Name:  "Farm",
				Value: 20,
			},
			Skill{
				Name:  "Herd",
				Value: 10,
			},
			Skill{
				Name:  "Spirit Combat",
				Value: 15,
			},
			Skill{
				Name:  "Dagger",
				Value: 10,
			},
			Skill{
				Name:  "Broadsword",
				Value: 15,
			},
			Skill{
				Name:  "Battle Axe",
				Value: 10,
			},
			Skill{
				Name:  "1H Spear",
				Value: 10,
			},
			Skill{
				Name:  "Javelin",
				Value: 10,
			},
			Skill{
				Name:  "Medium Shield",
				Value: 15,
			},
			Skill{
				Name:  "Large Shield",
				Value: 10,
			},
		},
		// Skill Choices
		SkillChoices: []SkillChoice{
			// Choice of 2 skills
			SkillChoice{
				Skills: []Skill{
					// Skill 1
					Skill{
						Name:  "Composite Bow",
						Base:  5,
						Value: 10,
					},
					// Skill 2
					Skill{
						Name:  "Sling",
						Value: 10,
					},
				},
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
				Name:     "Occupation Lore (Local)",
				Value:    15,
				Category: "Knowledge",
			},
			Skill{
				Name:  "Farm",
				Value: 30,
			},
			Skill{
				Name:       "Craft",
				UserChoice: true,
				CoreString: "Craft",
				UserString: "Arms",
				Base:       10,
				Value:      15,
				Category:   "Manipulation",
			},
			Skill{
				Name:  "First Aid",
				Value: 10,
			},
			Skill{
				Name:  "Scan",
				Value: 10,
			},
			Skill{
				Name:  "Herd",
				Value: 15,
			},
			Skill{
				Name:  "Manage Household",
				Value: 30,
			},
			Skill{
				Name:  "Medium Shield",
				Value: 15,
			},
			Skill{
				Name:  "Broadsword",
				Value: 15,
			},
		},
		// Skill Choices
		SkillChoices: []SkillChoice{
			// Choice of 2 skills
			SkillChoice{
				Skills: []Skill{
					// Skill 1
					Skill{
						Name:  "Jump",
						Value: 10,
					},
					// Skill 2
					Skill{
						Name:  "Climb",
						Value: 10,
					},
				},
			},
		},
		// Passions
		Passions: []Ability{
			// Ability 1
			Ability{
				Name:       "Love (family)",
				CoreString: "Love",
				UserString: "family",
				UserChoice: true,
				Type:       "Passion",
				Base:       60,
				Value:      10,
			},
			// Ability 2
			Ability{
				Name:       "Loyalty (clan)",
				CoreString: "Loyalty",
				UserString: "clan",
				Type:       "Passion",
				UserChoice: true,
				Base:       60,
				Value:      10,
			},
			// Ability 3
			Ability{
				Name:       "Loyalty (tribe)",
				CoreString: "Loyalty",
				UserString: "tribe",
				Type:       "Passion",
				UserChoice: true,
				Base:       60,
				Value:      10,
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
				Name:     "Cult Lore (Orlanth)",
				Value:    15,
				Category: "Knowledge",
			},
			Skill{
				Name:     "Worship (Orlanth)",
				Value:    15,
				Category: "Magic",
			},
			Skill{
				Name:  "Meditate",
				Value: 25,
			},
			Skill{
				Name:  "Orate",
				Value: 30,
			},
			Skill{
				Name:  "Sing",
				Value: 10,
			},
			Skill{
				Name:     "Speak (Stormspeech)",
				Value:    10,
				Category: "Communication",
			},
		},
	},
	// Esrolia
}
