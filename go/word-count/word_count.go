package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	if re, err := regexp.Compile(`(\w+'t|\w+'s|\w+'re)|\w+`); err == nil {
		words := re.FindAllString(strings.ToLower(phrase), -1)
		var freq Frequency = make(Frequency)
		for _, word := range words {
			freq[word] += 1
		}

		return freq
	}

	return Frequency{}
}
