package elgamal

import (
	"fmt"
	"math/rand"
)

func MakeRandom(mode string) []byte {
	var bitLength int

	if mode == "params" {
		min := 2048
		max := 4096

		bitLength = rand.Intn(max-min) + min // range is min to max
	}

	var rand_set []byte
	fmt.Println(bitLength)
	// generate slice for random sequence
	if bitLength%8 == 0 {
		rand_set = make([]byte, bitLength/8)
	} else {
		rand_set = make([]byte, bitLength/8+1)
	}
	fmt.Println(len(rand_set))
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
			var tempByte byte

			for i := 0; i < bitToGenerate; i++ {
				tempByte = tempByte << 1
				curRandBit := byte(rand.Intn(2))
				tempByte = tempByte | curRandBit
			}
			// fmt.Printf("%08b\n", tempByte)
			rand_set[i] = tempByte
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

	return rand_set
}
