package main

import (
	"fmt"
	"phonenumber"
	"regexp"
)

func main() {
	telephonenumber := []string{"(223) 456-7890", "223.456.7890", "223 456   7890   ", "123456789",
		"22234567890", "12234567890", "+1 (223) 456-7890", "321234567890",
		"523-abc-7890", "523-@:!-7890", "(023) 456-7890", "(123) 456-7890",
		"(223) 056-7890", "(223) 156-7890", "1 (023) 456-7890", "1 (123) 456-7890",
		"1 (223) 056-7890", "1 (223) 156-7890"}

	re := regexp.MustCompile(`((\d)|(\d\D+))|((\d+)(\d+\D+))`)
	//re2 := regexp.MustCompile(`(\(\d{3}\))`)

	for index := 0; index < len(telephonenumber); index++ {
		result := re.FindAllString(telephonenumber[index], -1)
		fmt.Println(len(result), "->", result)
		data, _ := phonenumber.Format(telephonenumber[index])
		area := data
		exchageCode := data
		number := data
		fmt.Println(area)
		fmt.Println(exchageCode)
		fmt.Println(number)
	}
}
