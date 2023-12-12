package elgamal

import (
	"math/big"
)

type Cypher struct {
	x *big.Int
	y *big.Int
}

func NewCypher(xVal, yVal *big.Int) *Cypher {
	return &Cypher{
		x: xVal,
		y: yVal,
	}
}

/*
m - message to encrypt
b - public key of message recipient
*/
func Encrypt(m []byte, b *big.Int) []*Cypher {
	blocks := MakeBlocks(m)
	cyphers := make([]*Cypher, len(blocks))

	for i := 0; i < len(blocks); i++ {
		x, y := big.NewInt(0), big.NewInt(0)
		g, p := Params().g, Params().p
		k := PrivateGenerator()

		m_block := big.NewInt(0)
		m_block.SetBytes(m)

		x.Exp(g, k, p)
		y.Exp(b, k, p)
		y.Mul(y, m_block)
		y.Mod(y, p)

		cypherBlock := NewCypher(big.NewInt(0), big.NewInt(0))
		cypherBlock.x = x
		cypherBlock.y = y

		cyphers[i] = cypherBlock
	}

	// x, y := big.NewInt(0), big.NewInt(0)
	// g, p := Params().g, Params().p
	// k := PrivateGenerator()

	// m_block := big.NewInt(0)
	// m_block.SetBytes(m)

	// x.Exp(g, k, p)
	// y.Exp(b, k, p)
	// y.Mul(y, m_block)
	// y.Mod(y, p)

	return cyphers
}

// general decrypt function
func Decrypt(cyphers []*Cypher, a *big.Int) []byte {
	var decrypted_data []byte
	for i := 0; i < len(cyphers); i++ {
		cypherBlock := cyphers[i]
		decryptedBlock := cypherBlock.Decrypt(a)
		decrypted_data = append(decrypted_data, decryptedBlock...)
	}
	return decrypted_data
}

// structure method
func (cypher *Cypher) Decrypt(a *big.Int) []byte {
	p := Params().p

	s := big.NewInt(0)
	s.Exp(cypher.x, a, p)

	s_inverse := big.NewInt(0)
	s_inverse.ModInverse(s, p)

	m := big.NewInt(0)
	m.Mul(cypher.y, s_inverse)
	m.Mod(m, p)
	//fmt.Println(message_block)
	return m.Bytes()
}

// split message into the blocks
func MakeBlocks(m []byte) []*big.Int {
	blockSize := Params().BitSize()
	bitsInMessage := len(m) * 8
	//println(bitsInMessage)
	blocksRequried := 0
	if bitsInMessage%blockSize == 0 {
		blocksRequried = bitsInMessage / blockSize
	} else {
		blocksRequried = (bitsInMessage / blockSize) + 1
	}
	blocks := make([]*big.Int, blocksRequried)

	offset := 0
	length := blockSize / 8

	for i := 0; i < len(blocks); i++ {
		if i == len(blocks)-1 {
			length = len(m)
		}

		currBlock := big.NewInt(0)
		blocks[i] = currBlock.SetBytes(m[offset:length])
		offset = length
		length = offset + (blockSize / 8)
	}

	return blocks
}
