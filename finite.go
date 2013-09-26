package rounding

import (
	"math/big"
)

// Wrapper around QuoRem that returns the remainder
func remQuo(x, y, q, r *big.Int) *big.Int {
	q.QuoRem(x, y, r)
	return r
}

// Finite returns true if x has a finite decimal representation.
// (x is finite if the Denom can be represented as (2**x) * (5**y))
func Finite(x *big.Rat) bool {
	d := new(big.Int).Set(x.Denom())
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

// FinitePrec returns the precision of the finite deciml representation of x.
// WARNING: Running this on a value that does not have a finite decimal
// representation will result in an infinite loop. Always check with Finite()
// first if you are unsure.
func FinitePrec(x *big.Rat) int {
	d := x.Denom()
	n := new(big.Int).Set(x.Num())
	m := new(big.Int)

	var i int
	for m.Mod(n, d).Sign() != 0 {
		i++
		n.Mul(n, big10)
	}
	return i
}

// FiniteString returns the equivalent of x.FloatString(FinitPrec(x)).
// WARNING: Running this on a value that does not have a finite decimal
// representation will result in an infinite loop. Always check with Finite()
// first if you are unsure.
func FiniteString(x *big.Rat) string {
	return x.FloatString(FinitePrec(x))
}

// FiniteString returns the equivalent of x.FloatString(max(FinitPrec(x), prec)).
// WARNING: Running this on a value that does not have a finite decimal
// representation will result in an infinite loop. Always check with Finite()
// first if you are unsure.
func FiniteStringMin(x *big.Rat, prec int) string {
	p := FinitePrec(x)
	if p < prec {
		p = prec
	}
	return x.FloatString(p)
}
