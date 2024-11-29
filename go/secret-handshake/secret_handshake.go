package secret

import (
	"math"
)

func codeToStrings(code int) []string {
	str_codes:=[]string { "wink", "double blink", "close your eyes", "jump" }

	start,step:=0, 1
	if code > int(math.Pow(2, float64(len(str_codes)))-1)  {
		start, step=len(str_codes) - 1,-1
		code-=int(math.Pow(2, float64(len(str_codes))))
	}

	result:=make([]string, 0, 2)
	for code > 0 {
		value:=math.Pow(2, float64(start))
		if code & int(value) != 0 {
			result = append(result, str_codes[start])
			code-=int(value)
		}
		start+=step
	}

	return result
}

func Handshake(code uint) []string {
	return codeToStrings(int(code))
}
