package allergies

import (
	"math"
)

var (
	allAllergies = getAlergies(0)
)

func Allergies(allergies uint) []string {
	result := []string{}

	for allergies > 0 {
		exp := math.Log(float64(allergies)) / math.Log(2)
		if exp <= 16 {
			found := getAlergies(allAllergies, int(math.Pow(2, math.Trunc(exp))))
			if found.(string) != "" {
				result = append(result, found.(string))
			}
			allergies -= uint(math.Pow(2, math.Trunc(exp)))
		}
	}

	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	allFound := Allergies(allergies)
	return exists(allFound, allergen)
}

// is there such a similar function on golang greater than 1.18?
func exists(allergies []string, allergen string) bool {
	result := false
	for i := range allergies {
		if found := allergies[i]; allergen == found {
			result = true
			break
		}
	}

	return result
}

func getAllergiesByID() map[int]string {
	result := map[int]string{
		1: "eggs", 2: "peanuts", 4: "shellfish", 8: "strawberries",
		16: "tomatoes", 32: "chocolate", 64: "pollen", 128: "cats",
	}

	return result
}

// testing interface{} use when it comes to make code "smaller" pipe the data
func getAlergies(data ...interface{}) interface{} {
	if len(data) > 0 {
		var result interface{}
		switch data[0].(type) {
		case uint, int:
			result = getAllergiesByID()
		case map[int]string:
			result = data[0].(map[int]string)[data[1].(int)]
		}

		return result
	}

	return nil
}
