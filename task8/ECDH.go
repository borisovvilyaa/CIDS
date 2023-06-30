package main

import (
	cids "cids/task7"
	"fmt"
	"math/rand"
)

type KeyPair struct {
	PrivateKey int
	PublicKey  cids.ECPoint
}

func generatePrivateKey(curveOrder int) int {

	return rand.Intn(curveOrder)
}

func computePublicKey(privateKey int, basePoint cids.ECPoint) cids.ECPoint {
	return cids.ScalarMult(privateKey, basePoint)
}

func computeSharedSecret(privateKey int, publicKey cids.ECPoint) cids.ECPoint {
	return cids.ScalarMult(privateKey, publicKey)
}

func printUserDetails(user int, keyPair KeyPair) {
	fmt.Printf("User %d:\n", user)
	fmt.Printf("Private Key: %d\n", keyPair.PrivateKey)
	fmt.Printf("Public Key: (%f, %f)\n\n", keyPair.PublicKey.X, keyPair.PublicKey.Y)
}

func main() {
	//Config for generation privkey
	basePoint := cids.BasePointGGet()
	curveOrder := 256

	//first key
	privateKey1 := generatePrivateKey(curveOrder)
	keyPair1 := KeyPair{
		PrivateKey: privateKey1,
		PublicKey:  computePublicKey(privateKey1, basePoint),
	}

	//second key
	privateKey2 := generatePrivateKey(curveOrder)
	keyPair2 := KeyPair{
		PrivateKey: privateKey2,
		PublicKey:  computePublicKey(privateKey2, basePoint),
	}

	printUserDetails(1, keyPair1)
	printUserDetails(2, keyPair2)


	
	sharedSecret1 := computeSharedSecret(keyPair1.PrivateKey, keyPair2.PublicKey)
	sharedSecret2 := computeSharedSecret(keyPair2.PrivateKey, keyPair1.PublicKey)

	if cids.IsEqual(sharedSecret1, sharedSecret2) {
		fmt.Println("Shared secret successfully computed:")
		fmt.Printf("A common secret: %f\n", sharedSecret1.X)
	} else {
		fmt.Println("Error: Shared secrets do not match")
	}
}
