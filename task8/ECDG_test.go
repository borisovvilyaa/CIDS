package main

import (
	cids "cids/task7"
	"math/big"
	"testing"
)

// TestSharedSecretComparisonBig tests the comparison of shared secrets using very large numbers.
// It generates two private keys, computes the corresponding public keys, and then computes the shared secrets.
// The test checks that the shared secrets should be equal.
func TestSharedSecretComparisonBig(t *testing.T) {
	basePoint := cids.BasePointGGet()

	privateKey1 := big.NewInt(999999999999999999)
	privateKey2 := big.NewInt(999999999999999991)
	publicKey1 := computeSharedSecret(*privateKey1, basePoint)
	publicKey2 := computeSharedSecret(*privateKey2, basePoint)

	sharedSecret1 := computeSharedSecret(*privateKey1, publicKey2)
	sharedSecret2 := computeSharedSecret(*privateKey2, publicKey1)

	if !cids.IsEqual(sharedSecret1, sharedSecret2) {
		t.Error("Shared secrets should not be equal")
	}
}

// TestSharedSecretComparisonSmall tests the comparison of shared secrets using small numbers.

func TestSharedSecretComparisonSmall(t *testing.T) {
	basePoint := cids.BasePointGGet()

	privateKey1 := big.NewInt(1)
	privateKey2 := big.NewInt(2)
	publicKey1 := computeSharedSecret(*privateKey1, basePoint)
	publicKey2 := computeSharedSecret(*privateKey2, basePoint)

	sharedSecret1 := computeSharedSecret(*privateKey1, publicKey2)
	sharedSecret2 := computeSharedSecret(*privateKey2, publicKey1)

	if !cids.IsEqual(sharedSecret1, sharedSecret2) {
		t.Error("Shared secrets should not be equal")
	}
}

// TestAll runs all the test cases.

func TestAll(t *testing.T) {
	t.Run("TestSharedSecretComparison (very big nubmer)", TestSharedSecretComparisonBig)
	t.Run("TestSharedSecretComparison (small nubmer)", TestSharedSecretComparisonSmall)

}
