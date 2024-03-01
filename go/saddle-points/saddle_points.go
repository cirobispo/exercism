package matrix

import (
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Pair struct {
	row, col int
}

type Matrix struct {
	values [][]int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	result := make([][]int, 0, len(rows))
	for i := range rows {
		row := rows[i]
		columns := strings.Split(row, " ")
		result = append(result, make([]int, 0, len(columns)))
		for j := range columns {
			col := columns[j]
			value, _ := strconv.ParseInt(col, 10, 64)
			result[i] = append(result[i], int(value))
		}
	}

	return &Matrix{values: result}, nil
}

func (m Matrix) getRow(row int) []int {
	size := len(m.values)
	if row < 0 || row > size-1 {
		return []int{}
	}

	result := make([]int, 0, size)
	result = append(result, m.values[row]...)
	return result
}

func (m Matrix) getColumn(col int) []int {
	size := 0
	if len(m.values) > 0 {
		size = len(m.values[0])
	}

	if col < 0 || col > size-1 {
		return []int{}
	}

	result := make([]int, 0, size)
	for i := range m.values {
		row := m.values[i]
		result = append(result, row[col])
	}

	return result
}

type NEst int

const (
	NBigest   = 0
	NSmallest = 1
)

func getNEstIndex(items []int, what NEst) int {
	if len(items) < 1 {
		return -1
	}

	last := 0
	for i := range items {
		isBiggest:=(what == NBigest && items[last] < items[i])
		isSmallest:=(what == NSmallest && items[last] > items[i])
		if isBiggest || isSmallest {
			last = i
		}
	}

	return last
}

func (m Matrix) getBigColSmallLine() []Pair {	
	size := len(m.values)
	if row:=m.getRow(0); size == 0 || row[getNEstIndex(row, NBigest)] < 1 {
		return []Pair{}
	}

	result := make([]Pair, 0)
	for i := 0; i < size; i++ {
		row:=m.getRow(i)
		BColumnInLine:=getNEstIndex(row, NBigest)
		column:=m.getColumn(BColumnInLine)
		SRowInColumn:=getNEstIndex(column, NSmallest)
		for c:=BColumnInLine; c< len(row); c++ {
			if row[c] != row[BColumnInLine] {
				continue
			}
			column=m.getColumn(c)
			if column[SRowInColumn] == column[i] && (row[c] == column[i]) {
				result = append(result, Pair{row: i+1, col:c+1})
			}
		}
	}

	return result
}

func (m *Matrix) Saddle() []Pair {
	return m.getBigColSmallLine()
}
