package pythagorean

import "math"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	return []Triplet{}
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	result:=make([]Triplet, 0)
	c:=int(math.Round(float64(p) *.42))
	ab:=int(math.Round(float64(p) *.58))
	dist(&ab, &c)

	b:=int(math.Round(float64(ab) *.57))
	a:=int(math.Round(float64(ab) *.43))
	dist(&b, &a)
	result = append(result, makeTriplet(a, b, c))
	return result
}

func makeTriplet(a, b, c int) Triplet {
	return Triplet{a, b, c}
}

func dist(a, b *int) {
	for {
		if r:=int(math.Sqrt(float64(*a**a - *b**b))); r > 0 {
			*b+=r/2
			*a-=r/2
		} else if r < 0 {
			*a+=r/2
			*b-=r/2
		} else {
			break
		}
	}
}