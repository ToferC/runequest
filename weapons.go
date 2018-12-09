package runequest

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

// Weapon represents a Runequest weapon
type Weapon struct {
	Name      string
	Skill     *Skill
	STR       int
	DEX       int
	Damage    string
	STRDamage bool
	HP        int
	CurrentHP int
	ENC       int
	Length    float64
	SR        int
	Type      string
	Range     int
}

// BaseWeapons is an array of runequest weapons
var BaseWeapons = loadWeapons()

func translateDieCode(s string) *DieCode {
	// translates a string like 1d6+1 into a DieCode

	var dice, mod, max int
	dieCode := &DieCode{}

	str := strings.Split(s, "+")
	if str[1] != "" {
		mod, _ = strconv.Atoi(str[1])
	}

	diceString := strings.Split(str[0], "D")
	dice, _ = strconv.Atoi(diceString[0])
	max, _ = strconv.Atoi(diceString[1])

	dieCode.NumDice = dice
	dieCode.DiceMax = max
	dieCode.Modifier = mod

	return dieCode
}

func loadWeapons() []*Weapon {

	weapons := []*Weapon{}

	f, _ := os.Open("weapons.csv")

	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		sHP, _ := strconv.Atoi(record[5])

		if record[2] == "melee" {
			sSR, err := strconv.Atoi(record[2])
			if err != nil {
				sSR = 0
			}

			weapons = append(weapons, &Weapon{
				Name:      record[0],
				Type:      "melee",
				SR:        sSR,
				STRDamage: true,
				Damage:    record[3],
				HP:        sHP,
				CurrentHP: sHP,
			})
		} else {

			r, _ := strconv.Atoi(record[6])
			weapons = append(weapons, &Weapon{
				Name:      record[0],
				Type:      "missile",
				Range:     r,
				Damage:    record[3],
				HP:        sHP,
				CurrentHP: sHP,
			})
		}
	}
	return weapons
}
