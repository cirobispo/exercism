package brackets

type Blocking struct {
	parent   *Blocking
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

func NewBlocking(char rune, parent *Blocking) *Blocking {
	bt, _ := WhichBracket(char)
	return &Blocking{parent: parent, who: bt, isClosed: false, children: make([]*Blocking, 0)}
}

func (b *Blocking) AddChild(child *Blocking) {
	b.children = append(b.children, child)
}

func (b *Blocking) Close() {
	b.isClosed = true
}

func BuildBlockingPairs(input string) *Blocking {
	var result *Blocking

	addToResult := func(char rune) {
		if result == nil {
			result = NewBlocking(char, nil)
		} else {
			temp := NewBlocking(char, result)
			result.AddChild(temp)
			result = temp
		}
	}

	for i := range input {
		char := rune(input[i])
		if isAnyBracket(char) {
			bt, bo := WhichBracket(char)
			if bo == boRight {
				if result == nil || result.isClosed || result.who != bt {
					addToResult(char)
					break
				}
				result.Close()
				result = result.parent
				continue
			}

			addToResult(char)
		}
	}
	return result
}

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
