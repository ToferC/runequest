package runequest

// PersonalHistoryEvents is a map of events for character history
var PersonalHistoryEvents = map[string]*Event{
	"1582_base": &Event{
		Name:        "Base",
		Year:        1583,
		Description: "Lots of text",
		Participant: "Grandparent",
		HomelandModifiers: map[string]int{
			"Sartar":  -5,
			"Esrolia": 5,
		},
		Slug: "1582_base",
		Results: []*EventResult{
			&EventResult{
				Range:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Description: "Go to war!",
				Skills: []Skill{
					Skill{
						Name: "Dodge",
						Updates: []*Update{
							&Update{
								Date:  "1583",
								Value: 10,
								Event: "1583 - go to war",
							},
						},
					},
				},
				Passions: []Ability{
					Ability{
						Name:       "Love (family)",
						CoreString: "Love",
						UserString: "family",
						Updates: []*Update{
							&Update{
								Date:  "1583",
								Value: 10,
								Event: "1583 - go to war",
							},
						},
					},
				},
				Lunars:     RollDice(10, 1, 0, 1) * 10,
				Reputation: RollDice(4, 1, 0, 1),
				ImmediateFollowEvent: map[string]int{
					"1583_civil_war": 0,
				},
				NextFollowEvent: map[string]int{
					"1584_starving": 0,
				},
			},
			&EventResult{
				Range:       []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				Description: "Don't go to war",
				Skills: []Skill{
					Skill{
						Name: "Bargain",
						Updates: []*Update{
							&Update{
								Date:  "1624",
								Value: 10,
								Event: "1624 - don't go to war",
							},
						},
					},
				},
				Passions: []Ability{
					Ability{
						Name:       "Hate (Lunars)",
						CoreString: "Hate",
						UserString: "Lunars",
						Updates: []*Update{
							&Update{
								Date:  "1583",
								Value: 10,
								Event: "1583 - go to war",
							},
						},
					},
				},
				Lunars:     RollDice(10, 1, 0, 1) * 10,
				Reputation: RollDice(4, 1, 0, 1),
				ImmediateFollowEvent: map[string]int{
					"1583_civil_war": 0,
				},
				NextFollowEvent: map[string]int{
					"1584_starving": 0,
				},
			},
		},
	},
}
