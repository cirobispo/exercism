package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	msg string
}

func (s *SillyNephewError) Error() string {
	return s.msg
}

var negativeFodderError = errors.New("negative fodder")
var divisionByZeroError = errors.New("division by zero")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	var result float64 = 0
	var error_result error
	if weightFodder != nil {
		fodderAmount, err := weightFodder.FodderAmount()
		if fodderAmount < 0 {
			if err == ErrScaleMalfunction {
				error_result = negativeFodderError
			} else if err == nil {
				error_result = negativeFodderError
			} else {
				error_result = err
			}
		} else {
			if cows < 0 {
				msg := fmt.Sprintf("silly nephew, there cannot be %v cows", cows)
				error_result = &SillyNephewError{msg}
			} else if cows == 0 {
				error_result = divisionByZeroError
			} else {
				if err == ErrScaleMalfunction {
					result = (fodderAmount * 2) / float64(cows)
				} else if err != nil {
					error_result = err
				} else {
					result = fodderAmount / float64(cows)
				}
			}
		}
	}

	return result, error_result
}
