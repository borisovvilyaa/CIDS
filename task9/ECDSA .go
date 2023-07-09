package main

import (
	cids "cids/task7"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
)

type KeyPair struct {
	PrivateKey int
	PublicKey  cids.ECPoint
}

type signature struct {
	firstElement  int
	secondElement cids.ECPoint
}

func generatePrivateKey(curveOrder int) int {
	for {
		n := rand.Intn(curveOrder)
		if n != 0 {
			return rand.Intn(curveOrder)
		}
	}
}

func computePublicKey(privateKey int, basePoint cids.ECPoint) cids.ECPoint {
	return cids.ScalarMult(privateKey, basePoint)
}

func printUserDetails(user int, keyPair KeyPair) {
	fmt.Printf("User %d:\n", user)
	fmt.Printf("Private Key: %d\n", keyPair.PrivateKey)
	fmt.Printf("Public Key: (%f, %f)\n\n", keyPair.PublicKey.X, keyPair.PublicKey.Y)
}

func hashData(input string) [32]byte {
	bytes := []byte(input)
	hash := sha256.Sum256(bytes)
	return hash
}
func modInverse(a, m int) int {
	if m == 0 {
		return 0
	}

	g := gcdExtended(a, m)
	if g < 0 {
		g += m
	}

	return g
}

func gcdExtended(a, b int) int {
	if a == 0 {
		return b
	}

	gcd := gcdExtended(b%a, a)
	x := gcdExtended(b%a, a)

	return x - (b/a)*gcd
}

func HashToNumber(hash []byte) uint64 {
	// Хеш SHA256 имеет размер 32 байта (256 бит)
	// Мы будем преобразовывать первые 8 байт хеша в число типа uint64
	number := binary.BigEndian.Uint64(hash[:8])
	return number
}

func main() {
	countPoint := 0
	basePoint := cids.BasePointGGet()
	countPoint++
	curveOrder := 256

	//1 step. Create key pair
	privateKey := generatePrivateKey(curveOrder)
	keyPair := KeyPair{
		PrivateKey: privateKey,
		PublicKey:  computePublicKey(privateKey, basePoint),
	}
	countPoint++
	printUserDetails(1, keyPair)

	// Signature step.

	input := "Hello, world!"
	hash := hashData(input)

	fmt.Println("Input string:", input)
	fmt.Printf("SHA256 hash:%s\n\n", hex.EncodeToString(hash[:]))

	nonce := rand.Intn(100)
	fmt.Printf("Nonce: %d\n\n", nonce)

	point_multiplication := cids.ScalarMult(nonce, cids.BasePointGGet())

	r := point_multiplication.X

	fmt.Printf("Point multiplication: %f\n\n", point_multiplication)
	countPoint++

	k_inverse := modInverse(nonce, countPoint)
	number := HashToNumber(hash[:])
	fmt.Println(number, keyPair.PrivateKey, r, k_inverse)
	s := (int(number)*keyPair.PrivateKey + int(r)*k_inverse) % countPoint
	fmt.Printf("Signature: %f, %d\n\n", r, s)

	//verify signature

	s_inverse := modInverse(s, curveOrder)
	w := s_inverse % curveOrder
	u1_int := (int(HashToNumber(hash[:])) * w) % curveOrder
	u2_int := ((int(r) * w) % curveOrder)
	fmt.Println(u1_int, u2_int)
	u1 := cids.ScalarMult(u1_int, cids.BasePointGGet())
	u2 := cids.ScalarMult(u2_int, cids.BasePointGGet())

	fmt.Println(u1, u2)
	sumU := cids.AddECPoints(u1, u2)

	fmt.Printf("%v, %f\n", sumU.X, r)

	// verify signature
	fmt.Print("\n\nVerifying signature\n\n")
	s_inverse = modInverse(s, curveOrder)
	w = s_inverse % curveOrder
	u1_int = (int(HashToNumber(hash[:])) * w) % curveOrder
	u2_int = ((int(r) * w) % curveOrder)
	fmt.Println(u1_int, u2_int)
	u1 = cids.ScalarMult(u1_int, cids.BasePointGGet())
	u2 = cids.ScalarMult(u2_int, cids.BasePointGGet())

	fmt.Println(u1, u2)

	fmt.Println(u1, u2)
	point_sum := cids.AddECPoints(u1, u2)
	fmt.Print(r == point_sum.X, r, point_sum.X)
}

// func hashData(input string) [32]byte {
// 	bytes := []byte(input)
// 	hash := sha256.Sum256(bytes)
// 	return hash
// }

// func generateBlindingPoint(nonce int, basePoint cids.ECPoint) cids.ECPoint {
// 	return cids.ScalarMult(nonce, basePoint)
// }

// func generateSignature(privateKey, nonce int, hash [32]byte, blindingPoint cids.ECPoint, countPoint int) signature {
// 	firstComp := (int(blindingPoint.X) + privateKey) % countPoint
// 	secondComp := cids.ScalarMult(int(binary.BigEndian.Uint32(hash[:4]))/(privateKey)/(nonce), blindingPoint)
// 	return signature{
// 		firstElement:  firstComp,
// 		secondElement: secondComp,
// 	}
// }

// func 	verifySignature(input string, hash [32]byte, signature signature, keyPair KeyPair, countPoint int) bool {
// 	Q := cids.AddECPoints(cids.ScalarMult(int(signature.secondElement.Y), cids.BasePointGGet()), cids.ScalarMult(signature.firstElement, keyPair.PublicKey))
// 	fmt.Println(signature.firstElement, int(Q.X))
// 	return signature.firstElement == int(Q.X)
// }

// func main() {
// 	countPoint := 0
// 	basePoint := cids.BasePointGGet()
// 	countPoint++
// 	curveOrder := 256

// 	privateKey := generatePrivateKey(curveOrder)
// 	keyPair := KeyPair{
// 		PrivateKey: privateKey,
// 		PublicKey:  computePublicKey(privateKey, basePoint),
// 	}
// 	countPoint++

// 	printUserDetails(1, keyPair)

// 	input := "Hello, world!"
// 	hash := hashData(input)

// 	fmt.Println("Input string:", input)
// 	fmt.Printf("SHA256 hash:%s\n\n", hex.EncodeToString(hash[:]))

// 	nonce := rand.Intn(100)
// 	fmt.Printf("Nonce: %d\n\n", nonce)

// 	blindingPoint := generateBlindingPoint(nonce, basePoint)
// 	countPoint++
// 	fmt.Printf("Blinding point: (%f, %f)\n\n", blindingPoint.X, blindingPoint.Y)

// 	fmt.Printf("Count Point: %d\n\n", countPoint)

// 	signature := generateSignature(privateKey, nonce, hash, blindingPoint, countPoint)
// 	fmt.Printf("First Component: %d\n\n", signature.firstElement)
// 	fmt.Println("Second Component:", signature.secondElement)

// 	fmt.Printf("\n\n-------Verify-------\n\n")

// 	inputverify := "Hello, world!"
// 	hashverify := hashData(input)

// 	fmt.Println("Input string:", inputverify)
// 	fmt.Printf("SHA256 hash:%s\n\n", hex.EncodeToString(hashverify[:]))

// 	fmt.Println(verifySignature(input, hashverify, signature, keyPair, countPoint))
// }
