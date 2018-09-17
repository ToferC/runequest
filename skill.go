package runequest

import "fmt"

// Skill is a learned ability of an RPG Character
type Skill struct {
	Name            string
	category        string
	TakesSubject    bool
	Subject         string
	Base            int
	CategoryValue   int
	Value           int
	Total           int
	Min             int
	Max             int
	ExperienceCheck bool
}

func (s *Skill) String() string {

	s.Total = s.Base + s.Value + s.CategoryValue

	text := ""

	if s.TakesSubject {
		text += fmt.Sprintf("%s (%s) %d%%", s.Name, s.Subject, s.Total)
	} else {
		text += fmt.Sprintf("%s %d%%", s.Name, s.Total)
	}

	return text
}
