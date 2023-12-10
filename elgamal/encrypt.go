package elgamal

import (
	"fmt"
	"math/big"
)

type Cypher struct {
	x *big.Int
	y *big.Int
}

/*
m - message to encrypt
b - public key of message recipient
*/
func Encrypt(m []byte, b *big.Int) Cypher {
	x, y := big.NewInt(0), big.NewInt(0)
	g, p := Params().g, Params().p
	k := PrivateGenerator()

	m_block := big.NewInt(0)
	m_block.SetBytes(m)

	x.Exp(g, k, p)
	y.Exp(b, k, p)
	y.Mul(y, m_block)
	y.Mod(y, p)

	return Cypher{x, y}
}

func (cypher *Cypher) Decrypt(a *big.Int) {
	p := Params().p

	s := big.NewInt(0)
	s.Exp(cypher.x, a, p)

	s_inverse := big.NewInt(0)
	s_inverse.ModInverse(s, p)

	m := big.NewInt(0)
	m.Mul(cypher.y, s_inverse)
	m.Mod(m, p)

	message_block := string(m.Bytes())

	fmt.Println(message_block)

}
