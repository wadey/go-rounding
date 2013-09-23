package rounding

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestTrunc(t *testing.T) {
	b, _ := new(big.Rat).SetString("0.1999")

	Trunc(b, 3)
	if b.FloatString(4) != "0.1990" {
		t.Fatal(b)
	}
	Trunc(b, 2)
	if b.FloatString(2) != "0.19" {
		t.Fatal(b)
	}
	Trunc(b, 3)
	if b.FloatString(2) != "0.19" {
		t.Fatal(b)
	}

	b = big.NewRat(2, 3)
	Trunc(b, 2)
	if b.FloatString(3) != "0.660" {
		t.Fatal(b)
	}
}

func TestNegativeTrunc(t *testing.T) {
	b, _ := new(big.Rat).SetString("-0.1999")

	Trunc(b, 3)
	if b.FloatString(4) != "-0.1990" {
		t.Fatal(b)
	}
	Trunc(b, 2)
	if b.FloatString(2) != "-0.19" {
		t.Fatal(b)
	}
	Trunc(b, 3)
	if b.FloatString(2) != "-0.19" {
		t.Fatal(b)
	}

	b = big.NewRat(-2, 3)
	Trunc(b, 2)
	if b.FloatString(3) != "-0.660" {
		t.Fatal(b)
	}
}

func TestSignAfterTrunc(t *testing.T) {
	b, _ := new(big.Rat).SetString("0.001")

	if b.Sign() != 1 {
		t.Fatal(b)
	}
	Trunc(b, 2)
	if b.Sign() != 0 {
		t.Fatal(b)
	}

	b.SetString("-0.001")

	if b.Sign() != -1 {
		t.Fatal(b)
	}
	Trunc(b, 2)
	if b.Sign() != 0 {
		t.Fatal(b)
	}
}

func BenchmarkTrunc(b *testing.B) {
	x := new(big.Rat)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x.SetFrac64(rand.Int63n(200)-100, rand.Int63n(100)+1)
		prec := rand.Intn(20)
		b.StartTimer()

		Trunc(x, prec)
	}
}
