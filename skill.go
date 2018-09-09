package runequest

import "fmt"

// Skill is a learned ability of an RPG Character
type Skill struct {
	Name     string
	Base     int
	category SkillCategory
	Value    int
}

func (s *Skill) String() string {
	text := fmt.Sprintf("%s %d", s.Name, s.Value)
	return text
}

// SkillCategory is a grouping of skills
type SkillCategory struct {
	Name          string
	StatModifiers []*statMods
}

type statMods struct {
	statistic Statistic
	value     map[int]int
}
