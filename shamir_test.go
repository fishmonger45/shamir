package main

import (
	"fmt"
	"math/big"
	"testing"
	"testing/quick"
)

//func Test_Split(t *testing.T) {
//	payload := []byte("secret")
//	secret := new(big.Int)
//	secret.SetBytes(payload)

//	poly, err := NewPolynomial(5)
//	if err != nil {
//		t.Error(err)
//	}

//	// 5 total shares with 2 required for secret reconstruction
//	shares, err := Split(secret, poly, 2)
//	if err != nil {
//		t.Error(err)
//	}

//	fmt.Printf("shares: %v\n", shares)
//}

//func Test_Evaluate(t *testing.T) {
//	polynomial := Polynomial{

//		Coefficients: []*big.Int{
//			big.NewInt(1234),
//			big.NewInt(166),
//			big.NewInt(94),
//		},
//	}

//	result := polynomial.Evaluate(1)
//	fmt.Printf("result: %v\n", result)
//	fmt.Printf("(1234 + 166 + 94): %v\n", (1234 + 166 + 94))
//}

func Test_QuickCheck(t *testing.T) {
	counter := 0
	f := func(x, y, z int64) bool {
		a := big.NewInt(x)
		b := big.NewInt(y)
		c := big.NewInt(z)
		polynomial := Polynomial{
			Coefficients: []*big.Int{
				a, b, c,
			},
		}

		result := polynomial.Evaluate(1)

		counter += 1
		fmt.Printf("counter: %v\n", counter)

		return result == a.Add(a, b).Add(a, c)
	}

	if err := quick.Check(f, &quick.Config{
		MaxCount: 1,
	}); err != nil {
		t.Error(err)
	}
}

//func Test_Join(t *testing.T) {
//	payload := []byte("secret")
//	secret := new(big.Int)
//	secret.SetBytes(payload)

//	poly, err := NewPolynomial(5)
//	if err != nil {
//		t.Error(err)
//	}

//	// 5 total shares with 2 required for secret reconstruction
//	shares, err := Split(secret, poly, 2)
//	if err != nil {
//		t.Error(err)
//	}

//	reconstructed := Join(shares)

//	if reconstructed != secret {
//		fmt.Printf("reconstructed: %v\n", reconstructed)
//		fmt.Printf("secret: %v\n", secret)
//		t.FailNow()
//	}

//	fmt.Printf("shares: %v\n", shares)
//}
