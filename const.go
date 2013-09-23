// Rounding and truncation methods for big.Rat to close some of the gap in
// functionality between Rat and Decimal (such as Java's BigDecimal).
package rounding

import (
	"math/big"
)

const (
	expCacheLen = 20
)

var (
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
