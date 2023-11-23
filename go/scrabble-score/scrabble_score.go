package scrabble

import (
	"regexp"
	"strings"
)

func Score(word string) int {
	var letters = []struct {
		regex string
		value int
	}{{"(A|E|I|O|U|L|N|R|S|T)", 1}, {"(D|G)", 2}, {"(B|C|M|P)", 3}, {"(F|H|V|W|Y)", 4}, {"(K)", 5}, {"(J|X)", 8}, {"(Q|Z)", 10}}

	var result int
	for i := 0; i < len(letters); i++ {
		r := regexp.MustCompile(letters[i].regex)
		result += len(r.FindAllString(strings.ToUpper(word), -1)) * letters[i].value
	}
	return result
}
