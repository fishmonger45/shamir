package main

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"os"
)

var (
	ErrInvalidShareCount = errors.New("invalid share count")
	ErrRequiredTotal     = errors.New("total shares must be greater than the required")
)

type Share struct {
	Y *big.Int
	X int
}

type Polynomial struct {
	Coefficients []*big.Int
}

// Create a new polynomial from a random seed.
func NewPolynomial(required int) (*Polynomial, error) {
	if required < 0 {
		return nil, ErrInvalidShareCount
	}

	cs := make([]*big.Int, required)
	for i := 0; i < required; i++ {
		cs[i] = big.NewInt(rand.Int63())
	}

	return &Polynomial{
		Coefficients: cs,
	}, nil
}

func (polynomial *Polynomial) Evaluate(x int64) *big.Int {
	acc := new(big.Int)
	acc.Set(polynomial.Coefficients[0])
	fmt.Printf("acc: %v\n", acc)
	for exp := 1; exp < len(polynomial.Coefficients); exp++ {
		xbig := big.NewInt(x)
		fmt.Printf("xbig: %v\n", xbig)
		p := xbig.Exp(xbig, big.NewInt(int64(exp)), nil)
		fmt.Printf("p: %v\n", p)
		//w := big.NewInt(int64(p))
		r := new(big.Int)
		r.Mul(polynomial.Coefficients[exp], p)
		if r.Int64() != polynomial.Coefficients[exp].Int64() {
			fmt.Printf("r: %v\n", r)
			fmt.Printf("polynomial.Coefficients[exp]: %v\n", polynomial.Coefficients[exp])
			os.Exit(1)
		}
		acc.Add(acc, r)
	}

	return acc
}

// Split a secret into a total number of shares with a minimum number of shares required for secret reconstruction
func Split(secret *big.Int, polynomial *Polynomial, required int) ([]Share, error) {
	if required < 0 {
		return nil, ErrInvalidShareCount
	}

	// Total shares cannot be less than required shares
	if len(polynomial.Coefficients) < required {
		return nil, ErrRequiredTotal
	}
	shares := make([]Share, 0)

	for x := 1; x <= len(polynomial.Coefficients); x++ {
		value := polynomial.Evaluate(int64(x))
		shares = append(shares,
			Share{
				Y: value,
				X: x,
			})
	}

	return shares, nil
}

func Join(shares []Share) *big.Int {
	acc := big.NewInt(0)
	for j, _ := range shares {
		val := new(big.Int)
		val.Set(shares[j].Y)

		acc2 := big.NewInt(1)
		for m, _ := range shares {
			if m == j {
				continue
			}

			numerator := new(big.Int)
			numerator.Set(big.NewInt(int64(shares[m].X)))

			denominator := new(big.Int)
			denominator.Set(numerator)
			denominator.Sub(denominator, big.NewInt(int64(shares[j].X)))

			final := new(big.Int)
			final.Set(numerator)
			final.Div(final, denominator)

			acc2.Mul(acc2, final)
		}

		val.Mul(val, acc2)
		acc.Add(acc, val)
	}

	return acc
}
