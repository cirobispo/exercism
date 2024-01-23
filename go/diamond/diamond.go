package diamond

import (
	"fmt"
	"strings"
)

func Gen(char byte) (string, error) {
	if char >='A' && char <= 'Z' {
		diff:=int(char-65)
		return fmt.Sprint(drawDiamond(char, diff)), nil
	}
	return "", fmt.Errorf("some error")
}

func drawDiamond(startChar byte, size int) string {
	drawHalf:=func(begin, max, inc int) string {
		var result strings.Builder

		edgeDist:=0
		i:=begin 
		for {
			result.WriteString(drawLine(startChar-byte(i), size*2+1, edgeDist+i))
			if (inc == -1 && i <= 0) || (inc > 0 && i >= max) {
				break
			}
			result.WriteString(fmt.Sprintln())
			i+=inc
		}
		
		return result.String()
	}

	result:=drawHalf(size, 0, -1)
	if size > 0 {
		result+="\n" + drawHalf(1, size, 1)
	}

	return result
}

func drawLine(char byte, lineSize, eborder int) string {
	var sb strings.Builder
	if lineSize - eborder > 0 {
		sb.WriteString(strings.Repeat(" ", eborder))
		if (eborder*2 + 1) != lineSize {
			sb.WriteRune(rune(char))
			sb.WriteString(strings.Repeat(" ", lineSize-(eborder+1) * 2))
			sb.WriteRune(rune(char))
		} else {
			sb.WriteRune(rune(char))
		}
		sb.WriteString(strings.Repeat(" ", eborder))
	}
	return sb.String()
}