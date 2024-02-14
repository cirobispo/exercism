package queenattack

import (
	"fmt"
	"math"
	"strconv"
)

type boardIndex [2]int

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition {
		return false, fmt.Errorf("white and black queen are at the same place, %s", whitePosition)
	}

	wbi := boardPosToIndex(whitePosition)
	bbi := boardPosToIndex(blackPosition)

	if !isValid(wbi) || !isValid(bbi) {
		return false, fmt.Errorf("white/black queen position is not allowed")
	}

	if p := isSameLineOrColumn(wbi, bbi); p > 0 {
		return true, nil
	} else {
		rdif:=math.Abs(float64(wbi[0]) - float64(bbi[0]))
		cdif:=math.Abs(float64(wbi[1]) - float64(bbi[1]))
		return rdif == cdif, nil
	}
}

func boardPosToIndex(position string) boardIndex {
	column := int(byte(position[0]) - byte('a'))
	line := 0
	if r, err := strconv.ParseInt(position[1:], 10, 32); err == nil {
		line = int(r - 1)
	}

	return boardIndex{line, column}
}

func isValid(pos boardIndex) bool {
	return (pos[0] >= 0 && pos[0] < 8) && (pos[1] >= 0 && pos[1] < 8)
}

func isSameLineOrColumn(a, b boardIndex) int {
	result := 0
	if a[0] == b[0] {
		result += 1
	}

	if a[1] == b[1] {
		result += 2
	}

	return result
}
