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

	end := false // tracks end of player history
	mod := 0
	next := "1583_base"

	for !end {

		event := runequest.PersonalHistoryEvents[next]

		// Start history
		result, end := runequest.DetermineHistory(c, event, mod)

		// Identify last EventResult
		if len(c.History) > 0 {
			last := c.History[len(c.History)-1]
			next := last.Results[0].NextFollowEvent
			immediate := last.Results[0].ImmediateFollowEvent
			// Check for immediate followup

			if immediate != "" {

				for immediate != "" {
					// if immediate follow-up, go there
					e := runequest.PersonalHistoryEvents[immediate]
					mod := last.Results[0].ImmediateFollowMod
					r, end := runequest.DetermineHistory(c, e, mod)

					// Identify last EventResult
					last := c.History[len(c.History)-1]
					next := last.Results[0].NextFollowEvent

					// Check for immediate followup
					immediate := last.Results[0].ImmediateFollowEvent
					fmt.Println(next)
					fmt.Println(end, immediate, r)
				}
			}
			fmt.Println(next)
		}

		// if no immediate event or after immediate event
		fmt.Println(result, end, next)
	}

	// Character done
	for k, v := range c.History {
		fmt.Println(k, v)
	}
}
