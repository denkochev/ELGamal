package main

import (
	"elgamal/elgamal"
	"fmt"
	"math/big"
)

func main() {

	x := big.Int{}
	x.SetBytes(elgamal.MakeRandom("params"))

	fmt.Println(x.Text(10))
}
