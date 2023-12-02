package wordy

import (
	"fmt"
	"regexp"
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

func parseCalc(textCalc ...string) Calc {
	numOpNum:=func(n1 float64, op operation, n2 float64) Calc {
		c:=Calc{n1, op}
		if result, ok:=c.Execute(n2); ok {
			return result
		}

		return c
	}

	opNum:=func(c Calc, op operation, n2 float64) Calc {
		c.operation = op
		if result, ok:=c.Execute(n2); ok {
			return result
		}

		return c
	}

	textToOp:=func(textOp string) operation {
		switch textOp {
		case "plus":
			return opPlus
		case "minus":
			return opMinus
		case "multiplied by":
			return opMultiply
		case "divided by":
			return opDivide
		default:
			return opEqual
		}
	}

	textToFloat:=func(textValue string, value *float64) bool {
		if result, ok:=strconv.ParseFloat(textValue, 64); ok == nil {
			*value = result
			return true
		}
		return false
	}

	var result Calc = Calc{0, opEqual}
	for i:= range textCalc {
		params:=strings.Split(strings.Trim(textCalc[i], " "), " ")
		if len(params) >= 2 && params[1] == "by" {
			params[0]=params[0] + " by"
			temp:=params[:1]
			temp=append(temp, params[2:]...)
			params=temp
		} 
		
		if pCount:=len(params); pCount >= 3 {
			param0, param2:=0.0,0.0
			if pCount > 3 {
				params=params[1:]
			}

			if textToFloat(params[0], &param0) && textToFloat(params[2], &param2) {
				result=numOpNum(param0, textToOp(params[1]), param2)
			}
		} else if pCount == 2 {
			if param, ok:=strconv.ParseFloat(params[1], 64); ok == nil {
				result=opNum(result, textToOp(params[0]), param)
			}
		} else {
			if param, ok:=strconv.ParseFloat(params[0], 64); ok == nil {
				result=opNum(result, opEqual, param)
			}
		}
	}

	return result
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
		if n2 > 0 {
			result, OK= Calc{c.n1 / n2, opEqual}, true
		} else {
			result, OK= Calc{0, opEqual}, false
		}
	default:
		result, OK= Calc{c.n1, opEqual}, true
	}

	return result, OK
}

func Answer(question string) (int, bool) {
	req:= regexp.MustCompile(`(-?\d*\s)(plus|minus|multiplied by|divided by)(\s-?\d*)`)
	// fqs :=req.FindAll([]byte(question), -1)
	fqs :=req.FindAll([]byte(question), -1)
	fmt.Println(question)
	var result Calc
	for i:=range fqs {
		fmt.Println(string(fqs[i]))
		c:=parseCalc(string(fqs[i]))
		if result.operation == opNaN {
			result = c
		} else {
			result.operation = c.operation
			if r, ok :=result.Execute(c.n1); ok {
				result = r
			}
		}
	}

	return int(result.n1), true
}
