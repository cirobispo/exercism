package atbash

import (
	"math"
	"strings"
)

//                                           1        11        11 1
//         6   67    7   78    8   89     9 90        01        12 2
//         56789012345678901234567890     78901234567890123456789012
// Plain:  ABCDEFGHIJKLMNOPQRSTUVWXYZ*****abcdefghijklmnopqrstuvwxyz
// Cipher: ZYXWVUTSRQPONMLKJIHGFEDCBA*****zyxwvutsrqponmlkjihgfedcba
//                                        1 11        11        1
//         98   8    87   7    76   6     2 21        10        09 9
//         09876543210987654321098765     21098765432109876543210987

func isItUpper(d rune) bool {
	return (d > 64) && (d < 91)
}

func isItLower(d rune) bool {
	return (d > 96) && (d < 123)
}

func isItLetter(d rune) bool {
	return isItUpper(d) || isItLower(d)
}

func isItNumber(d rune) bool {
	return (d >= '0') && (d <= '9')
}

func charAtbash(d rune) rune {
	bgn := 65
	dif := math.Abs(float64(int(d) - 65))
	if isItLower(d) {
		bgn = 97
		dif = math.Abs(float64(int(d) - 97))
	}

	return rune(bgn + (25 - int(dif)))
}

func Atbash(s string) string {
	var atb []rune = make([]rune, 0)

	spaceAdded := 0
	blockIt := func() {
		if len(atb) > 0 && (len(atb)-spaceAdded)%5 == 0 && atb[len(atb)-1] != ' ' {
			atb = append(atb, ' ')
			spaceAdded++
		}
	}

	for _, r := range s {
		if isItLetter(r) {
			blockIt()
			atb = append(atb, charAtbash(r))
		}

		if isItNumber(r) {
			blockIt()
			atb = append(atb, r)
		}
	}

	return strings.ToLower(string(atb))
}
