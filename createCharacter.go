package runequest

// NewCharacter generates a random starting character in Runequest
func NewCharacter(name string) *Character {
	c := Character{
		Name:           name,
		Role:           "Player Character",
		Statistics:     map[string]*Statistic{},
		Abilities:      map[string]*Ability{},
		PowerRunes:     map[string]*Ability{},
		ElementalRunes: map[string]*Ability{},
		ConditionRunes: map[string]*Ability{},
		RuneSpells:     map[string]*Spell{},
		SpiritMagic:    map[string]*Spell{},
		Homeland:       &Homeland{},
		Occupation:     &Occupation{},
		Cult:           &Cult{},
		Skills:         map[string]*Skill{},
		LocationForm:   "Humanoids",
		HitLocations:   LocationForms["Humanoids"],
		HitLocationMap: SortLocations(LocationForms["Humanoid"]),
		CreationSteps:  CreationStatus,

		MeleeAttacks:  map[string]*Attack{},
		RangedAttacks: map[string]*Attack{},

		SkillCategories: map[string]*SkillCategory{},

		Movement: []*Movement{},

		History: []*Event{},

		Grandparent: &FamilyMember{
			Alive: true,
		},
		Parent: &FamilyMember{
			Alive: true,
		},
	}

	// Copy base maps for new characters
	for k, v := range RuneQuestStats {
		c.Statistics[k] = &v
	}

	for k, v := range Skills {
		c.Skills[k] = &v
	}

	for k, v := range Abilities {
		c.Abilities[k] = &v
	}

	for k, v := range ElementalRunes {
		c.ElementalRunes[k] = &v
	}

	for k, v := range PowerRunes {
		c.PowerRunes[k] = &v
	}

	for k, v := range ConditionRunes {
		c.ConditionRunes[k] = &v
	}

	return &c
}
