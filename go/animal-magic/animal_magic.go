package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	result := make([]string, 0, 8)

	shuffle := func(ItemsToShuffle int) {
		type callback func(left int)
		var myCallback callback

		shuffle := func(leftToShuffle int) {
			index := rand.Intn(leftToShuffle)
			result = append([]string(result), animals[index])
			animals = append(animals[0:index], animals[index+1:leftToShuffle]...)
			leftToShuffle -= 1
			if leftToShuffle > 1 {
				myCallback(leftToShuffle)
			} else {
				result = append(result, animals[0])
			}
		}
		myCallback = shuffle
		myCallback(ItemsToShuffle)
	}
	shuffle(len(animals))
	return result
}
