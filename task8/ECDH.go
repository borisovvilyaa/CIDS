package main

import (
	cids "cids/task7"
	"fmt"
	"math/big"
	"math/rand"
)

type KeyPair struct {
	PrivateKey big.Int
	PublicKey  cids.ECPoint
}

func generatePrivateKey(curveOrder int) *big.Int {

	return big.NewInt(rand.Int63n(int64(curveOrder)))
}


func computeSharedSecret(privateKey big.Int, publicKey cids.ECPoint) cids.ECPoint {
	return cids.ScalarMult(privateKey, publicKey)
}

func printUserDetails(user int, keyPair KeyPair) {
	fmt.Printf("User %d:\n", user)
	fmt.Printf("Private Key: %v\n", keyPair.PrivateKey.String())
	fmt.Printf("Public Key: (%v, %v)\n\n", keyPair.PublicKey.X, keyPair.PublicKey.Y)
}

func main() {
	//Config for generation privkey
	basePoint := cids.BasePointGGet()
	curveOrder := 256

	//first key
	privateKey1 := generatePrivateKey(curveOrder)
	keyPair1 := KeyPair{
		PrivateKey: *privateKey1,
		PublicKey:  computeSharedSecret(*privateKey1, basePoint),
	}

	//second key
	privateKey2 := generatePrivateKey(curveOrder)
	keyPair2 := KeyPair{
		PrivateKey: *privateKey2,
		PublicKey:  computeSharedSecret(*privateKey2, basePoint),
	}

	printUserDetails(1, keyPair1)
	printUserDetails(2, keyPair2)

	sharedSecret1 := computeSharedSecret(keyPair1.PrivateKey, keyPair2.PublicKey)
	sharedSecret2 := computeSharedSecret(keyPair2.PrivateKey, keyPair1.PublicKey)

	if cids.IsEqual(sharedSecret1, sharedSecret2) {
		fmt.Println("Shared secret successfully computed:")
		fmt.Printf("A common secret: %v\n", sharedSecret1.X)
	} else {
		fmt.Println("Error: Shared secrets do not match")
	}
}
