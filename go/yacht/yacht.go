package yacht

import "sort"

type yDice int
type yCategory int

const (
	ycOne yCategory =1
	ycTwo yCategory =2
	ycThree yCategory =3
	ycFour yCategory =4
	ycFive yCategory =5
	ycSix yCategory =6
	ycFullHouse yCategory=10
	ycFourOfAKind yCategory=11
	ycLittleStraight yCategory=12
	ycBigStraight yCategory=13
	ycChoice yCategory=20
	ycYatch yCategory=21
)

func Score(dice []int, category string) int {
	cat:=textToCategory(category)
	if cat == ycYatch {
		_, count:=moda(dice)
		if count == 5 {
			return 50
		}
		return 0
	} else if cat >= ycOne && cat <= ycSix {
		return getCount(dice, yDice(cat)) * int(cat)
	} else if cat == ycChoice{
		return getChoice(dice)
	} else if cat == ycFullHouse {
		if isFullHouse(dice) {
			return getChoice(dice)
		}
	} else if cat == ycFourOfAKind {
		if d, count:=moda(dice); count >= 4 {
			return int(d)*4 
		}
	} else if cat == ycLittleStraight || cat == ycBigStraight {
		if isStraight(dice) {
			if (cat == ycBigStraight && dice[0]==2) || (cat == ycLittleStraight && dice[0]==1) {
				return 30
			}
		}
	}
	return 0
}

func textToCategory(category string) yCategory {
	switch category {
	case "ones":
		return ycOne
	case "twos":
		return ycTwo
	case "threes":
		return ycThree
	case "fours":
		return ycFour
	case "fives":
		return ycFive
	case "sixes":
		return ycSix
	case "full house":
		return ycFullHouse
	case "four of a kind":
		return ycFourOfAKind
	case "little straight":
		return ycLittleStraight
	case "big straight":
		return ycBigStraight
	case "yacht":
		return ycYatch
	default:
		return ycChoice
	}
}

func catalogDices(dices []int) map[yDice]int {
	result:=make(map[yDice]int)
	for i:=range dices {
		d:=dices[i]
		result[yDice(d)]+=1
	}

	return result
}

func isFullHouse(dices []int) bool {
	result:=catalogDices(dices)
	if len(result) == 2 {
		_, count:=moda(dices); return count == 3
	}
	return false
}

func moda(dices []int) (yDice, int) {
	var result yDice
	biggest:=0
	
	for i, v:=range catalogDices(dices) {
		if v > biggest {
			result=i
			biggest=v
		}
	}

	return result, biggest
}

func getCount(dices []int, dice yDice) int {
	result:=catalogDices(dices)
	return result[dice]
}

func getChoice(dice []int) int {
	result:=0
	for i:=range dice {
		result+=dice[i]
	}

	return result
}

func isStraight(dices []int) bool {
	if _, count:=moda(dices); count == 1 {
		sort.Slice(dices, func(i, j int) bool {
			return dices[i] < dices[j]
		})

		for i:=1; i<len(dices); i++ {
			if dices[i]-dices[i-1] != 1 {
				return false
			}
		}
		return true
	}
	return false
}