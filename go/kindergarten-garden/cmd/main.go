package main

import (
	"fmt"
	"kindergarten"
)

func main() {
	//diagram := "\nVRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV"
	//children := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"}

	diagram := "rc\ngg\n"
	children := []string{"Alice", "Alice"}
	g, _ := kindergarten.NewGarden(diagram, children)
	fmt.Println(g.Plants("David"))
}
