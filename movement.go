package runequest

// Movement represents a movement type and value
type Movement struct {
	Name  string
	Value int
}

def (m Movement) String() {
	fmt.Sprintf("%s: %d", m.Name, m.Value)
}
