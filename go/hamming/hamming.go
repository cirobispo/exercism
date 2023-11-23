package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var result int = 0
	var err error

	oddWord := len(a) != len(b)

	if !oddWord {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				result += 1
			}
		}
	} else {
		err = errors.New("index out of bounds")
	}

	return result, err
}
