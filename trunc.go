package rounding

import (
	"math/big"
)

// Sets x to the decimal value truncated at the given precision
// Returns the argument that was modified in place
// Example:
//   for DecTrunc(NewRat(2, 3), 2), x would be set to 66/100 (0.66)
func Trunc(x *big.Rat, prec int) *big.Rat {
	if x.Sign() == -1 {
		return Trunc(x.Abs(x), prec).Neg(x)
	}
	m := exp10(prec)
	n, d := x.Num(), x.Denom()
	n.Div(n.Mul(n, m), d)
	d.Set(m)
	return x
}
