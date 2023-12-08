package elgamal

import "math/big"

type GamalParams struct {
	// safe p prime number
	p *big.Int
	// p-1 value
	p_sub_1 *big.Int
	//
	g *big.Int
	// bit size of p
	bitSize int
}

// params for group
var gamalParams = GamalParams{
	/*
		A "safe prime" is an integer p
		 which is such that both p
		 and (p−1)/2
		 are prime. If you generate your modulus for ElGamal signatures are a "safe prime",
		 then there are only two prime factors of p−1: these are 2, and (p−1)/2
		.
	*/
	p: FromHex("58b24e74e15b45483db571d78ee01067481713697e4c2c0ef7b069eec3693ab4efc2b13625413c69602e5ac2c80ad84754ba7d9dd4c499d6f1043e9e93d30427aa71d2d6784a4e8e69465318c657dd6f593ae5f948d9b3275b545401ac094ed738cf5ae6bb0f642cbdefba37062b64261f583c34e1024b1a192fe5807df3d55a6fac1c3b8a28a3900b46fb01ce1a2fe9ed6eebad75ff0752594b808cb50e135ba8da2002301f7fd008ec57b9933ec1767ed4ea9384b3ebd74e98b6dc61ff8e52e53eacc55fe9b97dd54f17e7a3dc2b7bb3aee3d870202d83e447216bf3833aef3a5d4e8b1c250569cb8f373f8e5deb2748e4f8227a4ea7122483422da7705578c734683b620e12f4583d03d41369c3c732fdd94d0fc4824ff318ae2d85fbad4c4e1f037c4d102f0b185a76c2d949f3483b16a1711a35c0bc4afe4e75eb1040d7d40ff1a925671e27bdc779fe3d7889e8599a374395d83648cf91a75f5a858b8adf4278537fadc4e91a64432e0007f158073654839eb969a813b60f2f58b2aa15d4b16abfae48f42efa726f197b4dc5a8ef6d05651fffd571e440733405800b2f873c62553b39858ba07b57c481ca233e5d83773124996110f8164c83d4b967c3a0fdd7e05ccfd5e78bbfa771835dad85f2cab4ca3b2b540a3"),

	p_sub_1: FromHex("58b24e74e15b45483db571d78ee01067481713697e4c2c0ef7b069eec3693ab4efc2b13625413c69602e5ac2c80ad84754ba7d9dd4c499d6f1043e9e93d30427aa71d2d6784a4e8e69465318c657dd6f593ae5f948d9b3275b545401ac094ed738cf5ae6bb0f642cbdefba37062b64261f583c34e1024b1a192fe5807df3d55a6fac1c3b8a28a3900b46fb01ce1a2fe9ed6eebad75ff0752594b808cb50e135ba8da2002301f7fd008ec57b9933ec1767ed4ea9384b3ebd74e98b6dc61ff8e52e53eacc55fe9b97dd54f17e7a3dc2b7bb3aee3d870202d83e447216bf3833aef3a5d4e8b1c250569cb8f373f8e5deb2748e4f8227a4ea7122483422da7705578c734683b620e12f4583d03d41369c3c732fdd94d0fc4824ff318ae2d85fbad4c4e1f037c4d102f0b185a76c2d949f3483b16a1711a35c0bc4afe4e75eb1040d7d40ff1a925671e27bdc779fe3d7889e8599a374395d83648cf91a75f5a858b8adf4278537fadc4e91a64432e0007f158073654839eb969a813b60f2f58b2aa15d4b16abfae48f42efa726f197b4dc5a8ef6d05651fffd571e440733405800b2f873c62553b39858ba07b57c481ca233e5d83773124996110f8164c83d4b967c3a0fdd7e05ccfd5e78bbfa771835dad85f2cab4ca3b2b540a2"),

	g: FromHex("2"),

	bitSize: 3780,
}

func Params() *GamalParams {
	return &gamalParams
}

func (params *GamalParams) GetP() *big.Int {
	return params.p
}

func (params *GamalParams) Get_p_sub_1() *big.Int {
	return params.p_sub_1
}

func (params *GamalParams) GetG() *big.Int {
	return params.g
}

func (params *GamalParams) BitSize() int {
	return params.bitSize
}
