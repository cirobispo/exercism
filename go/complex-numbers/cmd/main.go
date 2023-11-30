package main

import (
	"complexnumbers"
	"fmt"
	"math"
)

func main () {
	fmt.Println( complexnumbers.Polar2Rect(complexnumbers.Number{0, math.Pi}) )
}