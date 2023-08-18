package main

import "fmt"

func main() {
	var bom_dia = "Bom dia!"
	var addr *string = &bom_dia
	var addr_addr *string = addr
	fmt.Printf("addr: %v, addr_addr: %v\n", addr, addr_addr)

	fmt.Println("bom_dia->addr_addr:", *addr_addr)
	var boa_noite string = "Boa noite!"
	addr_addr = &boa_noite
	fmt.Println("boa_noite->addr_addr:", *addr_addr)

	boa_noite = "não sou mais boa noite!"
	fmt.Println("addr_addr:", *addr_addr)
	fmt.Println("boa_noite:", boa_noite)
	fmt.Println("addr:", *addr)

	fmt.Println("this is the main entry point.")

}
