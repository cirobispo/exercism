package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type matchCallback func(string) bool
type validatorCommand func(string) bool

type stringValidator interface {
	Regex() string
	Match(string) bool
	Validate(string) bool
	setNext(n stringValidator)
}

type validator struct {
	regex   string
	command validatorCommand
	next    stringValidator
}

func (v *validator) Regex() string {
	return v.regex
}

func (v *validator) Match(data string) bool {
	r := regexp.MustCompile(v.regex)
	match := r.Match([]byte(data))
	return match
}

func (v *validator) setCommand(cmd validatorCommand) {
	v.command = cmd
}

func (v *validator) setNext(n stringValidator) {
	v.next = n
}

func (v *validator) Validate(id string) bool {
	isValid := v.Match(id)
	if isValid {
		command := v.command
		isCard := command(id)
		return isCard
	} else {
		if v.next != nil {
			return v.next.Validate(id)
		}
		return false
	}
}

func New(regex string) *validator {
	return &validator{regex: regex}
}

func calcNumbers(id string) int {
	var mult int = 2
	var result int = 0
	for i := len(id) - 1; i > -1; i-- {
		mult += int(math.Pow(-1, float64(len(id)-i)))
		//		fmt.Printf("\t%v", float64(len(id)-i))
		value, _ := strconv.Atoi(string(id[i]))
		value = int(float64(value) * math.Abs(float64(mult)))
		if value > 9 {
			value -= 9
		}
		result += value
	}
	return result
}

func isCreditCardValid(id string) bool {
	calc_creditcard := func(id string) int {
		var amount int = 0

		blocks := strings.Split(id, " ")
		for i := len(blocks) - 1; i > -1; i-- {
			amount += calcNumbers(blocks[i])
		}
		// fmt.Printf("%v", amount)
		return amount
	}

	var value int = -1
	value = calc_creditcard(id)

	fmt.Printf("\nCREDITCARD valor: %v, valido: %v", value, value%10 == 0)
	return (value%10 == 0)
}

func isCanadianSIMValid(id string) bool {
	calc_sim := func(id string) int {
		var amount int = 0

		blocks := strings.ReplaceAll(id, " ", "")
		amount += calcNumbers(blocks)
		//fmt.Printf("%v", amount)
		return amount
	}

	value := calc_sim(id)

	fmt.Printf("\nCANADIAN SIN - valor: %v, valido: %v", value, value%10 == 0 && value > 0)
	return (value%10 == 0 && value > 0)
}

func isNotNumber(id string) bool {
	fmt.Printf("\nNão possui somente digitos - valor: %v, false", id)
	return false
}

func unknownData(id string) bool {
	fmt.Printf("\nDado não possui formato conhecido - valor: %v, false", id)
	return false
}

func main() {
	notNumber := New("[\\D]")
	notNumber.setCommand(isNotNumber)

	//	cc_validator := New("(^\\d{3}\\s\\d{3}\\s\\d{3}\\s\\d{3}$)")
	cc_validator := New("(\\d{3})")
	cc_validator.setCommand(isCreditCardValid)

	//	cc_validator1 := New("(^\\d{4}\\s\\d{4}\\s\\d{4}\\s\\d{4}$)")
	cc_validator1 := New("(\\d{4}$)")
	cc_validator1.setCommand(isCreditCardValid)

	//	canadian_sim := New("(^\\d{3}\\s\\d{3}\\s\\d{3}$)")
	canadian_sim := New("(\\d{3})")
	canadian_sim.setCommand(isCanadianSIMValid)

	sin := New("(\\d{1,2})")
	sin.setCommand(isCanadianSIMValid)

	unknown := New("([0-9])")
	unknown.setCommand(unknownData)

	notNumber.setNext(cc_validator)
	cc_validator.setNext(cc_validator1)
	cc_validator1.setNext(canadian_sim)
	canadian_sim.setNext(sin)
	sin.setNext(unknown)

	allData := []string{"1", "0", "059", "59", "055 444 285", "8273 1232 7352 0569",
		"1 2345 6789 1234 5678 9012", "1 2345 6789 1234 5678 9013", "095 245 88",
		"234 567 891 234", "059a", "055-444-285", "055# 444$ 285", " 0", "0000 0", "091",
		"9999999999 9999999999 9999999999 9999999999", "109", "055b 444 285", ":9",
		"59%59"}
	// f f v v v f f f v v f f f f v v v v f f f
	for _, item := range allData {
		data := strings.ReplaceAll(item, " ", "")
		notNumber.Validate(data)
	}
}
