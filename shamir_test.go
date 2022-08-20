package main

import (
	"fmt"
	"math/big"
	"testing"
)

func Test_Split(t *testing.T) {
	payload := []byte("secret")
	secret := new(big.Int)
	secret.SetBytes(payload)

	poly, err := NewPolynomial(5)
	if err != nil {
		t.Error(err)
	}

	// 5 total shares with 2 required for secret reconstruction
	shares, err := Split(secret, poly, 2)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("shares: %v\n", shares)
}
