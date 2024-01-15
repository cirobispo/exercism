package perfect

import (
	"errors"
)

// Define the Classification type here.
type Classification uint

const (
	ClassificationPerfect   Classification = 3
	ClassificationAbundant  Classification = 2
	ClassificationDeficient Classification = 1
)

var ErrOnlyPositive = errors.New("only positive number are allowed")

func getDivider(n int64) int64 {
	result := int64(0)
	for i := int64(1); i < n; i++ {
		if r := n % int64(i); r == 0 {
			result += i
			if result > n {
				break
			}
		}
	}

	return result
}

func Classify(n int64) (Classification, error) {
	var result Classification
	var err error = ErrOnlyPositive

	if n > 0 {
		sum := getDivider(n)
		err = nil
		if r := sum - n; r > 0 {
			result = ClassificationAbundant
		} else if r < 0 {
			result = ClassificationDeficient
		} else {
			result = ClassificationPerfect
		}
	}

	return result, err
}
