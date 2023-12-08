package elgamal

import (
	"fmt"
	"math/big"
	"math/rand"
)

// generate prime bigInt number
func GenerateBigSafePrime(print bool) (*big.Int, *big.Int) {
	p := big.NewInt(0)
	p_minus_1_div_2 := big.NewInt(0)
	for {
		randomset, bitsize := MakeRandom("params")

		p.SetBytes(randomset)
		res := p.ProbablyPrime(10)

		if res {
			// (p - 1)
			p_minus_1_div_2.Sub(p, big.NewInt(1))
			// (p−1)/2
			p_minus_1_div_2.Div(p_minus_1_div_2, big.NewInt(2))

			res := p_minus_1_div_2.ProbablyPrime(10)
			if res {
				if print {
					fmt.Printf("Random %d bits number\n", bitsize)
					fmt.Println(p.Text(16))
					fmt.Println("----------------------------------------------------------------------------------------")
					fmt.Println(p_minus_1_div_2.Text(16))
				}
				break
			}
		}
	}

	return p, p_minus_1_div_2
}

// generate random (fips140 random) number [2048, 4096] bit length
func MakeRandom(mode string) ([]byte, int) {
	var bitLength int

	if mode == "params" {
		min := 2048
		max := 4096

		bitLength = rand.Intn(max-min) + min // range is min to max
	}

	var rand_set []byte
	// fmt.Println(bitLength)
	// generate slice for random sequence
	if bitLength%8 == 0 {
		rand_set = make([]byte, bitLength/8)
	} else {
		rand_set = make([]byte, bitLength/8+1)
	}
	// fmt.Println(len(rand_set))
	for i := 0; i < len(rand_set); i++ {
		var curByte byte
		/*
			generate []byte slice with exact bitLength random bits,
			if len []byte > bigLength first byte will required amount of zeroes
		*/
		if i == 0 && len(rand_set)*8 > bitLength {
			// fmt.Println("------------------------------")
			bitToGenerate := len(rand_set)*8 - bitLength
			// fmt.Println("bits to generate -> ", bitToGenerate)

			for i := 0; i < bitToGenerate; i++ {
				curByte = curByte << 1
				curRandBit := byte(rand.Intn(2))
				curByte = curByte | curRandBit
			}
			// fmt.Printf("%08b\n", tempByte)
			rand_set[i] = curByte
			continue
		}

		for j := 0; j < 8; j++ {
			curByte = curByte << 1
			curRandBit := byte(rand.Intn(2))
			curByte = curByte | curRandBit
		}
		// fmt.Printf("%08b\n", curByte)

		rand_set[i] = curByte
	}

	return rand_set, bitLength
}
