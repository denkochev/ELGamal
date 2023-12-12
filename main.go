package main

import (
	"elgamal/elgamal"
	"fmt"
)

func main() {
	// data to work with, can be any data any bits length in []byte format
	message := "Modern programming languages are fast. Call me Ishmael. Some years ago-never mind how long precisely-having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. Call me Ishmael. Some years ago-never mind how long precisely-having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. Modern programming languages are fast. "
	fmt.Println(len([]byte(message)) * 8)
	// create an instance of signature, UnstableSignature returns signature struct and key pair struct we can work with
	signature, keypair := elgamal.UnstableSignature(message)
	// verify signature struct with public key
	fmt.Println(signature.Verify([]byte(message), keypair.GetPublic()))
	// get keypair structure with private and public bigInt
	alice_keyPair := elgamal.GenerateKeyPair()

	/*
		IMPORTANT: to encrypt data in ElGamal scheme we have to get public Key of recipient
		to decrypt data recipient has to use encrypted data and his own private key
	*/

	// cypher our data through Encrypt function that returns list of Cypher structs
	make_cypher := elgamal.Encrypt([]byte(message), alice_keyPair.GetPublic())
	// decrypt data using decrypt function that takes *[]Cypher and private key of recipient
	plain_data := elgamal.Decrypt(make_cypher, alice_keyPair.GetPrivate())

	// check decrypted data
	parseToText := string(plain_data)
	fmt.Println(parseToText)
}
