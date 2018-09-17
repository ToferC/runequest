package runequest

// HitLocation represents a body area that can take damage
type HitLocation struct {
	Name     string
	HitLoc   []int
	Max      int
	Value    int
	Armor    int
	Disabled bool
}
