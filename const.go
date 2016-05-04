// Package rounding provides rounding and truncation methods for big.Rat.
// It uses the same rounding terminology as Java's BigDecimal.
//
// For more information see the README at:
//   https://github.com/wadey/go-rounding
package rounding

import (
	"math/big"
)

const (
	expCacheLen = 20
)

var (
	big1   = big.NewInt(1)
	big2   = big.NewInt(2)
	big5   = big.NewInt(5)
	big10  = big.NewInt(10)
	big100 = big.NewInt(100)

	exp10cache = func() (cache [expCacheLen]*big.Int) {
		for i := 0; i < expCacheLen; i++ {
			m := big.NewInt(int64(i))
			cache[i] = m.Exp(big10, m, nil)
		}
		return
	}()
)

// Return 10 ** exp
func exp10(exp int) *big.Int {
	if exp < expCacheLen {
		return exp10cache[exp]
	}
	m := big.NewInt(int64(exp))
	return m.Exp(big10, m, nil)
}
