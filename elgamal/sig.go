package elgamal

import "math/big"

type Signature struct {
	r *big.Int
	s *big.Int
}

func (signature *Signature) Sign(m []byte, a *big.Int) {
	k := PrivateGenerator()

	r := big.NewInt(0)
	r.Exp(Params().g, k, Params().p) //  r = g^k mod p

	e := Digest(m) // H(m)

	inv_k := big.NewInt(0)
	inv_k.ModInverse(k, Params().p) // k^-1

	ar := big.NewInt(0)
	ar.Mul(a, r) // a*r

	sig_part_1 := big.NewInt(0)
	sig_part_1.Sub(e, ar) // (H(m) - a*r)

	s := big.NewInt(0)
	s.Mul(sig_part_1, inv_k)   // (H(m) - a*r) * k^(-1)
	s.Mod(s, Params().p_sub_1) // s = (H(m) - a*r) * k^(-1) mod (p-1)

	signature.r = r
	signature.s = s
}
