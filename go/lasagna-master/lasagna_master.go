package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodle := 0
	sauce := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodle += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}

	return noodle, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	friendLI := friendList[len(friendList)-1]
	myList[len(myList)-1] = friendLI
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	result := make([]float64, 0, len(amounts))
	result = append(result, amounts...)
	for i := 0; i < len(amounts); i++ {
		result[i] = result[i] * float64(portions) / 2
	}

	return result
}
