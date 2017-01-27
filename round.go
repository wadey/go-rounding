package rounding

import (
	"math/big"
)

// RoundingMode describes how to round a given number.
// n is the integer that should be rounded such that the last digit is zero.
// l is the current last digit of n. For most cases you shouldn't need to
// implement this yourself, use one of the provided implementations in the
// rounding package.
type RoundingMode func(n, l *big.Int)

func roundUp(n, l *big.Int) {
	li := l.Int64()
	if li != 0 {
		if n.Sign() > 0 {
			n.Add(n, l.SetInt64(10-li))
		} else {
			n.Sub(n, l.SetInt64(10-li))
		}
	}
}

func roundDown(n, l *big.Int) {
	if n.Sign() > 0 {
		n.Sub(n, l)
	} else {
		n.Add(n, l)
	}
}

func roundFloor(n, l *big.Int) {
	if n.Sign() > 0 {
		n.Sub(n, l)
	} else {
		li := l.Int64()
		if li != 0 {
			n.Sub(n, l.SetInt64(10-li))
		}
	}
}

func roundHalfUp(n, l *big.Int) {
	if l.Int64() >= 5 {
		roundUp(n, l)
	} else {
		roundDown(n, l)
	}
}

func roundHalfDown(n, l *big.Int) {
	if l.Int64() > 5 {
		roundUp(n, l)
	} else {
		roundDown(n, l)
	}
}

func roundHalfEven(n, l *big.Int) {
	li := l.Int64()
	if li == 5 {
		k := new(big.Int).Rem(n, big100)
		ki := k.Int64() / 10
		if ki%2 == 0 {
			roundDown(n, l)
		} else {
			roundUp(n, l)
		}
	} else if li > 5 {
		roundUp(n, l)
	} else {
		roundDown(n, l)
	}
}

var (
	// Down rounds towards zero.
	Down RoundingMode = roundDown

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
	trunc(x, prec+1)
	n, d := x.Num(), x.Denom()
	l := new(big.Int).Rem(n, big10)
	l.Abs(l)

	method(n, l)

	// To force renormalization
	return x.SetFrac(n, d)
}
