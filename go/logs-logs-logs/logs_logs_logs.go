package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var result string
	for _, c := range log {
		if c == '❗' {
			result = "recommendation"
			break
		} else if c == '🔍' {
			result = "search"
			break
		} else if c == '☀' {
			result = "weather"
			break
		}
	}

	if result == "" {
		result = "default"
	}

	return result
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var result []rune = make([]rune, 0, len(log))
	for _, c := range log {
		if c == oldRune {
			result = append([]rune(result), newRune)
		} else {
			result = append([]rune(result), c)
		}
	}
	return string(result)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
