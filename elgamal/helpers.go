package elgamal

import (
	"crypto/sha256"
	"math/big"
)

func FromHex(s string) *big.Int {
	if s == "" {
		return big.NewInt(0)
	}
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid hex in source file: " + s)
	}
	return r
}

// returns H(m) in big.Int | sha256 hash
func Digest(m []byte) *big.Int {
	h := sha256.New()
	h.Write(m)
	e := big.NewInt(0)
	e.SetBytes(h.Sum(nil)) // e = Hash(msg) (in bigInt format)
	return e
}
