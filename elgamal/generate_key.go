package elgamal

import (
	"math/big"
)

type KeyPair struct {
	private *big.Int
	public  *big.Int
}

func GenerateKeyPair() *KeyPair {
	a := PrivateGenerator() // random in [1, p-1] | NUMBER HAS TO BE COPRIME WITH P-1 !

	b := big.NewInt(0)
	b.Exp(Params().g, a, Params().p) // g^a mod p
	return &KeyPair{
		private: a,
		public:  b,
	}
}

func (pair *KeyPair) GetPrivate() *big.Int {
	return pair.private
}

func (pair *KeyPair) GetPublic() *big.Int {
	return pair.public
}
