package runequest

// Cult represents a Religion in Runequest
type Cult struct {
	Name        string
	Description string
	Modifiers   []*Modifier
	SpellList   []*Spell
}
