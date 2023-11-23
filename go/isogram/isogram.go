package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	l := make(map[byte]bool)
	var result bool = true
	for i := 0; i < len(word); i++ {
		_, found := l[word[i]]
		if found {
			result = false
			break
		}
		if word[i] != '-' && word[i] != ' ' {
			l[word[i]] = true
		}
	}
	return result
}
