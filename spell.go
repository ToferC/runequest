package runequest

// Spell is a magical effect
type Spell struct {
	Name        string
	Description string
	Range       int
	Duration    string
	Cost        int
	Source      *DerivedStat
}
