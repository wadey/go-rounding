// Rounding and truncation methods for big.Rat to close some of the gap in
// functionality between Rat and Decimal (such as Java's BigDecimal).
package rounding

import (
	"math/big"
)

var (
	big10  = big.NewInt(10)
	big100 = big.NewInt(100)
)

// Return 10 ** exp
func exp10(exp int) *big.Int {
	m := big.NewInt(int64(exp))
	return m.Exp(big10, m, nil)
}
