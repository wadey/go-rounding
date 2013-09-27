package rounding

import (
	"fmt"
	"math/big"
)

func ExampleTrunc() {
	x := big.NewRat(2, 3)
	Trunc(x, 2)
	fmt.Println(x.FloatString(2))
	// Output: 0.66
}

func ExampleRound() {
	x, _ := new(big.Rat).SetString("0.125")
	Round(x, 2, HalfEven)
	fmt.Println(x.FloatString(2))
	// Output: 0.12
}

func ExampleFinite() {
	x, _ := new(big.Rat).SetString("0.125")
	fmt.Println(Finite(x))
	fmt.Println(Finite(big.NewRat(2, 3)))
	// Output:
	// true
	// false
}

func ExampleFinitePrec() {
	x, _ := new(big.Rat).SetString("0.125")
	fmt.Println(FinitePrec(x))
	// Output: 3
}

func ExampleFiniteString() {
	x, _ := new(big.Rat).SetString("0.125")
	fmt.Println(FiniteString(x))
	// Output: 0.125
}

func ExampleFiniteStringMin() {
	x, _ := new(big.Rat).SetString("5")
	fmt.Println(FiniteStringMin(x, 2))
	// Output: 5.00
}
