package railfence

import (
	"strings"
)

func Encode(message string, rails int) string {
	buildGrid := func(size int) [][]rune {
		result := make([][]rune, 0, size)
		for i := 0; i < size; i++ {
			result = append(result, make([]rune, 0, size))
		}
		return result
	}

	gridToText := func(grid [][]rune) string {
		var sb strings.Builder
		for i := 0; i < len(grid); i++ {
			sb.WriteString(string(grid[i]))
		}

		return sb.String()
	}

	nextRune := func(message string, index *int) rune {
		result := rune(message[*index])
		*index++
		return result
	}

	textLength := len(message)
	verticeSize := rails - 1
	segmentCount := (textLength / verticeSize)
	if textLength%verticeSize != 0 {
		segmentCount++
	}

	index := 0
	grid := buildGrid(rails)
	row, updown := 0, 1

	for s := 0; s < segmentCount; s++ {
		for index < textLength && ((row > 0 && updown == -1) || (row < verticeSize && updown == 1)) {
			grid[row] = append(grid[row], nextRune(message, &index))
			row += updown
		}
		updown *= -1
	}

	return gridToText(grid)
}

func Decode(message string, rails int) string {
	getSegment := func(takeTail bool, head, tail string) *string {
		stack := &head
		if takeTail {
			stack = &tail
		}

		return stack
	}

	nextRune := func(data *string, index int) rune {
		return rune((*data)[index])
	}

	textLength := len(message)
	verticeSize := rails - 1

	headSize := textLength / (verticeSize * 2)
	if (textLength % (verticeSize * 2)) > 0 {
		headSize++
	}

	tailSize := textLength / (verticeSize * 2)
	blockSize := (textLength - (headSize + tailSize)) / (verticeSize - 1)

	updown, row := 1, 0
	var sb strings.Builder
	for b := 0; b < blockSize; b++ {
		segment := getSegment(updown == -1, message[0:headSize], message[textLength-tailSize:])
		sb.WriteRune(nextRune(segment, (b * verticeSize) / (rails*2 - 2)))

		for (updown == 1 && row < verticeSize-1) || (updown == -1 && row > -1) {
			sb.WriteRune(rune(message[headSize+(row*blockSize+b)]))
			row += updown
		}
		updown *= -1
		row += updown
	}
	if (textLength % verticeSize) == 1 {
		sb.WriteRune(rune(message[len(message) - 1]))
	}

	return sb.String()
}
