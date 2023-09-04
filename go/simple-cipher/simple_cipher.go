package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
	dist int
}

type vigenere string

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if (distance > 0 && distance < 26) || (distance > -26 && distance < 0) {
		return shift{dist: distance}
	}

	return nil
}

func (c shift) Encode(input string) string {
	return cipher(strings.ToLower(input), c.dist)
}

func (c shift) Decode(input string) string {
	return cipher(strings.ToLower(input), c.dist*-1)
}

func NewVigenere(key string) Cipher {
	if key != "" {
		re := regexp.MustCompile(`^a+$|^\d+$|,| |^[A-Z]+$`)
		wrongOnes := re.FindAllString(key, -1)
		if len(wrongOnes) == 0 {
			var result vigenere = vigenere(key)
			return result
		}
	}

	return nil
}

func vigenereCipher(key, input string, inverse bool) string {
	pK := 0
	result := make([]rune, 0, len(input))
	for _, c := range strings.ToLower(string(input)) {
		if unicode.IsLetter(c) {
			var nc rune
			if inverse {
				dist := (int(string(key)[pK]) - 97) * -1
				nc = changeCharPos(97, 122, c, dist)
			} else {
				dist := int(c) - 97
				nc = changeCharPos(97, 122, rune(string(key)[pK]), dist)
			}

			result = append(result, nc)
			if pK < len(string(key))-1 {
				pK++
			} else {
				pK = 0
			}
		}
	}

	return string(result)
}

func (v vigenere) Encode(input string) string {
	return vigenereCipher(string(v), input, false)
}

func (v vigenere) Decode(input string) string {
	return vigenereCipher(string(v), input, true)
}

func changeCharPos(boc, eoc int, char rune, d int) rune {
	if d > 0 {
		if int(char) == eoc {
			char = rune(boc - 1)
		}
		return changeCharPos(boc, eoc, char+1, d-1)
	} else if d < 0 {
		if int(char) == boc {
			char = rune(eoc + 1)
		}
		return changeCharPos(boc, eoc, char-1, d+1)
	}

	return char
}

func charCipher(d rune, distance int) rune {
	boc := 65
	eoc := 90
	if unicode.IsLower(d) {
		boc = 97
		eoc = 122
	}

	return changeCharPos(boc, eoc, d, distance)
}

func cipher(s string, distance int) string {
	var atb []rune = make([]rune, 0)

	for _, r := range s {
		if unicode.IsLetter(r) {
			atb = append(atb, charCipher(r, distance))
		}

		if unicode.IsNumber(r) {
			atb = append(atb, r)
		}
	}

	return strings.ToLower(string(atb))
}
