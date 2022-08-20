package main

import (
	"errors"
	"math"
	"math/big"
	"math/rand"
)

var (
	ErrInvalidShareCount = errors.New("invalid share count")
	ErrRequiredTotal     = errors.New("total shares must be greater than the required")
)

type Share struct {
	Part  *big.Int
	Piece int
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

// Split a secret into a total number of shares with a minimum number of shares required for secret reconstruction
func Split(secret *big.Int, polynomial *Polynomial, required int) ([]Share, error) {
	if required < 0 {
		return nil, ErrInvalidShareCount
	}

	// Total shares cannot be less than required shares
	if len(polynomial.Coefficients) < required {
		return nil, ErrRequiredTotal
	}
	cs := make([]*big.Int, 0)
	ss := make([]Share, 0)

	cs = append(cs, secret)
	for i := 0; i < required; i++ {
		cs = append(cs, big.NewInt(rand.Int63()))
	}

	for x := 1; x <= len(polynomial.Coefficients); x++ {

		acc := new(big.Int)
		acc.Set(cs[0])
		for exp := 1; exp <= required; exp++ {
			p := math.Pow(float64(x), float64(exp))
			w := big.NewInt(int64(p))
			r := new(big.Int)
			r.Mul(cs[exp], w)
			acc.Add(acc, r)
		}

		ss = append(ss,
			Share{
				Part:  acc,
				Piece: x,
			})
	}

	return ss, nil
}
