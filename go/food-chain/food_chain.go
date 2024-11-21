package foodchain

//
// HOW I HATE THIS STUPID SONGS. BUNCH OF CRAP TEXT TO PUT TOGETHER
//

import (
	"fmt"
	"strings"
)

type animal struct {
	name        string
	exclamation string
}

func buildAnimalList() []animal {
	result := make([]animal, 0, 8)
	result = append(result, animal{name: "", exclamation: ""})
	result = append(result, animal{name: "fly", exclamation: ""})
	result = append(result, animal{name: "spider", exclamation: "%s wriggled and jiggled and tickled inside her."})
	result = append(result, animal{name: "bird", exclamation: "How absurd to swallow a bird!"})
	result = append(result, animal{name: "cat", exclamation: "Imagine that, to swallow a cat!"})
	result = append(result, animal{name: "dog", exclamation: "What a hog, to swallow a dog!"})
	result = append(result, animal{name: "goat", exclamation: "Just opened her throat and swallowed a goat!"})
	result = append(result, animal{name: "cow", exclamation: "I don't know how she swallowed a cow!"})
	result = append(result, animal{name: "horse", exclamation: "She's dead, of course!"})
	return result
}

func Verse(v int) string {
	animals := buildAnimalList()
	var sb strings.Builder
	if v > 0 && v < len(animals) {
		if v > 1 && v < 8 {
			sb.WriteString(sheSwallowed(animals[v]))
			sb.WriteString(sheComplain(animals[v], "It", true))
			for i := v; i > 1; i-- {
				sb.WriteString(sheSwallowedToCatch(animals[i], animals[i-1]))
			}
			sb.WriteString(sheWillDie(animals[1]))
		} else {
			sb.WriteString(sheSwallowed(animals[v]))
			if v != 8 {
				sb.WriteString(sheWillDie(animals[1]))
			} else {
				sb.WriteString(sheComplain(animals[v], "", false))
			}
		}
	}
	result := sb.String()
	return result
}

func Verses(start, end int) string {
	var sb strings.Builder
	for i := start; i <= end; i++ {
		sb.WriteString("\n\n")
		sb.WriteString(Verse(i))
	}
	return sb.String()[2:]
}

func Song() string {
	return Verses(1, 8)
}

func sheSwallowed(swallow animal) string {
	return fmt.Sprintf("I know an old lady who swallowed a %s.\n", swallow.name)
}

func sheComplain(swallow animal, extra string, brake bool) string {
	result := swallow.exclamation
	if strings.Contains(swallow.exclamation, "%s") {
		result = fmt.Sprintf(swallow.exclamation, extra)
	}

	if brake {
		result = fmt.Sprintf("%s\n", result)
	}
	return result
}

func sheSwallowedToCatch(swallow, catch animal) string {
	result := fmt.Sprintf("She swallowed the %s to catch the %s.\n", swallow.name, catch.name)
	if swallow.name == "bird" {
		result = fmt.Sprintf("She swallowed the %s to catch the %s %s\n", swallow.name, catch.name, sheComplain(catch, "that", false))
	}
	return result
}

func sheWillDie(swallow animal) string {
	return fmt.Sprintf("I don't know why she swallowed the %s. Perhaps she'll die.", swallow.name)
}
