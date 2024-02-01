package wordsearch

import (
	"fmt"
	"math"
)

type puzzleMapping struct {
	mapping       [][]rune
	height, width int
}

func NewPuzzleMapping(puzzle []string) puzzleMapping {
	height := len(puzzle)
	width := 0
	if height > 0 {
		width = len(puzzle[0])
	}
	mapping := make([][]rune, height)
	for i := range puzzle {
		p := puzzle[i]
		mapping[i] = append(mapping[i], []rune(p)...)
	}

	return puzzleMapping{mapping: mapping, height: height, width: width}
}

func (p puzzleMapping) wordFitsIn(wordSize, row, col int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < 8; i++ {
		//vWay get *-1 to invert Sin aligning to matrix
		vWay := int(math.Round(math.Sin((math.Pi/4)*float64(i))) * -1)
		hWay := int(math.Round(math.Cos((math.Pi / 4) * float64(i))))
		isHorizontal := hWay > 0 && p.width-col >= wordSize || hWay < 0 && col+1 >= wordSize
		isVertical := vWay > 0 && p.height-row >= wordSize || vWay < 0 && row+1 >= wordSize
		if isHorizontal && vWay == 0 || isVertical && hWay == 0 || hWay != 0 && isHorizontal && vWay != 0 && isVertical {
			result = append(result, []int{vWay, hWay})
		}
	}

	return result
}

func (p puzzleMapping) checkWordIn(word string, rowPos, colPos, vWay, hWay int) (bool, [][]int) {
	found := true
	for idx := range word {
		charMap := p.mapping[rowPos+idx*vWay][colPos+idx*hWay]
		if charMap != rune(word[idx]) {
			found = false
			break
		}
	}

	result := make([][]int, 0)
	if size := len(word); found { // just to let code 1 line smaller
		result = append(result, []int{colPos, rowPos})
		result = append(result, []int{colPos + (size-1)*hWay, rowPos + (size-1)*vWay})
	}

	return found, result
}

func (p puzzleMapping) Exists(word string) (bool, [][]int) {
	if p.height < 0 || p.width < 0 {
		return false, make([][]int, 0)
	}

	found, result := false, [][]int{}
	for l := 0; l < p.height; l++ {
		for c := 0; c < p.width; c++ {
			if dirs := p.wordFitsIn(len(word), l, c); len(dirs) > 0 {
				for i := range dirs {
					dir := dirs[i]
					found, result = p.checkWordIn(word, l, c, dir[0], dir[1])
					if found {
						goto respond
					}
				}
			}
		}
	}

respond:
	return found, result
}

func addToResult(word string, pm *puzzleMapping, result map[string][2][2]int) error {
	if found, data := pm.Exists(word); found {
		value := [2][2]int{{data[0][0], data[0][1]}, {data[1][0], data[1][1]}}
		result[word] = value
		return nil
	}

	return fmt.Errorf("%s not found", word)
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	pm := NewPuzzleMapping(puzzle)

	var err error
	result := make(map[string][2][2]int)
	for i := range words {
		word := words[i]
		if err = addToResult(word, &pm, result); err != nil {
			break
		}
	}

	return result, err
}
