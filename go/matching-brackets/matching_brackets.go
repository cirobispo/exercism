package brackets

type Blocking struct {
	who      BracketType
	isClosed bool
	children []*Blocking
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

func isBracket(char rune) bool {
	return char == '(' || char == ')'
}

func isSquaredBracket(char rune) bool {
	return char == '[' || char == ']'
}

func isCurlyBracket(char rune) bool {
	return char == '{' || char == '}'
}

func isAnyBracket(char rune) bool {
	return isBracket(char) ||
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

func NewBlocking(char rune) *Blocking {
	bt, _ := WhichBracket(char)
	return &Blocking{who: bt, isClosed: false, children: make([]*Blocking, 0)}
}

func (b *Blocking) AddChild(child *Blocking) {
	b.children = append(b.children, child)
}

func (b *Blocking) Close() {
	b.isClosed = true
}

func BuildBlockingPairs(input string) *Blocking {
	var result *Blocking

	var temp *Blocking
	for i := range input {
		char := rune(input[i])
		if !isAnyBracket(char) {
			continue
		}

		bt, bo := WhichBracket(char)
		if temp == nil && bo == boLeft {
			temp = NewBlocking(char)
			if result == nil {
				result = temp
			} else {
				result.AddChild(temp)
			}
		} else if temp != nil && bo == boLeft {
			temp = NewBlocking(char)
			result.AddChild(temp)
		} else if temp == nil && bo == boRight {
			break
		} else if temp != nil && bo == boRight {
			if bt == temp.who {
				temp.Close()
			} else {
				break
			}

			temp = nil
		}
	}

	return result
}

// if result != nil {
// 	result = parents[len(parents)-1]
// }

func CheckBlocking(b *Blocking) bool {
	if !b.isClosed {
		return false
	}

	result := true
	for i := range b.children {
		child := b.children[i]
		if result = CheckBlocking(child); !result {
			break
		}
	}

	return result
}

func Bracket(input string) bool {
	b := BuildBlockingPairs(input)
	if b == nil {
		return true
	}

	return CheckBlocking(b)
}
