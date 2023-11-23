package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	var raindrops = []struct {
		divisor int
		sound   string
	}{{3, "Pling"}, {5, "Plang"}, {7, "Plong"}}
	var result string
	for i := 0; i < len(raindrops); i++ {
		rest := number % raindrops[i].divisor
		if rest == 0 {
			result += raindrops[i].sound
		}
	}
	if result == "" {
		result = strconv.Itoa(number)
	}
	return result
}
