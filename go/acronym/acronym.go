// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	fmt.Println(s)
	return breakIntoInicials(s)
}


func breakIntoInicials(s string) string {
	rx:=regexp.MustCompile(`\w*`)
	s=strings.Replace(s, "'", "", -1) // remove '
	s=strings.Replace(s, "_", "", -1) // remove _
	
	words:=rx.FindAllString(s, -1)

	result :=make([]rune, 0, len(words))
	for _, item:=range words {
		item=strings.TrimSpace(item)
		if len(item) > 0 && item!="-" {
			result = append(result, rune(item[0]))
		}
	}
	return strings.ToUpper( string(result))
}