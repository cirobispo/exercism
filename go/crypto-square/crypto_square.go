package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func removeRunes(pt string) []rune {
	message:=make([]rune, 0, len(pt))
	pt=strings.ToLower(pt)
	for i:= range pt {
		if unicode.IsLetter(rune(pt[i])) || unicode.IsDigit(rune(pt[i]))  {
			message=append(message, rune(pt[i]))
		}
	}

	return message
}

func getSquareSize(sq float64) (float64, float64) {
	r, c:= math.Ceil(sq), math.Trunc(sq)
	if sq > c {
		c++
	}
	return r, c
} 

func Encode(pt string) string {
	message:=removeRunes(pt)

	size:=len(message)
	rCount, cCount:=getSquareSize(math.Sqrt(float64(size)))

	if size > 1 { 
		result:=make([]rune, 0, size)
		for c:=0;c<int(cCount); c++{
			result = append(result, ' ')
			for r:=0;r<size; r+=int(rCount) {
				if index:=r+c; index < size { 
					result = append(result, message[index])
				} else {
					result = append(result, ' ')
				}
			}
		}

		return string(result)[1:]
	}
	return pt
}
