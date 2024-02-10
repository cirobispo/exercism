package minesweeper

import (
	"fmt"
	"math"
	"strings"
)

type mineSweeper [][]int

func NewMineSwepper(board []string) mineSweeper {
	result := make([][]int, 0, len(board))
	for r := range board {
		row := board[r]
		result = append(result, make([]int, 0, len(row)))
		for c := range row {
			cel := row[c]
			value := 0
			if cel == '*' {
				value = -1
			}
			result[r] = append(result[r], value)
		}
	}

	return mineSweeper(result)
}

func (m mineSweeper) MapAdjacentMines() {
	for r := range [][]int(m) {
		row := m[r]
		for c := range row {
			if m[r][c] == -1 {
				adjacents := m.getAdjacents(r, c)
				for idx := range adjacents {
					adj := adjacents[idx]
					m[adj[0]][adj[1]]++
				}
			}
		}
	}
}

func (m mineSweeper) getAdjacents(row, col int) [][]int {
	height := len(m)
	width := len(m[row])
	result := make([][]int, 0)
	for i := 0; i < 8; i++ {
		vWay := -1 * int(math.Round(math.Sin((math.Pi/4)*float64(i))))
		hWay := int(math.Round(math.Cos((math.Pi / 4) * float64(i))))
		if col+hWay < width && col+hWay >= 0 && row+vWay < height && row+vWay >= 0 {
			if m[row+vWay][col+hWay] != -1 {
				result = append(result, []int{row + vWay, col + hWay})
			}
		}
	}

	return result
}

func (m mineSweeper) GetBoard() []string {
	result := make([]string, 0, len(m))
	for r := range [][]int(m) {
		row := m[r]
		var sb strings.Builder
		for c := range row {
			if value := m[r][c]; value == -1 {
				sb.WriteRune('*')
			} else if value == 0 {
				sb.WriteRune(' ')
			} else {
				sb.WriteString(fmt.Sprintf("%d", value))
			}
		}
		result = append(result, sb.String())
	}

	return result
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	ms := NewMineSwepper(board)
	ms.MapAdjacentMines()

	return ms.GetBoard()
}
