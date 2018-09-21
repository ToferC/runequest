package runequest

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// RollDice rolls and sum dice
func RollDice(max, min, bonus, numDice int) int {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	result := 0
	for i := 1; i < numDice+1; i++ {
		roll := r1.Intn(max+1-min) + min
		result += roll
	}

	result += bonus

	return result
}

// Sorting Functions

// ByValue implements the sort interface for abilities
type ByValue []*Ability

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value > a[j].Value }

// AddRuneModifiers adds stat modifiers based on runes
func (c *Character) AddRuneModifiers() {

	var runes []*Ability

	for _, a := range c.Abilities {
		if a.Type == "Elemental Rune" {
			runes = append(runes, a)
		}
	}

	// Sort Runes
	sort.Sort(ByValue(runes))
	fmt.Println(runes)

	primary, secondary := runes[0].Name, runes[1].Name

	switch {
	case primary == "Air":
		c.Statistics["STR"].Value += 2
	case primary == "Earth":
		c.Statistics["CON"].Value += 2
	case primary == "Darkness":
		c.Statistics["SIZ"].Value += 2
	case primary == "Fire/Sky":
		c.Statistics["INT"].Value += 2
	case primary == "Water":
		c.Statistics["DEX"].Value += 2
	case primary == "Moon":
		c.Statistics["POW"].Value += 2
	}

	switch {
	case secondary == "Air":
		c.Statistics["STR"].Value++
	case secondary == "Earth":
		c.Statistics["CON"].Value++
	case secondary == "Darkness":
		c.Statistics["SIZ"].Value++
	case secondary == "Fire/Sky":
		c.Statistics["INT"].Value++
	case secondary == "Water":
		c.Statistics["DEX"].Value++
	case secondary == "Moon":
		c.Statistics["POW"].Value++
	}

}

// ChooseRandom chooses a random element in a slice
// when given l as len(slice)
func ChooseRandom(l int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	return r.Intn(l)
}
