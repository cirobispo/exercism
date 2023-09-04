package main

import (
	"cipher"
	"fmt"
)

func main() {
	ceasar := cipher.NewCaesar()

	fmt.Println(ceasar.Encode("iamapandabear"))
}
