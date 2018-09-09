package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {
	c := runequest.Character{
		Name:       "Bigby",
		Statistics: runequest.RuneQuestStats,
	}

	fmt.Println(c)
}
