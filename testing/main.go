package main

import (
	"fmt"
	"github.com/toferc/runequest"
)

func main() {

	weapons := runequest.BaseWeapons
	for _, w := range weapons {
		fmt.Println(w)
	}
}
