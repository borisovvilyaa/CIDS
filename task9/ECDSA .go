package main

import (
	elliptic "cids/task7"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
)

type Signature struct {
	r *big.Int
	s *big.Int
}

type KeyPair struct {
	PrivateKey *big.Int
	PublicKey  elliptic.ECPoint
}

const n = 11400000

func GenerateKey() KeyPair {
	sk := big.NewInt(int64(rand.Intn(n - 1)))

	return KeyPair{
		PrivateKey: sk,
		PublicKey:  elliptic.ScalarMult(*sk, elliptic.BasePointGGet()),
	}
}

func PrintKeyPair(key KeyPair) {
	fmt.Printf("Public Key: (%v, %v) \n", key.PublicKey.X, key.PublicKey.Y)
	fmt.Println("Private Key: ", key.PrivateKey)
}

func CreateSignature(key KeyPair, message *big.Int) Signature {
	k := big.NewInt(int64(rand.Intn(n - 1)))
	kG := elliptic.ScalarMult(*k, elliptic.BasePointGGet())

	r := new(big.Int).Mod(kG.X, big.NewInt(n))
	s := new(big.Int).ModInverse(k, big.NewInt(n))
	s.Add(message, new(big.Int).Mul(key.PrivateKey, r))
	s.Mod(s, big.NewInt(n))

	return Signature{
		r: r,
		s: s,
	}
}

func PrintSignature(signature Signature) {
	fmt.Printf("r: %v\n", signature.r)
	fmt.Printf("s: %v\n", signature.s)
}

func VerifySignature(key KeyPair, message *big.Int, signature Signature) bool {
	w := new(big.Int).ModInverse(big.NewInt(n), signature.s)
	w.Mod(w, big.NewInt(n))

	u1 := new(big.Int).Mul(message, w)
	u1.Mod(u1, big.NewInt(n))

	u2 := new(big.Int).Mul(signature.r, w)
	u2.Mod(u1, big.NewInt(n))

	xVerify := elliptic.AddECPoints(elliptic.ScalarMult(*u1, elliptic.BasePointGGet()), elliptic.ScalarMult(*u2, key.PublicKey))

	v := new(big.Int).Mod(xVerify.X, big.NewInt(n))

	return v.Cmp(signature.r) == 0
}

func SerializePrivateKey(key *big.Int) string {
	return key.String()
}

func DeserializePrivateKey(data string) (*big.Int, error) {
	key := new(big.Int)
	_, success := key.SetString(data, 10)
	if !success {
		return nil, fmt.Errorf("failed to deserialize private key")
	}
	return key, nil
}

func SerializePublicKey(key elliptic.ECPoint) string {
	x := key.X.String()
	y := key.Y.String()
	return x + "," + y
}

func DeserializePublicKey(data string) (elliptic.ECPoint, error) {
	var key elliptic.ECPoint
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return key, fmt.Errorf("invalid public key format")
	}
	x := new(big.Int)
	y := new(big.Int)
	_, success1 := x.SetString(parts[0], 10)
	_, success2 := y.SetString(parts[1], 10)
	if !success1 || !success2 {
		return key, fmt.Errorf("failed to deserialize public key")
	}
	key.X = x
	key.Y = y
	return key, nil
}

func SerializeSignature(signature Signature) string {
	r := signature.r.Bytes()
	s := signature.s.Bytes()
	data := append(r, s...)
	return hex.EncodeToString(data)
}

func DeserializeSignature(data string) (Signature, error) {
	signature := Signature{}
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return signature, fmt.Errorf("failed to deserialize signature")
	}
	rSize := (len(bytes) + 1) / 2
	rBytes := bytes[:rSize]
	sBytes := bytes[rSize:]
	signature.r = new(big.Int).SetBytes(rBytes)
	signature.s = new(big.Int).SetBytes(sBytes)
	return signature, nil
}

func main() {
	keys := GenerateKey()
	PrintKeyPair(keys)
	message := sha256.Sum256([]byte("Hello, world!"))
	messageInt := new(big.Int).SetBytes(message[:])

	signature := CreateSignature(keys, messageInt)
	PrintSignature(signature)

	serializedPrivateKey := SerializePrivateKey(keys.PrivateKey)
	serializedPublicKey := SerializePublicKey(keys.PublicKey)
	serializedSignature := SerializeSignature(signature)

	fmt.Println("Serialized Private Key:", serializedPrivateKey)
	fmt.Println("Serialized Public Key:", serializedPublicKey)
	fmt.Println("Serialized Signature:", serializedSignature)

	deserializedPrivateKey, err := DeserializePrivateKey(serializedPrivateKey)
	if err != nil {
		fmt.Println("Failed to deserialize private key:", err)
		return
	}

	deserializedPublicKey, err := DeserializePublicKey(serializedPublicKey)
	if err != nil {
		fmt.Println("Failed to deserialize public key:", err)
		return
	}

	deserializedSignature, err := DeserializeSignature(serializedSignature)
	if err != nil {
		fmt.Println("Failed to deserialize signature:", err)
		return
	}

	fmt.Println("Deserialized Private Key:", deserializedPrivateKey)
	fmt.Println("Deserialized Public Key:", deserializedPublicKey)
	fmt.Println("Deserialized Signature:", deserializedSignature)

	fmt.Println("Signature Verification:", VerifySignature(KeyPair{PrivateKey: deserializedPrivateKey, PublicKey: deserializedPublicKey}, messageInt, deserializedSignature))
}
