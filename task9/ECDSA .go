package main

import (
	cids "cids/task7"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
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
	basePoint := cids.BasePointGGet()
	curveOrder := cids.CurveOrderNGet()

	p := 114973

	//1 step. Create key pair
	privateKey := generatePrivateKey(curveOrder - 1)
	keyPair := KeyPair{
		PrivateKey: privateKey,
		PublicKey:  computePublicKey(privateKey, basePoint),
	}
	printUserDetails(1, keyPair)

	// Signature step.

	input := "Hello, world!"
	hash := hashData(input)
	number := HashToNumber(hash[:])
	fmt.Println("Input string:", input)
	fmt.Printf("SHA256 hash number:%d\n", number)
	fmt.Printf("SHA256 hash:%s\n\n", hex.EncodeToString(hash[:]))

	nonce := rand.Intn(cids.CurveOrderNGet())
	fmt.Printf("Nonce: %d\n\n", nonce)

	point_multiplication := cids.ScalarMult(nonce, cids.BasePointGGet())
	fmt.Printf("Point multiplication: %f\n\n", point_multiplication)

	r := int(point_multiplication.X) % p
	s := int(math.Pow(float64(nonce), -1)*float64(int(number)+int(keyPair.PrivateKey)*r)) % p
	fmt.Println(r, s)

	//verify signature

	w := int(math.Pow(float64(s), -1)) % p
	u1 := (int(number) * w) % p
	u2 := r * w % p
	x1 := cids.ScalarMult(u1, cids.BasePointGGet())
	x2 := cids.ScalarMult(u2, keyPair.PublicKey)

	x := cids.AddECPoints(x1, x2)

	v := int(x.X) % p
	fmt.Println(v)
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
