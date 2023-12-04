package elgamal

import "math/big"

type GamalParams struct {
	p *big.Int
	g *big.Int
}

// params for group
var gamalParams = GamalParams{}

// getter for group
func Params() *GamalParams {
	return &gamalParams
}
