package runequest

import "fmt"

// SkillCategory is a grouping of skills
type SkillCategory struct {
	Name          string
	StatModifiers []statMods
	Value         int
}

func (sc SkillCategory) String() string {

	var text string

	if sc.Value > -1 {
		text += fmt.Sprintf("\n\n**%s +%d%%", sc.Name, sc.Value)
	} else {
		text += fmt.Sprintf("\n\n**%s %d%%", sc.Name, sc.Value)
	}
	return text
}

// CategoryOrder sets the order of skills in Runequest
var CategoryOrder = []string{"Agility", "Communication", "Knowledge", "Magic", "Manipulation", "Perception", "Stealth"}

type statMods struct {
	statistic string
	values    map[int]int
}

// DetermineSkillCategoryValues figures out base values for skill categories based on stats
func (c *Character) DetermineSkillCategoryValues() {
	for _, s := range c.SkillCategories {
		// For each category

		for _, sm := range s.StatModifiers {
			// For each modifier in a category

			stat := c.Statistics[sm.statistic]
			// Identify the stat

			// Map against specific values
			switch {
			case stat.Total <= 4:
				s.Value += sm.values[4]
			case stat.Total <= 8:
				s.Value += sm.values[8]
			case stat.Total <= 12:
				s.Value += sm.values[12]
			case stat.Total <= 16:
				s.Value += sm.values[16]
			case stat.Total <= 20:
				s.Value += sm.values[20]
			case stat.Total > 20:
				if sm.values[20] > 0 {
					s.Value += sm.values[20] + ((stat.Value-20)/4)*5
				} else {
					s.Value += sm.values[20] - ((stat.Value-20)/4)*5
				}
			}
		}
	}

	for _, skill := range c.Skills {
		sc := SkillCategories[skill.Category]

		skill.CategoryValue = sc.Value
	}

}

// SkillCategories is a map of skill categories
var SkillCategories = map[string]*SkillCategory{
	// Agility
	"Agility": &SkillCategory{
		Name: "Agility",
		StatModifiers: []statMods{
			statMods{
				statistic: "STR",
				values: map[int]int{
					4:  -5,
					8:  0,
					12: 0,
					16: 0,
					20: 5,
				},
			},
			statMods{
				statistic: "SIZ",
				values: map[int]int{
					4:  5,
					8:  0,
					12: 0,
					16: 0,
					20: -5,
				},
			},
			statMods{
				statistic: "DEX",
				values: map[int]int{
					4:  -10,
					8:  -5,
					12: 0,
					16: 5,
					20: 10,
				},
			},
			statMods{
				statistic: "POW",
				values: map[int]int{
					4:  -5,
					8:  0,
					12: 0,
					16: 0,
					20: 5,
				},
			},
		},
	},
	// Communication
	"Communication": &SkillCategory{
		Name: "Communication",
		StatModifiers: []statMods{
			statMods{
				statistic: "INT",
				values: map[int]int{
					4:  -5,
					8:  0,
					12: 0,
					16: 0,
					20: 5,
				},
			},
			statMods{
				statistic: "POW",
				values: map[int]int{
					4:  -5,
					8:  0,
					12: 0,
					16: 0,
					20: 5,
				},
			},
			statMods{
				statistic: "CHA",
				values: map[int]int{
					4:  -10,
					8:  -5,
					12: 0,
					16: 5,
					20: 10,
				},
			},
		},
	},
	// Knowledge
	"Knowledge": &SkillCategory{
		Name: "Knowledge",
		StatModifiers: []statMods{
			statMods{
				statistic: "INT",
				values: map[int]int{
					4:  -10,
					8:  -5,
					12: 0,
					16: 5,
					20: 10,
				},
			},
			statMods{
				statistic: "POW",
				values: map[int]int{
					4:  -5,
					8:  0,
					12: 0,
					16: 0,
					20: 5,
				},
			},
		},
	},
	// Magic
	"Magic": &SkillCategory{
		Name: "Magic",
		StatModifiers: []statMods{
			statMods{
				statistic: "POW",
				values: map[int]int{
					4:  -10,
					8:  -5,
					12: 0,
					16: 5,
					20: 10,
				},
			},
			statMods{
				statistic: "CHA",
				values: map[int]int{
					4:  -5,
					8:  0,
					12: 0,
					16: 0,
					20: 5,
				},
			},
		},
	},
	// Manipulation
	"Manipulation": &SkillCategory{
		Name: "Manipulation",
		StatModifiers: []statMods{
			statMods{
				statistic: "STR",
				values: map[int]int{
					4:  -5,
					8:  0,
					12: 0,
					16: 0,
					20: 5,
				},
			},
			statMods{
				statistic: "DEX",
				values: map[int]int{
					4:  -10,
					8:  -5,
					12: 0,
					16: 5,
					20: 10,
				},
			},
			statMods{
				statistic: "INT",
				values: map[int]int{
					4:  -10,
					8:  -5,
					12: 0,
					16: 5,
					20: 10,
				},
			},
			statMods{
				statistic: "POW",
				values: map[int]int{
					4:  -5,
					8:  0,
					12: 0,
					16: 0,
					20: 5,
				},
			},
		},
	},
	// Perception
	"Perception": &SkillCategory{
		Name: "Perception",
		StatModifiers: []statMods{
			statMods{
				statistic: "INT",
				values: map[int]int{
					4:  -10,
					8:  -5,
					12: 0,
					16: 5,
					20: 10,
				},
			},
			statMods{
				statistic: "POW",
				values: map[int]int{
					4:  -5,
					8:  0,
					12: 0,
					16: 0,
					20: 5,
				},
			},
		},
	},
	// Stealth
	"Stealth": &SkillCategory{
		Name: "Stealth",
		StatModifiers: []statMods{
			statMods{
				statistic: "SIZ",
				values: map[int]int{
					4:  10,
					8:  5,
					12: 0,
					16: -5,
					20: -10,
				},
			},
			statMods{
				statistic: "DEX",
				values: map[int]int{
					4:  -10,
					8:  -5,
					12: 0,
					16: 5,
					20: 10,
				},
			},
			statMods{
				statistic: "INT",
				values: map[int]int{
					4:  -10,
					8:  -5,
					12: 0,
					16: 5,
					20: 10,
				},
			},
			statMods{
				statistic: "POW",
				values: map[int]int{
					4:  5,
					8:  0,
					12: 0,
					16: 0,
					20: -5,
				},
			},
		},
	},
}
