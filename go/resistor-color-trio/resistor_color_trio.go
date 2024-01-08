package resistorcolortrio

import (
	"fmt"
	"math"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	size:=len(colors)
	value:=0
	subcolors:=colors[0:2]
	for i:=range subcolors {
		c:=int(math.Pow(10, float64(size-i-2)))
		value+=c*getColorValue(subcolors[i])
	}

	if value > 0 {
		lcv:=getColorValue(colors[size-1])
		value=value*int(math.Pow(10, float64(lcv)))
		return getValue(value)
	}

	return "0 ohms"
}

func getValue(value int) string {
	mult:=int(math.Trunc(math.Log10(float64(value)))) / 3
	value=value / int(math.Pow(10, float64(mult * 3)))

	return fmt.Sprintf("%d %sohms", value, getSciNumDesc(mult))
}

func getColorValue(color string) int {
	colors:=map[string]int{
		"black":0, "brown":1, "red":2, "orange":3, "yellow":4,
		"green":5, "blue":6, "violet":7, "grey":8, "white":9,
	}

	return colors[color]
}

func getSciNumDesc(mult int) string {
	if mult !=0 {
		switch mult {
			case 1: return "kilo"
			case 2: return "mega"
			case 3: return "giga"
			case -1: return "mili"
			case -2: return "micro"
			case -3: return "nano"
		}
	} 
	return ""
}
