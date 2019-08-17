package runequest

// Event represents a full event in previous character history
type Event struct {
	Year        int
	Name        string
	Description string
	Participant string
	// Character, Parent, Grandparent
	HomelandsModifiers map[string]int
	// Map Homeland name to modifier on d20 roll
	ImmediateFollowEvent *Event
	NextYearEvent        *Event
	Results              []*EventResult
	Slug                 string
}

// EventResult is a specific die range of random results from previous
// Character history
type EventResult struct {
	Range       []int
	Description string
	Updates     []Update
	Lunars      int
	Reputation  int
	Equipment   string
	Lethal      bool
}

// Boon represents a random benefit from an EventResult
type Boon struct {
	Range       []int
	Description string
	Updates     []Update
	Lunars      int
	Reputation  int
	Equipment   string
}

// RandomCauseOfDeath is a random extinction event for family
type RandomCauseOfDeath struct {
	Range       []int
	Description string
	Updates     []Update
	Lunars      int
	Reputation  int
	Equipment   string
	Lethal      bool
}
