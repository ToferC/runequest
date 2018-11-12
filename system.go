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

// ByTotal implements the sort interface for abilities
type ByTotal []*Ability

func (a ByTotal) Len() int           { return len(a) }
func (a ByTotal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTotal) Less(i, j int) bool { return a[i].Total > a[j].Total }

// DetermineRuneModifiers adds stat modifiers based on runes
func (c *Character) DetermineRuneModifiers() []string {

	// Set empty array for sorting
	var runes []*Ability

	// Add abilities to array
	for _, a := range c.ElementalRunes {
		a.UpdateAbility()
		runes = append(runes, a)
	}

	// Sort Runes
	sort.Sort(ByTotal(runes))
	fmt.Println(runes)

	primary, secondary := runes[0].Name, runes[1].Name

	return []string{primary, secondary}

}
