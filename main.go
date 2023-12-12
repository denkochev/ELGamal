package main

import (
	"elgamal/elgamal"
	"fmt"
)

func main() {
	message := "Modern programming langugages are fast"
	signature, keypair := elgamal.UnstableSignature(message)

	fmt.Println(signature.Verify([]byte(message), keypair.GetPublic()))

	alice_keyPair := elgamal.KeyPair{}
	alice_keyPair.GenerateKeyPair()

	make_cypher := elgamal.Encrypt([]byte(message), alice_keyPair.GetPublic())
	plain_data := elgamal.Decrypt(make_cypher, alice_keyPair.GetPrivate())

	parseToText := string(plain_data)
	fmt.Println(parseToText)
}
