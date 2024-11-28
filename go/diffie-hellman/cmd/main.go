package main

import (
	"diffiehellman"
	"fmt"
	"math/big"
)

func main () {
	t:=diffiehellman.PrivateKey(big.NewInt(4))
	diffiehellman.PublicKey(big.NewInt(4), t, t.Int64())
	fmt.Println(t)
}