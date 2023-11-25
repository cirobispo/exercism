package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	result:=make([]string, 0, 4*takeDown)
	for sbs:=takeDown; sbs>0; sbs-- {
		result=append(result, recite1part(startBottles)...)
		result=append(result, reciteMidpart(startBottles)...)
		result=append(result, reciteLastpart(startBottles)...)
		
		startBottles--

		if sbs > 1 {
			result=append(result, "")
		}
	}

	return result
}
func numbers(index int, lowerCase bool) string {
	numbers:=[]string{"No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	if lowerCase {
		return strings.ToLower(numbers[index])
	}

	return numbers[index]
}

func recite1part(startBottles int) []string {
	result:=make([]string, 0, 2)
	for i:=0; i<2;i++ {
		btl:="bottles"
		if startBottles == 1 {
			btl="bottle"
		}
		text:=fmt.Sprintf("%s green %s hanging on the wall,", numbers(startBottles, false), btl)
		result=append(result, text)
	}

	return result
}

func reciteMidpart(startBottles int) []string {
	result:=make([]string, 0)
	result=append(result, "And if one green bottle should accidentally fall,")

	return result
}

func reciteLastpart(startBottles int) []string {
	result:=make([]string, 0)
	rest, btl:=startBottles-1, "bottles"
	if rest == 1 {
		btl="bottle"
	}

	text:=fmt.Sprintf("There'll be %s green %s hanging on the wall.", numbers(rest, true), btl)
	result=append(result, text)

	return result
}