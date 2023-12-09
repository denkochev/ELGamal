package main

import (
	"elgamal/elgamal"
	"fmt"
)

func main() {
	message := "Modern programming langugages are fast"
	signature, keypair := elgamal.UnstableSignature(message)

	fmt.Println(signature.Verify([]byte(message), keypair.GetPublic()))

	// pair := elgamal.KeyPair{}
	// pair.GenerateKeyPair()

	// message := "Modern programming langugage are fast"

	// signature := elgamal.Signature{}
	// signature.Sign([]byte(message), pair.GetPrivate())

	// isSignatureValid := signature.Verify([]byte(message), pair.GetPublic())
	// fmt.Println(isSignatureValid)
}
