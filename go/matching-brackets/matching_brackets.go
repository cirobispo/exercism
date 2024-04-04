package brackets

type MBracket struct {
	bt         BracketType
	childCount int
	isOpen     bool
}

type BracketType int
type BracketOrientation int

const (
	btBracket        BracketType = 1
	btSquaredBracket BracketType = 2
	btCurlyBracket   BracketType = 3
)

const (
	boLeft  BracketOrientation = 0
	boRight BracketOrientation = 1
)

func isRegularBracket(char rune) bool {
	return char == '(' || char == ')'
}

func isSquaredBracket(char rune) bool {
	return char == '[' || char == ']'
}

func isCurlyBracket(char rune) bool {
	return char == '{' || char == '}'
}

func isAnyBracket(char rune) bool {
	return isRegularBracket(char) ||
		isSquaredBracket(char) ||
		isCurlyBracket(char)
}

func WhichBracket(char rune) (BracketType, BracketOrientation) {
	side := boLeft
	switch {
	case isCurlyBracket(char):
		if char == '}' {
			side = boRight
		}
		return btCurlyBracket, side

	case isSquaredBracket(char):
		if char == ']' {
			side = boRight
		}
		return btSquaredBracket, side

	default:
		if char == ')' {
			side = boRight
		}
		return btBracket, side
	}
}

func isBracketsInPairs(input string) bool {
	brackets_found := make([]*MBracket, 0)

	lastOpen := func() *MBracket {
		size := len(brackets_found)
		for i := size - 1; i >= 0; i-- {
			bf := brackets_found[i]
			if bf.isOpen {
				return bf
			}
		}

		return nil
	}

	for i := range input {
		char := rune(input[i])
		if isAnyBracket(char) {
			bt, bo := WhichBracket(char)
			lb := lastOpen()
			if bo == boRight && lb != nil && lb.bt == bt {
				lb.isOpen = false
			} else {
				brackets_found = append(brackets_found, &MBracket{bt: bt, childCount: 0, isOpen: true})
				if lb != nil && lb.isOpen {
					lb.childCount++
				}
			}
		}
	}

	result := true
	for i := range brackets_found {
		bf := brackets_found[i]
		if bf.isOpen {
			result = false
			break
		}
	}
	return result
}

func Bracket(input string) bool {
	return isBracketsInPairs(input)
}
