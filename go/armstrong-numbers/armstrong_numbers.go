package armstrong

import (
	"math"
)

func IsNumber(n int) bool {
	result := true
	if n > 0 {
		exp := int(math.Ceil(math.Log10(float64(n))))
		if n%10 == 0 {
			exp++
		}

		var sum float64
		var remvalue int = n
		for i := 1; i <= exp; i++ {
			divisor := math.Pow(10, float64(exp-i))
			base := int(float64(remvalue) / divisor)
			sum += math.Pow(float64(base), float64(exp))
			remvalue -= int(divisor * float64(base))
		}

		result = (int(sum) == n)
	}
	return result
}
