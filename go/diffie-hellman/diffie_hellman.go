package diffiehellman

import (
	crand "crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	var max, k big.Int
	max.Sub(p, big.NewInt(3))
	r, err := crand.Int(crand.Reader, &max)
	if err != nil {
		panic(err)
	}
	return k.Add(r, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	var result *big.Int = big.NewInt(0)
	result.Exp(big.NewInt(g), private, p)
	return result
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	pvt_k:=PrivateKey(p)
	pub_k:=PublicKey(pvt_k, p, g)

	return pvt_k, pub_k
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	var result *big.Int = big.NewInt(0)
	result.Exp(public2, private1, p)
	return result
}
