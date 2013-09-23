package rounding

import (
	"math/big"
)

type RoundingMode int

const (
	// Rounding mode to round away from zero.
	Up RoundingMode = iota
	// Rounding mode to round towards zero.
	Down
	// Rounding mode to round towards positive infinity.
	Ceil
	// Rounding mode to round towards negative infinity.
	Floor
	// Rounding mode to round towards "nearest neighbor" unless both neighbors are equidistant, in which case round up.
	HalfUp
	// Rounding mode to round towards "nearest neighbor" unless both neighbors are equidistant, in which case round down.
	HalfDown
	// Rounding mode to round towards the "nearest neighbor" unless both neighbors are equidistant, in which case, round towards the even neighbor.
	HalfEven
)

// Round sets x to its value rounded to the given precision using the given rounding mode.
// Returns x, which was modified in place.
func Round(x *big.Rat, prec int, method RoundingMode) *big.Rat {
	Trunc(x, prec+1)
	n, d := x.Num(), x.Denom()
	l := new(big.Int).Rem(n, big10)
	l.Abs(l)
	li := l.Int64()

	switch method {
	case HalfUp:
		if li >= 5 {
			method = Up
		} else {
			method = Down
		}
	case HalfDown:
		if li > 5 {
			method = Up
		} else {
			method = Down
		}
	case HalfEven:
		if li == 5 {
			k := new(big.Int).Rem(n, big100)
			ki := k.Int64() / 10
			if ki%2 == 0 {
				method = Down
			} else {
				method = Up
			}
		} else if li > 5 {
			method = Up
		} else {
			method = Down
		}
	}

	switch method {
	case Up:
		if li != 0 {
			if x.Sign() > 0 {
				n.Add(n, l.SetInt64(10-li))
			} else {
				n.Sub(n, l.SetInt64(10-li))
			}
		}
	case Down:
		if x.Sign() > 0 {
			n.Sub(n, l)
		} else {
			n.Add(n, l)
		}
	case Ceil:
		if x.Sign() > 0 {
			if li != 0 {
				n.Add(n, l.SetInt64(10-li))
			}
		} else {
			n.Add(n, l)
		}
	case Floor:
		if x.Sign() > 0 {
			n.Sub(n, l)
		} else {
			if li != 0 {
				n.Sub(n, l.SetInt64(10-li))
			}
		}
	}

	// To force renormalization
	return x.SetFrac(n, d)
}
