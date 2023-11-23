package chessboard

/**
type Square struct {
	rank int
	file string
}
 /**/

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

func rowAmount(column []bool, rank int) int {
	var total int = 0
	boolToInt := func(value bool) int {
		if value {
			return 1
		}
		return 0
	}

	if rank >= -1 && rank < len(column) {
		if rank == -1 {
			for _, value := range column {
				total += boolToInt(value)
			}
		} else if rank >= 0 {
			if len(column) > rank {
				total = boolToInt(column[rank])
			}
		}
	}
	return total
}

func columnAmount(cb Chessboard, rank int) int {
	var amount int = 0
	var files []string = []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for _, column := range files {
		amount += rowAmount(cb[column], rank-1)
	}
	return amount
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var amount int = 0
	amount = rowAmount(cb[file], -1)
	return amount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var amount int = 0
	if rank == 0 {
		rank = -1
	}
	amount = columnAmount(cb, rank)
	return amount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var amount int = 0
	amount += len(cb) * 8
	return amount
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var amount int = 0
	amount = columnAmount(cb, 0)
	return amount
}
