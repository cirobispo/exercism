package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	lines := strings.Split(s, "\n")
	result := make(Matrix, 0, len(lines))
	cellsSize := -1
	for i := range lines {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			return [][]int{}, fmt.Errorf("some line has no cells")
		}

		cells := strings.Split(line, " ")
		row := make([]int, 0, len(cells))
		if size := len(cells); cellsSize != -1 && size != cellsSize {
			return [][]int{}, fmt.Errorf("lines have different cell count")
		} else {
			cellsSize = size
		}
		for j := range cells {
			cell := cells[j]
			value, err := strconv.Atoi(cell)
			if err != nil {
				return [][]int{}, fmt.Errorf("value %v is not a proper int", cell)
			}
			row = append(row, value)
		}
		result = append(result, row)
	}

	return result, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	lineCount := len(m)
	if lineCount < 1 {
		return [][]int{}
	}

	cellCount := len(m[0])
	result := make([][]int, 0, cellCount)
	for c := 0; c < cellCount; c++ {
		result = append(result, []int{})
		for l := 0; l < lineCount; l++ {
			result[c] = append(result[c], m[l][c])
		}
	}

	return result
}

func (m Matrix) Rows() [][]int {
	lineCount := len(m)
	if lineCount < 1 {
		return [][]int{}
	}

	result := make([][]int, 0, lineCount)
	for i := 0; i < lineCount; i++ {
		result = append(result, []int{})
		result[i] = append(result[i], m[i]...)
	}

	return result
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row > len(m)-1 {
		return false
	}

	line := m[row]
	if col < 0 || col > len(line)-1 {
		return false
	}

	m[row][col] = val
	return true
}
