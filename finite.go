package rounding

import (
	"fmt"
	"math/big"
)

// Wrapper around QuoRem that returns the remainder
func remQuo(x, y, q, r *big.Int) *big.Int {
	q.QuoRem(x, y, r)
	return r
}

// Finite returns true if x has a finite decimal representation.
// x is finite if the Denom can be represented as (2**x) * (5**y).
func Finite(x *big.Rat) bool {
	// calling x.Denom() can modify x (populates b) so be extra careful
	xx := new(big.Rat).Set(x)

	d := xx.Denom()
	i := new(big.Int)
	m := new(big.Int)

	for {
		switch {
		case d.Cmp(big1) == 0:
			return true
		case remQuo(d, big2, i, m).Sign() == 0:
			d.Set(i)
		case remQuo(d, big5, i, m).Sign() == 0:
			d.Set(i)
		default:
			return false
		}
	}
}

// FinitePrec returns the precision of the finite decimal representation of x.
// WARNING: Running this on a value that does not have a finite decimal
// representation will panic.
func FinitePrec(x *big.Rat) int {
	if !Finite(x) {
		panic(fmt.Errorf("rounding.FinitePrec: called with non-finite value: %v", x))
	}
	// calling x.Denom() can modify x (populates b) so be extra careful
	xx := new(big.Rat).Set(x)

	d := xx.Denom()
	n := xx.Num()
	m := new(big.Int)

	var i int
	for m.Mod(n, d).Sign() != 0 {
		i++
		n.Mul(n, big10)
	}
	return i
}

// FiniteString returns the equivalent of x.FloatString(FinitePrec(x)).
// WARNING: Running this on a value that does not have a finite decimal
// representation will panic.
func FiniteString(x *big.Rat) string {
	return x.FloatString(FinitePrec(x))
}

// FiniteStringMin returns the equivalent of x.FloatString(max(FinitePrec(x), prec)).
// WARNING: Running this on a value that does not have a finite decimal
// representation will panic.
func FiniteStringMin(x *big.Rat, prec int) string {
	p := FinitePrec(x)
	if p < prec {
		p = prec
	}
	return x.FloatString(p)
}
