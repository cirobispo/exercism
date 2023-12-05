package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

type operation int 

const (
	opNaN       operation=0
	opPlus      operation=1
	opMinus     operation=2
	opMultiply  operation=3
	opDivide    operation=4
	opEqual     operation=5
)

type Calc struct {
	n1 float64
	operation 
}

func(c Calc) Execute(n2 float64) (Calc, bool) {
	var result Calc
	var OK bool
	switch c.operation {
	case opPlus:
		result, OK= Calc{c.n1 + n2, opEqual}, true
	case opMinus:
		result, OK= Calc{c.n1 - n2, opEqual}, true
	case opMultiply:
		result, OK= Calc{c.n1 * n2, opEqual}, true
	case opDivide:
		if n2 != 0 {
			result, OK= Calc{c.n1 / n2, opEqual}, true
		} else {
			result, OK= Calc{0, opEqual}, false
		}
	default:
		result, OK= Calc{c.n1, opEqual}, true
	}

	return result, OK
}

func textToOp(textOp string) operation {
	switch textOp {
	case "plus":
		return opPlus
	case "minus":
		return opMinus
	case "multiplied":
		return opMultiply
	case "divided":
		return opDivide
	default:
		return opEqual
	}
}

type WordState struct {
	state string
	calc Calc
}

func NewWordState(question string) (*WordState, string) {
	if word, rest:=getNextWord(question); word != "" {
		return &WordState{state: word, calc: Calc{}}, rest
	}
	return nil, question
}

func (w *WordState) UpdateState(question string, verifier func(next string, calc *Calc) bool) (string, bool) {
	OK:=false
	next, rest:=getNextWord(question)
	if verifier != nil {
		OK=verifier(next, &w.calc)
		if OK {
			w.state = next
		}
	}

	return rest, OK
}


func getNextWord(text string) (string, string) {
	words:=strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	if len(words) > 0 {
		if len(words[0]) == len(text) {
			return words[0], ""
		}
		return words[0], text[len(words[0])+1:]
	}
	return "", text
}

func stateWhat(question string, ws *WordState) (string, bool) {
	rest, ok:=ws.UpdateState(question, func(next string, calc *Calc) bool {
		return next == "is"
	})

	return rest, ok
}

func stateIs(question string, ws *WordState) (string, bool) {
	rest, ok:=ws.UpdateState(question, func(next string, calc *Calc) bool {
		if next[len(next)-1] == '?' {
			next=next[:len(next)-1]
		}
		if value, err:=strconv.ParseInt(next, 0, 64); err == nil {
			calc.n1 = float64(value)
			return true
		}
		return false
	})

	return rest, ok
}

func stateN1(question string, ws *WordState) (string, bool) {
	rest, ok:=ws.UpdateState(question, func(next string, calc *Calc) bool {
		if isOp:= next == "plus" || next == "minus" || next == "multiplied" || next == "divided"; isOp {
			ws.calc.operation=textToOp(next)
			return true
		}
		return false
	})

	return rest, ok
}

func stateOperation(question string, ws *WordState) (string, bool) {
	rest, ok:=ws.UpdateState(question, func(next string, calc *Calc) bool {
		if next[len(next)-1] == '?' {
			next=next[:len(next)-1]
		}
		if value, err:=strconv.ParseInt(next, 0, 64); err == nil {
			result, _:= calc.Execute(float64(value))
			calc.n1 = result.n1
			return true
		}
		return false
	})

	return rest, ok
}

func stateBy(question string, ws *WordState) (string, bool) {
	rest, ok:=ws.UpdateState(question, func(next string, calc *Calc) bool {
		return next == "by"
	})

	return rest, ok
}


func Answer(question string) (int, bool) {
	fmt.Println(question)
	ok:=false
	ws, rest:=NewWordState(question)
	if (ws != nil) {
		for len(rest) > 0 {
			switch ws.state {
			case "what":
				if rest, ok=stateWhat(rest, ws); !ok {
					goto endloop
				}
			case "is":
				if rest, ok=stateIs(rest, ws); !ok {
					goto endloop
				}
				ws.state = "n1"
			case "n1": 
				if rest, ok=stateN1(rest, ws); !ok {
					goto endloop
				}
				if ws.state == "multiplied" || ws.state == "divided" {
					ws.state = "by"
				} else {
					ws.state = "operation"
				}
			case "operation":
				if rest, ok=stateOperation(rest, ws); !ok {
					goto endloop
				}
				ws.state = "n1"
			case "by":
				if rest, ok=stateBy(rest, ws); !ok {
					goto endloop
				}
				ws.state = "operation"
			default:
				ok=false
				goto endloop
			}
		}
	}
	
endloop:
	fmt.Println(ws.calc.n1)
	if ok {
		return int(ws.calc.n1), true
	} 

	return 0, false
}
