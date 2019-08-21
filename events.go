package runequest

// FamilyMember tracks family history through previous character history
type FamilyMember struct {
	Name        string
	Occupation  string
	Relation    string
	Born        int
	Died        int
	Surviving   bool
	Description string
}

// Event represents a full event in previous character history
type Event struct {
	Year        int
	Name        string
	Description string
	Participant string
	// Character, Parent, Grandparent
	HomelandModifiers map[string]int
	// Map Homeland name to modifier on d20 roll
	OccupationModifiers map[string]int
	// Map Occupation name to modifier on d20 roll
	Results []*EventResult
	Slug    string
}

// EventResult is a specific die range of random results from previous
// Character history
type EventResult struct {
	Range                []int
	Description          string
	Skills               []Skill
	Passions             []Ability
	Lunars               int
	Reputation           int
	Equipment            string
	Lethal               bool
	Boon                 bool
	RandomDeath          bool
	ImmediateFollowEvent map[string]int // Slug
	NextFollowEvent      map[string]int // Slug
}

// Boon represents a random benefit from an EventResult
type Boon struct {
	Range       []int
	Description string
	Skills      []Skill
	Passsions   []Ability
	Lunars      int
	Reputation  int
	Equipment   string
}

// RandomCauseOfDeath is a random extinction event for family
type RandomCauseOfDeath struct {
	Range       []int
	Description string
	Skills      []Skill
	Passsions   []Ability
	Lunars      int
	Reputation  int
	Equipment   string
	Lethal      bool
}
