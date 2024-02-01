package piglatin

import (
	"strings"
)

type RuleAction func(word string) (bool, int)

type CR interface {
	execute(word string) (bool, int)
	next() CR
}

type Rule struct {
	action RuleAction
	next   *Rule
}

func NewRule(action RuleAction, rule *Rule) *Rule {
	return &Rule{action: action, next: rule}
}

func (r *Rule) execute(word string) (bool, int) {
	rule := r
	ok, pos := false, -1
	for rule != nil {
		if ok, pos = rule.action(word); ok {
			break
		}
		rule = rule.next
	}

	return ok, pos
}

func Sentence(sentence string) string {
	var result strings.Builder
	rule := buildRules()
	msg := strings.ToLower(sentence)
	size := len(sentence)
	var word strings.Builder
	for i := range msg {
		if i == size-1 || msg[i] == ' ' {
			if msg[i] != ' ' {
				word.WriteByte(msg[i])
			}
			if data := word.String(); len(data) > 0 {
				if ok, pos := rule.execute(data); ok {
					newWord := doPigLatin(data, pos)
					result.WriteString(newWord)
				} else {
					result.WriteString(data + "ay")
				}
			}
			word.Reset()
			if msg[i] == ' ' {
				result.WriteString(" ")
			}
		} else {
			word.WriteByte(msg[i])
		}
	}

	return result.String()
}

func doPigLatin(text string, pos int) string {
	begin := text[0:pos]
	end := text[pos:]
	return end + begin + "ay"
}

func buildRules() *Rule {
	r4 := NewRule(isFourthRule, nil)
	r3 := NewRule(isThirdRule, r4)
	r2 := NewRule(isSecondRule, r3)
	qu := NewRule(func(word string) (bool, int) { return startWith(word, "qu") }, r2)
	r1 := NewRule(isFirstRule, qu)
	return r1
}

func isFirstRule(word string) (bool, int) {
	if len(word) > 0 {
		// if isLikeVowel(word) {
		// 	return true, 2
		// } else {
		return isVowel(word[0]), 0
		// }
	}
	return false, -1
}

func isSecondRule(word string) (bool, int) {
	if size := len(word); size > 0 {
		if ok, _ := startWith(word, "yt"); ok {
			return true, 0
		} else if ok, _ := startWith(word, "xr"); ok {
			return true, 0
		} else if !isVowel(word[0]) {
			if ok, pos := firstVowel(word); ok {
				begin := word[:pos]
				if begin != "sq" {
					return true, pos
				}
			}
		}
	}

	return false, -1
}

func isThirdRule(word string) (bool, int) {
	if size := len(word); size > 0 {
		if isConsonantCluster(word, 2) && word[1] == 'q' && word[2] == 'u' {
			return true, 3
		}
	}

	return false, -1
}

func isFourthRule(word string) (bool, int) {
	if size := len(word); size > 0 {
		if ok, pos := startWith(word, "yt"); ok && pos == 0 {
			return ok, pos
		} else if ok, pos := startWith(word, "xr"); ok && pos == 0 {
			return ok, pos
		} else if isConsonantCluster(word, 2) && size > 2 && word[2] == 'y' {
			return true, 2
		} else if size >= 2 && !isVowel(word[0]) && word[1] == 'y' {
			return true, 1
		}
	}

	return false, -1
}

func isVowel(c byte) bool {
	if c > 0 {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	return false
}

func firstVowel(word string) (bool, int) {
	ok, pos := false, -1
	for i := range word {
		if isVowel(word[i]) {
			ok, pos = true, i
			break
		}
	}

	return ok, pos
}

func isLikeVowel(v string) bool {
	if size := len(v); size > 1 {
		sub := v[:2]
		return sub == "xr" || sub == "yt"
	}

	return false
}

func startWith(word, start string) (bool, int) {
	if ss := len(start); len(word) >= ss {
		return word[:ss] == start, ss
	}

	return false, -1
}

func isConsonantCluster(word string, maxcluster int) bool {
	result := false
	if size := len(word); size > 0 && size > maxcluster {
		result = true
		for i := 0; i < maxcluster; i++ {
			if isVowel(word[i]) {
				result = false
				break
			}
		}
	}

	return result
}
