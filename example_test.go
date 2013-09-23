package rounding

import (
	"fmt"
	"math/big"
)

func ExampleTrunc() {
	x := big.NewRat(2, 3)
	Trunc(x, 2)
	fmt.Println(x.FloatString(2))
	// Output:
	// 0.66
}

func ExampleRound() {
	x, _ := new(big.Rat).SetString("0.125")
	Round(x, 2, HalfEven)
	fmt.Println(x.FloatString(2))
	// Output:
	// 0.12
}
