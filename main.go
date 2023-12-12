package main

import (
	"elgamal/elgamal"
	"fmt"
)

func main() {
	message := "Modern programming languages are fast. Call me Ishmael. Some years ago-never mind how long precisely-having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. Call me Ishmael. Some years ago-never mind how long precisely-having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. Modern programming languages are fast. "
	fmt.Println(len([]byte(message)) * 8)
	signature, keypair := elgamal.UnstableSignature(message)

	fmt.Println(signature.Verify([]byte(message), keypair.GetPublic()))

	alice_keyPair := elgamal.KeyPair{}
	alice_keyPair.GenerateKeyPair()

	make_cypher := elgamal.Encrypt([]byte(message), alice_keyPair.GetPublic())
	plain_data := elgamal.Decrypt(make_cypher, alice_keyPair.GetPrivate())

	parseToText := string(plain_data)
	fmt.Println(parseToText)
}
