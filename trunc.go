package rounding

import (
	"math/big"
)

// Trunc sets x to the decimal value truncated at the given precision.
// Returns x, which was modified in place.
func Trunc(x *big.Rat, prec int) *big.Rat {
	trunc(x, prec)

	// To force renormalization
	return x.SetFrac(x.Num(), x.Denom())
}

// truncate without normalization (rounding.Round depends on this)
func trunc(x *big.Rat, prec int) *big.Rat {
	if x.Sign() == -1 {
		return trunc(x.Abs(x), prec).Neg(x)
	}
	m := exp10(prec)
	n, d := x.Num(), x.Denom()
	n.Div(n.Mul(n, m), d)
	d.Set(m)
	return x
}
