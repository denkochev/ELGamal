package elgamal

import (
	"math/big"
)

type KeyPair struct {
	private *big.Int
	public  *big.Int
}

func (pair *KeyPair) GenerateKeyPair() {
	a := PrivateGenerator() // random in [1, p-1]

	b := big.NewInt(0)
	b.Exp(Params().g, a, Params().p) // g^a mod p

	pair.private = a
	pair.public = b
}

func (pair *KeyPair) GetPublic() *big.Int {
	return pair.public
}
