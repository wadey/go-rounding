package rounding

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestFinite(t *testing.T) {
	test := func(x, y int64, finite bool) {
		if Finite(big.NewRat(x, y)) != finite {
			t.Errorf("Expected Finite(%v, %v) == %v", x, y, finite)
		}
	}
	test(1318, 185, false)
	test(66, 100, true)
	test(1, 2, true)
	test(-1, 2, true)
	test(1, 7, false)
	test(-1, 7, false)
}

func TestFinitePrec(t *testing.T) {
	test := func(x, y int64, prec int) {
		if actual := FinitePrec(big.NewRat(x, y)); actual != prec {
			t.Errorf("Expected FinitePrec(%v, %v) == %v, not %v", x, y, prec, actual)
		}
	}
	test(1, 16, 4)
	test(-1, 16, 4)
	test(1, 2, 1)
	test(5, 1, 0)
}

func TestFiniteString(t *testing.T) {
	test := func(x, y int64, s string) {
		if actual := FiniteString(big.NewRat(x, y)); actual != s {
			t.Errorf("Expected FiniteString(%v, %v) == %#v, not %#v", x, y, s, actual)
		}
	}
	test(1, 16, "0.0625")
	test(66, 100, "0.66")
	test(5, 1, "5")
}

func TestFiniteStringMin(t *testing.T) {
	test := func(x, y int64, prec int, s string) {
		if actual := FiniteStringMin(big.NewRat(x, y), prec); actual != s {
			t.Errorf("Expected FiniteStringMin(%v, %v, %v) == %#v, not %#v", x, y, prec, s, actual)
		}
	}
	test(60, 100, 2, "0.60")
	test(5, 1, 2, "5.00")
	test(-5, 1, 2, "-5.00")
}

func BenchmarkFinite(b *testing.B) {
	x := new(big.Rat)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x.SetFrac64(rand.Int63n(200)-100, rand.Int63n(100)+1)
		b.StartTimer()

		Finite(x)
	}
}

func BenchmarkFinitePrec(b *testing.B) {
	x := new(big.Rat)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x.SetFrac64(rand.Int63n(200)-100, rand.Int63n(100)+1)
		for !Finite(x) {
			x.SetFrac64(rand.Int63n(200)-100, rand.Int63n(100)+1)
		}
		b.StartTimer()

		FinitePrec(x)
	}
}
