package dndcharacter

//
// THIS KIND OF EXERCISE IS A PIECE OF CRAP. SUPER BAD DESCRIBED.
//

import (
	"math"
	"math/rand"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	if score > 10 {
		score--
	}
	return int(math.Round((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return rollDices(4)
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constituition := Ability()

	charac := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: int(constituition),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constituition),
	}

	return charac
}

func rollDices(throlls int) int {
	dices := make([]int, 0, throlls)
	for i := 0; i < throlls; i++ {
		dices = append(dices, rand.Intn(6)+1)
	}
	sort.Slice(dices, func(i, j int) bool {
		return dices[i] > dices[j]
	})

	result := 0
	for i := 0; i < throlls-1; i++ {
		result += dices[i]
	}

	return result
}
