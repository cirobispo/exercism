package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	digCount := len(digits)
	serieCount := digCount - span + 1
	if digCount >= span && span > 0 {
		var bestProduct int64 = 0
		series := make(map[string]int64, serieCount)
		for i := 0; i < serieCount; i++ {
			serie := digits[i : i+span]
			for _, c := range serie {
				if v, err := strconv.Atoi(string(c)); err == nil {
					if _, f := series[serie]; !f {
						series[serie] = 1
					}
					series[serie] = series[serie] * int64(v)
				} else {
					return 0, errors.New("conversion error")
				}
			}

			if series[serie] > bestProduct {
				bestProduct = series[serie]
			}
		}
		return bestProduct, nil
	}

	return 0, errors.New("impossible make series")
}
