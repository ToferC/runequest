package runequest

// HitLocation represents a body area that can take damage
type HitLocation struct {
	Name     string
	HitLoc   []int
	Boxes    int
	Shock    []bool
	Kill     []bool
	Armor    int
	Disabled bool
}
