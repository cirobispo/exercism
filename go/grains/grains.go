package grains

import (
	"errors"
	"math"
)

func square(number int) uint64 {
	v := math.Pow(2, float64(number))
	return uint64(v)
}

func Square(number int) (uint64, error) {
	if number > 0 && number < 65 {
		return square(number), nil
	}
	return 0, errors.New("value out of bounds")
}

func Total() uint64 {
	v := square(64)
	return v
}
