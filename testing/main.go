package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {
	c := runequest.NewCharacter("Bob")

	c.Description = "Man"

	c.Grandparent.Homeland = "Esrolia"
	c.Grandparent.Occupation = "Hunter"

	end := false
	event := runequest.PersonalHistoryEvents["1583_base"]
	mod := 0

	for !end {

		result := runequest.DetermineHistory(c, event, mod)

		last := c.History[len(c.History)-1]
		immediate := last.Results[0].ImmediateFollowEvent

		next := last.Results[0].NextFollowEvent

		if immediate != "" {
			mod := last.Results[0].ImmediateFollowMod
			e := runequest.PersonalHistoryEvents[immediate]
			r := runequest.DetermineHistory(c, e, mod)
		}

		e := runequest.PersonalHistoryEvents[follow]

		if c.Grandparent.Alive {
			r := runequest.DetermineHistory(c, e, m2)
		}
	}

	// Character done
	for k, v := range c.History {
		fmt.Println(k, v)
	}
}
