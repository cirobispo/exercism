package twelve

import (
	"fmt"
	"strings"
)

func Verse(i int) string {
	verses:=make([]string, 13)
	verses[0]=""
	verses[1]="Partridge in a Pear Tree"
	verses[2]="Turtle Doves"	
	verses[3]="French Hens"	
	verses[4]="Calling Birds"	
	verses[5]="Gold Rings"	
	verses[6]="Geese-a-Laying"	
	verses[7]="Swans-a-Swimming"	
	verses[8]="Maids-a-Milking"	
	verses[9]="Ladies Dancing"	
	verses[10]="Lords-a-Leaping"	
	verses[11]="Pipers Piping"	
	verses[12]="Drummers Drumming"

	glue:=""
	result:=""
	for j:=i;j > 0;j-- {
		result=result + glue + fmt.Sprintf("%s %s", num2text(number(), j, true), verses[j])
		glue=", "
		if j == 2 {
			glue=", and "
		} 
	}
	result=fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", num2text(numberth(), i, true), result)
	return result

	return ""
}

func Song() string {
	result:=Verse(1)
	for i:=2; i<=12; i++ {
		result+=fmt.Sprintf("\n%s", Verse(i));
	}

	return result
}

func number() []string {
	n_value:=[]string{"", "A", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve"}
	return n_value
}

func numberth() []string {
	nth_value:=[]string{"", "First", "Second", "Third", "Fourth", "Fifth", "Sixth", "Seventh", "Eighth", "Ninth", "Tenth", "Eleventh", "Twelfth"}
	return nth_value
}

func num2text(textValues []string, value int, lowerCase bool) string {
	result:= textValues[value]

	if lowerCase{
		return strings.ToLower(result)
	}

	return result
}