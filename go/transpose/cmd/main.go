package main

import (
	"fmt"
	"transpose"
)

func main() {
	t := transpose.New([]string{
		"T", "EE", "AAA", "SSSS", "EEEEE", "RRRRRR"})

	data := t.Juxtapose()
	response := []string{"TEASER", " EASER", "  ASER", "   SER", "    ER", "     R"}
	for i := range data {
		fmt.Printf("[%v] <=> [%v]\n", response[i], data[i])
	}
}
