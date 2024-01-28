package house

import (
	"fmt"
	"strings"
)

//
// HOW I HATE THIS STUPID SONGS. BUNCH OF CRAP TEXT TO PUT TOGETHER
//
var verses = []string {
	"This is the horse and the hound and the horn",
	"that belonged to the farmer sowing his corn",
	"that kept the rooster that crowed in the morn",
	"that woke the priest all shaven and shorn",
	"that married the man all tattered and torn",
	"that kissed the maiden all forlorn",
	"that milked the cow with the crumpled horn",
	"that tossed the dog",
	"that worried the cat",
	"that killed the rat",
	"that ate the malt",
	"This is the house that Jack built.",
}

func Verse(v int) string {
	// "that lay in the house that Jack built.",

	if v > 0 && v <= 12 {
		size:=len(verses)
		result:=verses[size - v]
		if v == 1 {
			return result
		} else {
			var sb strings.Builder
			for i:=size-v; i<size; i++ {
				line:=verses[i]
				if i == size-v {
					idx:=strings.Index(line, "the")
					line="This is " + line[idx:]
				}
				if i == (size-1) {
					idx:=strings.Index(line, "the")
					line="that lay in " + line[idx:]
				}
				sb.WriteString(fmt.Sprintf("\n%s", line))
			}

			return sb.String()[1:]
		}
	}

	return ""
}

func Song() string {
	var sb strings.Builder
	for i:=1; i<13; i++ {
		sb.WriteString("\n\n")
		sb.WriteString(Verse(i))
	}

	return sb.String()[2:]
}


