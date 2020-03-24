package rounding

import (
	"math/big"
)

// RoundingMode describes how to round a given number.
// sign is the sign of the number (needed when n is zero)
// n is the integer that should be rounded such that the last digit is zero.
// l is the current last digit of n.
// r indicates whether the rest of the original number, after last digit, is non-zero
// NOTE: This method is only called when rounding is necessary, so if l == 0 it
// means there was more precision that was already truncated (so r is true).
// For most cases you shouldn't need to implement this yourself, use one of
// the provided implementations in the rounding package.
type RoundingMode func(sign int, n, l *big.Int, r bool)

func roundUp(sign int, n, l *big.Int, _ bool) {
	li := l.Int64()
	if sign >= 0 {
		n.Add(n, l.SetInt64(10-li))
	} else {
		n.Sub(n, l.SetInt64(10-li))
	}
}

func roundDown(sign int, n, l *big.Int, _ bool) {
	if sign > 0 {
		n.Sub(n, l)
	} else {
		n.Add(n, l)
	}
}

func roundCeil(sign int, n, l *big.Int, _ bool) {
	if sign > 0 {
		li := l.Int64()
		n.Add(n, l.SetInt64(10-li))
	} else {
		n.Add(n, l)
	}
}

func roundFloor(sign int, n, l *big.Int, _ bool) {
	if sign > 0 {
		n.Sub(n, l)
	} else {
		li := l.Int64()
		n.Sub(n, l.SetInt64(10-li))
	}
}

func roundHalfUp(sign int, n, l *big.Int, r bool) {
	if l.Int64() >= 5 {
		roundUp(sign, n, l, r)
	} else {
		roundDown(sign, n, l, r)
	}
}

func roundHalfDown(sign int, n, l *big.Int, r bool) {
	if l.Int64() > 5 {
		roundUp(sign, n, l, r)
	} else if l.Int64() <= 4 {
		roundDown(sign, n, l, r)
	} else {
		// l.Int64() == 5
		if r {
			roundUp(sign, n, l, r)
		} else {
			roundDown(sign, n, l, r)
		}
	}
}

func roundHalfEven(sign int, n, l *big.Int, r bool) {
	li := l.Int64()
	if li == 5 {
		if !r {
			k := new(big.Int).Rem(n, big100)
			ki := k.Int64() / 10
			if ki%2 == 0 {
				roundDown(sign, n, l, r)
			} else {
				roundUp(sign, n, l, r)
			}
		} else {
			roundUp(sign, n, l, r)
		}
	} else if li > 5 {
		roundUp(sign, n, l, r)
	} else {
		roundDown(sign, n, l, r)
	}
}

var (
	// Up rounds away from zero.
	Up RoundingMode = roundUp

	// Down rounds towards zero.
	Down RoundingMode = roundDown

	// Ceil rounds towards positive infinity.
	Ceil RoundingMode = roundCeil

	// Floor rounds towards negative infinity.
	Floor RoundingMode = roundFloor

	// HalfUp rounds towards "nearest neighbor" unless both neighbors are equidistant, in which case round up.
	HalfUp RoundingMode = roundHalfUp

	// HalfDown rounds towards "nearest neighbor" unless both neighbors are equidistant, in which case round down.
	HalfDown RoundingMode = roundHalfDown

	// HalfEven rounds towards the "nearest neighbor" unless both neighbors are equidistant, in which case, round towards the even neighbor.
	HalfEven RoundingMode = roundHalfEven
)

// Round sets x to its value rounded to the given precision using the given rounding mode.
// Returns x, which was modified in place.
func Round(x *big.Rat, prec int, method RoundingMode) *big.Rat {
	sign := x.Sign()
	orig := new(big.Rat).Set(x)
	trunc(x, prec+1)
	n, d := x.Num(), x.Denom()
	l := new(big.Int).Rem(n, big10)
	l.Abs(l)

	r := x.Cmp(orig) != 0
	// Only run the rounding method if just truncating won't suffice
	if l.Sign() != 0 || r {
		method(sign, n, l, r)
	}

	// To force renormalization
	return x.SetFrac(n, d)
}
