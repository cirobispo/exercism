package sorting

import (
	"fmt"
	"strconv"
)

func formatFloat(f float64) string {
	value := strconv.FormatFloat(f, 'f', 1, 32)
	return value
}

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %v", formatFloat(f))
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	var value string
	if nb != nil {
		n := nb.Number()
		value = fmt.Sprintf("This is a box containing the number %v", formatFloat(float64(n)))
	}
	return value
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fnbType := fmt.Sprintf("%T", fnb)
	//	fmt.Printf("\nType: %T, type(S): %v, value: %v, value(2): %v \n", fnb, fnbType, fnb, "")
	var result int = 0
	if fnbType == "sorting.FancyNumber" {
		result, _ = strconv.Atoi(fnb.(FancyNumber).n)
	}
	return result
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fnbType := fmt.Sprintf("%T", fnb)
	//fmt.Printf("\nType: %T, type(S): %v, value: %v, value(2): %v \n", fnb, fnbType, fnb, "")
	var result, value string = "", "0.0"
	if fnbType == "sorting.FancyNumber" {
		if v, e := strconv.ParseFloat(fnb.Value(), 64); e == nil {
			value = formatFloat(v)
		}
	}
	result = fmt.Sprintf("This is a %v containing the number %v", "fancy box", value)
	return result
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	var result string = "Return to sender"
	fmt.Printf(fmt.Sprintf("Type: %T\n", i))
	switch i.(type) {
	case int:
		result = DescribeNumber(float64(i.(int)))
		break
	case float64:
		result = DescribeNumber(i.(float64))
		break
	case NumberBox:
		result = fmt.Sprintf("This is a box containing the number %v", formatFloat(float64(i.(NumberBox).Number())))
		break
	case FancyNumber:
		result = fmt.Sprintf("This is a fancy box containing the number %v", formatFloat(float64(ExtractFancyNumber(i.(FancyNumberBox)))))
		break
	case FancyNumberBox:
		result = DescribeFancyNumberBox(i.(FancyNumberBox))
		break
	}
	return result
}
