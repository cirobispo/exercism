package allyourbase

import (
	"errors"
	"math"
)

type Base struct {
	symbols []rune
	hasSymbols bool
	size int
}

func CreateBaseFromSymbols(symbols []rune) *Base {
	result:=&Base{ hasSymbols: true }
	result.symbols=append(result.symbols, symbols...)
	result.size=len(result.symbols)
	return result
}

func CreateBaseFromSize(size int) *Base {
	result:=&Base{ size: size, hasSymbols: false }
	return result
}

func (b *Base) Size() int {
	return b.size
}

type BaseConverter struct {
	inputBase, outputBase *Base
}

func NewBC(inputBase, outputBase *Base) BaseConverter {
	return BaseConverter{inputBase: inputBase, outputBase: outputBase}
}

func isAccordingToBase(values []int, base int) bool {
	result:=true
	for i:=range values {
		if value:=values[i]; value >= base || value < 0 {
			result=false
			break
		}
	
	}

	return result
}

func (bc BaseConverter) ToBase10(b Base, values []int) int {
	li:=len(values)-1
	sum:=0
	bsize:=float64(bc.inputBase.Size())
	for i:=range values {
		value:=values[li-i]
		sum+=int(float64(value) * math.Pow(bsize, float64(i)))
	}

	return sum
}

func getExp(sum, base int) float64 {
	return math.Log(float64(sum))/math.Log(float64(base))
}

func getConst(sum, base int, exp float64) int {
	iexp:=math.Trunc(exp)
	return sum / int(math.Pow(float64(base), iexp))
}

func getValue(cnst, base int, exp float64) int {
	return cnst * int(math.Pow(float64(base), exp))
}

func getConstAndExp(sum int, base Base) (int, int) {
	b:=base.Size()

	fexp:=getExp(sum, b)
	if fexp >= 1 {
		cnst:=getConst(sum, b, fexp)
		return cnst, int(math.Trunc(fexp))
	}

	return sum, 0
}

func (bc BaseConverter) FromBase10(sum int, base Base) []int {
	var result []int

	b:=float64(base.Size())
	for {
		c, i:=getConstAndExp(sum, base)
		if len(result) < i+1 {
			result=make([]int, i+1)
		}
		index:=len(result) - (i+1)
		result[index]=c

		sum-=getValue(c, int(b), float64(i))
		if sum == 0 {
			break
		}
	}

	return result
}

func (bc BaseConverter) Convert(values []int) []int {
	sum:=bc.ToBase10(*bc.inputBase, values)
	result:=bc.FromBase10(sum, *bc.outputBase)

	return result
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	var result []int
	var err error
	if (inputBase < 2) {
		err=errors.New("input base must be >= 2")
		return result, err
	}

	if (outputBase < 2) {
		err=errors.New("output base must be >= 2")
		return result, err
	}

	if !isAccordingToBase(inputDigits, inputBase) {
		err=errors.New("all digits must satisfy 0 <= d < input base")
		return result, err
	}

	ib, ob:=CreateBaseFromSize(inputBase), CreateBaseFromSize(outputBase)
	if ib != nil && ob != nil {
		bc:=NewBC(ib, ob)
		result=bc.Convert(inputDigits)
	}

	return result, err
}

// 	CreateBaseFromSymbols([]rune{'0', '1'}),
// 	CreateBaseFromSymbols([]rune{'0','1','2','3','4','5','6','7','8','9'}),
// 	CreateBaseFromSymbols([]rune{'0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F'}),
// 	CreateBaseFromSymbols([]rune{'0','1','2','3','4','5','6'}),
// 	CreateBaseFromSymbols([]rune{'0','1','2'}),
