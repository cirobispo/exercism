package interest

import (
	"math"
)

func InterestRatePerLevel(balance float64, splitWhenPossible bool) (float32, float64) {
	if balance < 0 {
		return float32(3.213), balance
	} else if balance < 1000 {
		return float32(0.5), balance
	} else if balance < 5000 {
		return float32(1.621), balance
	} else {
		if splitWhenPossible {
			return float32(2.475), balance - AnnualBalanceUpdate(5000)
		} else {
			return float32(2.475), balance
		}
	}
}

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	value, _ := InterestRatePerLevel(balance, false)
	return value
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return (balance * float64(InterestRate(balance)) / 100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	calc := func(pv, fv float64, rate float64) int {
		yearsToGo := int(math.Ceil(math.Log(fv/pv) / math.Log(1+rate)))
		//fmt.Printf("\n\nPV: %v, FV: %v, rate: %v, Years: %v", pv, fv, rate, yearsToGo)
		return yearsToGo
	}

	if targetBalance > 5000 {
		var years int
		for balance < targetBalance {
			balance = AnnualBalanceUpdate(balance)
			years++
		}
		return years
	} else {
		rate := float64(InterestRate(balance) / 100)
		return calc(balance, targetBalance, rate)
	}
}
