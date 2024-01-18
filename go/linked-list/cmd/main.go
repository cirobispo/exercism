package main

import (
	"fmt"
	"linkedlist"
)

func main() {
	list:=linkedlist.NewList(1,2,3,4,5)
	Print(list)
	list.Reverse()
	Print(list)
}

func Print(list *linkedlist.List) {
	node:=list.First()
	for {
		fmt.Printf("%d,\t", node.Value.(int))
		node=node.Next()
		if node == nil {
			break
		}
	}
	fmt.Println()
}