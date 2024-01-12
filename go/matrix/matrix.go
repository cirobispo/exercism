package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	columns, lines int
	values         []int
}

func New(s string) (*Matrix, error) {
	lines := strings.Split(s, "\n")

	buildResult := func(lines []string) (*Matrix, error) {
		var result *Matrix
		var err = errors.New("error building matrix")
		//defining its size
		if ls := len(lines); ls > 0 {
			for i := range lines {
				line := lines[i]
				columns := strings.Split(line, " ")
				if cs := len(columns); cs > 0 {
					result = &Matrix{lines: ls, columns: cs}
					err = nil
				}
			}
		}

		return result, err
	}

	if result, err := buildResult(lines); err == nil {
		for i := range lines {
			line := lines[i]
			columns := strings.Split(line, " ")
			for j := range columns {
				value, _ := strconv.ParseInt(columns[j], 10, 64)
				result.values = append(result.values, int(value))
			}
		}

		return result, nil
	} else {
		return &Matrix{}, err
	}
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	result := make([][]int, 0, m.lines*m.columns)

	for i := 0; i < m.columns; i++ {
		for j := 0; j < m.lines; j++ {
			result[i][j] = m.values[(j*m.columns)+i]
		}
	}

	return result
}

func (m *Matrix) Rows() [][]int {
	result := make([][]int, 0, m.lines*m.columns)

	for j := 0; j < m.lines; j++ {
		for i := 0; i < m.columns; i++ {
			result[j][i] = m.values[(j*m.columns)+i]
		}
	}

	return result
}

func (m *Matrix) Set(row, col, val int) bool {
	onBoard := (row >= 0 && row < m.lines) && (col >= 0 && col < m.columns)
	if onBoard {
		m.values[(row*m.columns)+col] = val
		return true
	}

	return false
}
