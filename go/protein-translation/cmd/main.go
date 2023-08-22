package main

import (
	"fmt"
	"protein"
)

func main() {
	mm := protein.NewMultiMap()
	mm.Add([]string{":---"}, ":---")
	mm.Add([]string{"AUG"}, "Methionine")
	mm.Add([]string{"UUU", "UUC"}, "Phenylalanine")
	mm.Add([]string{"UUA", "UUG"}, "Leucine")
	mm.Add([]string{"UCU", "UCC", "UCA", "UCG"}, "Serine")
	mm.Add([]string{"UAU", "UAC"}, "Tyrosine")
	mm.Add([]string{"UGU", "UGC"}, "Cysteine")
	mm.Add([]string{"UGG"}, "Tryptophan")
	mm.Add([]string{"UAA", "UAG", "UGA"}, "STOP")

	mm.Add([]string{"AUG"}, "START")

	fmt.Println(mm.GetValue("UGC"))
	fmt.Println(mm.GetValue("UCA"))

	fmt.Println(mm.GetKeys("STOP"))
}
