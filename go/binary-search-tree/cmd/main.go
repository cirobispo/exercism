package main

import (
	"binarysearchtree"
	"fmt"
)

func main() {
	b:=binarysearchtree.NewBst(4)
	b.Insert(2)
	b.Insert(6)
	b.Insert(3)
	b.Insert(1)
	b.Insert(5)
	b.Insert(7)
	
	fmt.Println(b)
}