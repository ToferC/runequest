package main

import (
	"fmt"

	"github.com/toferc/runequest"
)

func main() {

	c := runequest.NewCharacter("Default")
	hl := runequest.SortLocations(c.HitLocations)
	fmt.Println(hl)

}
