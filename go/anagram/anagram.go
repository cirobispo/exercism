package anagram

import (
	"strings"
)

func buildCharMap(word *string) map[rune]int {
	var result map[rune]int = make(map[rune]int)
	for _, c := range *word {
		result[c] += 1
	}

	return result
}

func compareCharMap(l, r map[rune]int) bool {
	if len(l) == len(r) {
		result := true
		for k, v := range l {
			if search, found := r[k]; !found || v != search {
				result = false
				break
			}
		}
		return result
	}
	return false
}

func Detect(subject string, candidates []string) []string {
	text := strings.ToLower(subject)
	chars_sub := buildCharMap(&text)

	var result = make([]string, 0)
	for _, w := range candidates {
		word := strings.ToLower(w)
		if word != text {
			chars_chk := buildCharMap(&word)
			if compareCharMap(chars_sub, chars_chk) {
				result = append(result, w)
			}
		}
	}

	return result
}
