package main

import (
	"fmt"
	"summultiples"
)

func main() {
	fmt.Println( callSum(20, []int{3, 5}) )
	fmt.Println( callSum(100, []int{3, 5}) )
	fmt.Println( callSum(150, []int{5, 6, 8}) )
	fmt.Println( callSum(10000, []int{43, 47}) )
	
	// fmt.Println( callSum(10000, []int{2, 3, 5, 7, 11}) )
}

func callSum(limit int, divisors []int) int {
	return summultiples.SumMultiples(limit, divisors...)
}