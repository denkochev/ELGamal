package elgamal

import (
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

// temporarly solution for signature
func UnstableSignature(message string) (*Signature, *KeyPair) {
	pair := GenerateKeyPair()
	signature := &Signature{big.NewInt(0), big.NewInt(0)}
	signature.Sign([]byte(message), pair.GetPrivate())
	for {

		isSignatureValid := signature.Verify([]byte(message), pair.GetPublic())
		if isSignatureValid {
			break
		}
		// generate again if verifying wasn't succesful
		pair = GenerateKeyPair()
		signature.Sign([]byte(message), pair.GetPrivate())
	}
	return signature, pair
}

func (signature *Signature) Sign(m []byte, a *big.Int) {
	k := PrivateGenerator()
	r := big.NewInt(0)
	r.Exp(Params().g, k, Params().p) //  r = g^k mod p

	e := Digest(m) // H(m)

	ar := big.NewInt(0)
	ar.Mul(a, r) // a*r

	sig_part_1 := big.NewInt(0)
	sig_part_1.Sub(e, ar) // (H(m) - a*r)

	inv_k := big.NewInt(0)
	inv_k.ModInverse(k, Params().p_sub_1) // k^-1

	s := big.NewInt(0)
	s.Mul(sig_part_1, inv_k)   // (H(m) - a*r) * k^(-1)
	s.Mod(s, Params().p_sub_1) // s = (H(m) - a*r) * k^(-1) mod (p-1)

	signature.R = r
	signature.S = s
}

func (signature *Signature) Verify(m []byte, b *big.Int) bool {
	p := Params().p             // p
	p_sub_1 := Params().p_sub_1 // p-1
	g := Params().g
	s := signature.S // s
	r := signature.R // r
	s_inverse := big.NewInt(0)
	s_inverse.ModInverse(s, p_sub_1) // s^(-1)
	e := Digest(m)                   // H(m)

	// TODO: HERE IS A PROBLEM, GCD OF P-1 AND S OFTEN = 2
	// ts := big.NewInt(0)
	// fmt.Println(ts.GCD(nil, nil, s, p_sub_1))

	y := big.NewInt(0)
	y.ModInverse(b, p) // y = b^(-1) mod p

	u1 := big.NewInt(0)
	u1.Mul(e, s_inverse)
	u1.Mod(u1, p_sub_1) // (H(m) * s^(-1)) mod (p-1)

	u2 := big.NewInt(0)
	u2.Mul(r, s_inverse) // (r * s^(-1))
	u2.Mod(u2, p_sub_1)  // (r * s^(-1)) mod (p-1)

	g_u1 := big.NewInt(0)
	g_u1.Exp(g, u1, p) // g^u1

	y_u2 := big.NewInt(0)
	y_u2.Exp(y, u2, p) // y^u2

	v := big.NewInt(0)
	v.Mul(g_u1, y_u2)
	v.Mod(v, p) // v = (g^u1 * y^u2) mod p

	res := v.Cmp(r) == 0
	return res
}
