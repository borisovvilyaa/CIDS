// package main

// import (
// 	cids "cids/task7"
// 	"crypto/sha256"
// 	"encoding/binary"
// 	"encoding/hex"
// 	"fmt"
// 	"math"
// 	"math/rand"
// )

// type KeyPair struct {
// 	PrivateKey int
// 	PublicKey  cids.ECPoint
// }

// type signature struct {
// 	firstElement  int
// 	secondElement cids.ECPoint
// }

// func generatePrivateKey(curveOrder int) int {
// 	for {
// 		n := rand.Intn(curveOrder)
// 		if n != 0 {
// 			return rand.Intn(curveOrder)
// 		}
// 	}
// }

// func computePublicKey(privateKey int, basePoint cids.ECPoint) cids.ECPoint {
// 	return cids.ScalarMult(privateKey, basePoint)
// }

// func printUserDetails(user int, keyPair KeyPair) {
// 	fmt.Printf("User %d:\n", user)
// 	fmt.Printf("Private Key: %d\n", keyPair.PrivateKey)
// 	fmt.Printf("Public Key: (%f, %f)\n\n", keyPair.PublicKey.X, keyPair.PublicKey.Y)
// }

// func hashData(input string) [32]byte {
// 	bytes := []byte(input)
// 	hash := sha256.Sum256(bytes)
// 	return hash
// }
// func modInverse(a, m int) int {
// 	if m == 0 {
// 		return 0
// 	}

// 	g := gcdExtended(a, m)
// 	if g < 0 {
// 		g += m
// 	}

// 	return g
// }

// func gcdExtended(a, b int) int {
// 	if a == 0 {
// 		return b
// 	}

// 	gcd := gcdExtended(b%a, a)
// 	x := gcdExtended(b%a, a)

// 	return x - (b/a)*gcd
// }

// func HashToNumber(hash []byte) uint64 {
// 	// Хеш SHA256 имеет размер 32 байта (256 бит)
// 	// Мы будем преобразовывать первые 8 байт хеша в число типа uint64
// 	number := binary.BigEndian.Uint64(hash[:8])
// 	return number
// }

// func main() {
// 	basePoint := cids.BasePointGGet()
// 	curveOrder := cids.CurveOrderNGet()

// 	p := 114973

// 	//1 step. Create key pair
// 	privateKey := generatePrivateKey(curveOrder - 1)
// 	keyPair := KeyPair{
// 		PrivateKey: privateKey,
// 		PublicKey:  computePublicKey(privateKey, basePoint),
// 	}
// 	printUserDetails(1, keyPair)

// 	// Signature step.

// 	input := "Hello, world!"
// 	hash := hashData(input)
// 	number := HashToNumber(hash[:])
// 	fmt.Println("Input string:", input)
// 	fmt.Printf("SHA256 hash number:%d\n", number)
// 	fmt.Printf("SHA256 hash:%s\n\n", hex.EncodeToString(hash[:]))

// 	nonce := rand.Intn(cids.CurveOrderNGet())
// 	fmt.Printf("Nonce: %d\n\n", nonce)

// 	point_multiplication := cids.ScalarMult(nonce, cids.BasePointGGet())
// 	fmt.Printf("Point multiplication: %f\n\n", point_multiplication)

// 	r := int(point_multiplication.X) % p
// 	s := int(math.Pow(float64(nonce), -1)*float64(int(number)+int(keyPair.PrivateKey)*r)) % p
// 	fmt.Println(r, s)

// 	//verify signature

// 	w := int(math.Pow(float64(s), -1)) % p
// 	u1 := (int(number) * w) % p
// 	u2 := r * w % p
// 	x1 := cids.ScalarMult(u1, cids.BasePointGGet())
// 	x2 := cids.ScalarMult(u2, keyPair.PublicKey)

// 	x := cids.AddECPoints(x1, x2)

// 	v := int(x.X) % p
// 	fmt.Println(v)
// }

// // func hashData(input string) [32]byte {
// // 	bytes := []byte(input)
// // 	hash := sha256.Sum256(bytes)
// // 	return hash
// // }

// // func generateBlindingPoint(nonce int, basePoint cids.ECPoint) cids.ECPoint {
// // 	return cids.ScalarMult(nonce, basePoint)
// // }

// // func generateSignature(privateKey, nonce int, hash [32]byte, blindingPoint cids.ECPoint, countPoint int) signature {
// // 	firstComp := (int(blindingPoint.X) + privateKey) % countPoint
// // 	secondComp := cids.ScalarMult(int(binary.BigEndian.Uint32(hash[:4]))/(privateKey)/(nonce), blindingPoint)
// // 	return signature{
// // 		firstElement:  firstComp,
// // 		secondElement: secondComp,
// // 	}
// // }

// // func 	verifySignature(input string, hash [32]byte, signature signature, keyPair KeyPair, countPoint int) bool {
// // 	Q := cids.AddECPoints(cids.ScalarMult(int(signature.secondElement.Y), cids.BasePointGGet()), cids.ScalarMult(signature.firstElement, keyPair.PublicKey))
// // 	fmt.Println(signature.firstElement, int(Q.X))
// // 	return signature.firstElement == int(Q.X)
// // }

// // func main() {
// // 	countPoint := 0
// // 	basePoint := cids.BasePointGGet()
// // 	countPoint++
// // 	curveOrder := 256

// // 	privateKey := generatePrivateKey(curveOrder)
// // 	keyPair := KeyPair{
// // 		PrivateKey: privateKey,
// // 		PublicKey:  computePublicKey(privateKey, basePoint),
// // 	}
// // 	countPoint++

// // 	printUserDetails(1, keyPair)

// // 	input := "Hello, world!"
// // 	hash := hashData(input)

// // 	fmt.Println("Input string:", input)
// // 	fmt.Printf("SHA256 hash:%s\n\n", hex.EncodeToString(hash[:]))

// // 	nonce := rand.Intn(100)
// // 	fmt.Printf("Nonce: %d\n\n", nonce)

// // 	blindingPoint := generateBlindingPoint(nonce, basePoint)
// // 	countPoint++
// // 	fmt.Printf("Blinding point: (%f, %f)\n\n", blindingPoint.X, blindingPoint.Y)

// // 	fmt.Printf("Count Point: %d\n\n", countPoint)

// // 	signature := generateSignature(privateKey, nonce, hash, blindingPoint, countPoint)
// // 	fmt.Printf("First Component: %d\n\n", signature.firstElement)
// // 	fmt.Println("Second Component:", signature.secondElement)

// // 	fmt.Printf("\n\n-------Verify-------\n\n")

// // 	inputverify := "Hello, world!"
// // 	hashverify := hashData(input)

// // 	fmt.Println("Input string:", inputverify)
// // 	fmt.Printf("SHA256 hash:%s\n\n", hex.EncodeToString(hashverify[:]))

// // 	fmt.Println(verifySignature(input, hashverify, signature, keyPair, countPoint))
// // }

// package main

// import (
// 	elliptic "cids/task7"
// 	"crypto/sha256"
// 	"encoding/binary"
// 	"encoding/hex"
// 	"fmt"
// 	"math"
// 	"math/big"
// 	"math/rand"
// )

// type KeyPair struct {
// 	PrivateKey int
// 	PublicKey  elliptic.ECPoint
// }

// func isPrime(number int) bool {
// 	if number <= 1 {
// 		return false
// 	}

// 	// Проверяем делители до квадратного корня числа
// 	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
// 		if number%i == 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
// func generate_p(n int) int {
// 	for {
// 		number := rand.Intn(n - 1)
// 		if isPrime(number) {
// 			return number
// 		}
// 	}
// }

// func modInverse(k, n int) (int, error) {
// 	g := big.NewInt(int64(k))
// 	mod := big.NewInt(int64(n))

// 	g.Mod(g, mod)

// 	if g.Cmp(big.NewInt(0)) == 0 {
// 		return 0, fmt.Errorf("No modular inverse exists")
// 	}

// 	x := new(big.Int)
// 	y := new(big.Int)
// 	d := new(big.Int)

// 	d.GCD(x, y, g, mod)

// 	if d.Cmp(big.NewInt(1)) != 0 {
// 		return 0, fmt.Errorf("No modular inverse exists")
// 	}

// 	x.Mod(x, mod)

// 	if x.Cmp(big.NewInt(0)) < 0 {
// 		x.Add(x, mod)
// 	}

// 	return int(x.Int64()), nil
// }
// func HashToNumber(hash []byte) int {
// 	number := binary.BigEndian.Uint64(hash[:8])
// 	return int(number)
// }

// func hashData(input string) [32]byte {

// 	bytes := []byte(input)
// 	hash := sha256.Sum256(bytes)
// 	return hash
// }
// func main() {
// 	//Settings

// 	//it's number from wiki. I don't understud, it's number i need to calculating, or it's number constant
// 	//link: https://ru.wikipedia.org/wiki/ECDSA

// 	p := 114973
// 	n := 114467

// 	// Generate key pair
// 	privateKey := generate_p(n)
// 	KeyPair := KeyPair{
// 		PrivateKey: privateKey,
// 		PublicKey:  elliptic.PointMultiplication(elliptic.BasePointGGet(), privateKey),
// 	}
// 	fmt.Printf("\n-------------KeyPair-------------\nPrivate Key %d\nPublic Key: %v\n---------------------------------\n\n", KeyPair.PrivateKey, KeyPair.PublicKey)

// 	//Signature
// 	//Create nonce number
// 	nonce := rand.Intn(n - 1)
// 	for {
// 		if nonce%2 == 0 {
// 			break
// 		} else {
// 			nonce = rand.Intn(n - 1)
// 		}
// 	}
// 	fmt.Printf("Nonce: %d\n\n", nonce)

// 	nonceBasePoint := elliptic.PointMultiplication(elliptic.BasePointGGet(), nonce)
// 	fmt.Printf("Nonce base point: %f , %f\n\n", nonceBasePoint.X, nonceBasePoint.Y)

// 	//r
// 	r := int(nonceBasePoint.X) % p

// 	inv, err := modInverse(nonce, 2)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	input := "Hello, world!"
// 	hash := hashData(input)
// 	number := HashToNumber(hash[:])
// 	fmt.Println("Input string:", input)
// 	fmt.Printf("SHA256 hash number:%d\n", number)
// 	fmt.Printf("SHA256 hash:%s\n\n", hex.EncodeToString(hash[:]))

// 	s := (inv*number + KeyPair.PrivateKey*r) % p

// 	fmt.Printf("Signature: %d, %d", r, s)

// }

// v3

package main

// import (
// 	elliptic "cids/task7"
// 	"crypto/sha256"
// 	"fmt"
// 	"math/big"
// 	"math/rand"
// )

// type KeyPair struct {
// 	PrivateKey *big.Int
// 	PublciKey  elliptic.ECPoint
// }

// func main() {
// 	// Settings

// 	p := big.NewInt(114973)
// 	n := 114467

// 	sk := big.NewInt(int64(rand.Intn(n - 1)))
// 	keys := KeyPair{
// 		PrivateKey: sk,
// 		PublciKey:  elliptic.MultiplyECPoint(*sk, elliptic.BasePointGGet()),
// 	}
// 	fmt.Printf("\n-------------KeyPair-------------\n")
// 	fmt.Printf("Private Key: %s\n", keys.PrivateKey)
// 	fmt.Printf("Public Key: X: %s, Y: %s\n", keys.PublciKey.X.String(), keys.PublciKey.Y.String())
// 	fmt.Printf("---------------------------------\n\n")

// 	//Signature
// 	k := big.NewInt(int64(rand.Intn(n - 1)))
// 	kG := elliptic.ScalarMult(*k, elliptic.BasePointGGet())
// 	r := new(big.Int)
// 	r.Mod(kG.X, p)

// 	// Получаем хеш-значение
// 	hash := sha256.Sum256([]byte("Hello, world!"))

// 	// Создаем новый big.Int
// 	hashInt := new(big.Int)

// 	// Устанавливаем значение хеш-значения в big.Int
// 	hashInt.SetBytes(hash[:])
// 	fmt.Println("---------------------------------HASH----------------------------------------")
// 	fmt.Println(hashInt.String())
// 	fmt.Println("-----------------------------------------------------------------------------")

// 	// Вычисление обратного элемента k^-1
// 	kInv := new(big.Int).ModInverse(k, p)

// 	// Вычисление s
// 	s := new(big.Int).Mul(hashInt, keys.PrivateKey)
// 	s.Add(s, kInv)
// 	s.Mod(s, p)

// 	fmt.Println("---------Signature--------")
// 	fmt.Printf("| %-10s | %-10s |\n", "r", "s")
// 	fmt.Println("|------------|------------|")
// 	fmt.Printf("| %-10s | %-10s |\n", r.String(), s.String())
// 	fmt.Println("---------------------------")

// 	// Verify signature
// 	// Значения

// 	// Вычисление обратного элемента s^-1
// 	w := new(big.Int).ModInverse(s, p)

// 	// Вычисление u1
// 	u1 := new(big.Int).Mul(hashInt, w)
// 	u1.Mod(u1, p)

// 	u2 := new(big.Int).Mul(r, w)
// 	u2.Mod(u2, p)

// 	X := elliptic.AddECPoints(elliptic.ScalarMult(*u1, elliptic.BasePointGGet()), elliptic.ScalarMult(*u2, keys.PublciKey))
// 	v := new(big.Int).Mod(X.X, p)
// 	fmt.Println("v =", v, "r =", r)
// 	fmt.Println(v == r)
// }


// package main

// import (
// 	elliptic "cids/task7"
// 	"crypto/sha256"
// 	"fmt"
// 	"math/big"
// 	"math/rand"
// )

// type KeyPair struct {
// 	PrivateKey *big.Int
// 	PublciKey  elliptic.ECPoint
// }

// func main() {
// 	// Settings

// 	p := big.NewInt(114973)
// 	n := 114973

// 	sk := big.NewInt(int64(rand.Intn(n)))

// 	keys := KeyPair{
// 		PrivateKey: sk,
// 		PublciKey:  elliptic.MultiplyECPoint(*sk, elliptic.BasePointGGet()),
// 	}

// 	fmt.Printf("\n-------------KeyPair-------------\n")
// 	fmt.Printf("Private Key: %s\n", keys.PrivateKey)
// 	fmt.Printf("Public Key: X: %s, Y: %s\n", keys.PublciKey.X.String(), keys.PublciKey.Y.String())
// 	fmt.Printf("---------------------------------\n\n")

// 	//Signature
// 	k := big.NewInt(int64(rand.Intn(n)))
// 	kG := elliptic.ScalarMult(*k, elliptic.BasePointGGet())
// 	r := new(big.Int).Mod(kG.X, p)

// 	// Получаем хеш-значение
// 	hash := sha256.Sum256([]byte("Hello, world!"))

// 	// Создаем новый big.Int
// 	hashInt := new(big.Int)

// 	// Устанавливаем значение хеш-значения в big.Int
// 	hashInt.SetBytes(hash[:])
// 	fmt.Println("---------------------------------HASH----------------------------------------")
// 	fmt.Println(hashInt.String())
// 	fmt.Println("-----------------------------------------------------------------------------")

// 	// Вычисление обратного элемента k^-1
// 	kInv := new(big.Int).ModInverse(k, p)

// 	// Вычисление s
// 	s := new(big.Int).Mul(hashInt, keys.PrivateKey)
// 	s.Add(s, kInv)
// 	s.Mod(s, p)

// 	fmt.Println("---------Signature--------")
// 	fmt.Printf("| %-10s | %-10s |\n", "r", "s")
// 	fmt.Println("|------------|------------|")
// 	fmt.Printf("| %-10s | %-10s |\n", r.String(), s.String())
// 	fmt.Println("---------------------------")

// 	// Verify signature
// 	// Значения

// 	// Вычисление обратного элемента s^-1
// 	w := new(big.Int).ModInverse(s, p)

// 	// Вычисление u1
// 	u1 := new(big.Int).Mul(hashInt, w)
// 	u1.Mod(u1, p)

// 	u2 := new(big.Int).Mul(r, w)
// 	u2.Mod(u2, p)

// 	X := elliptic.AddECPoints(elliptic.ScalarMult(*u1, elliptic.BasePointGGet()), elliptic.ScalarMult(*u2, keys.PublciKey))
// 	v := new(big.Int).Mod(X.X, p)
// 	fmt.Println("v =", v, "r =", r)
// 	fmt.Println(v == r)
// }

