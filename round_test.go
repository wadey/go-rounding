package rounding

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestRoundUp(t *testing.T) {
	m := roundUp
	testRounding(t, "5.5", "6", 0, m)
	testRounding(t, "2.5", "3", 0, m)
	testRounding(t, "1.6", "2", 0, m)
	testRounding(t, "1.1", "2", 0, m)
	testRounding(t, "1.0", "1", 0, m)
	testRounding(t, "-1.0", "-1", 0, m)
	testRounding(t, "-1.1", "-2", 0, m)
	testRounding(t, "-1.6", "-2", 0, m)
	testRounding(t, "-2.5", "-3", 0, m)
	testRounding(t, "-5.5", "-6", 0, m)
}

func TestRoundDown(t *testing.T) {
	m := Down
	testRounding(t, "5.5", "5", 0, m)
	testRounding(t, "2.5", "2", 0, m)
	testRounding(t, "1.6", "1", 0, m)
	testRounding(t, "1.1", "1", 0, m)
	testRounding(t, "1.0", "1", 0, m)
	testRounding(t, "-1.0", "-1", 0, m)
	testRounding(t, "-1.1", "-1", 0, m)
	testRounding(t, "-1.6", "-1", 0, m)
	testRounding(t, "-2.5", "-2", 0, m)
	testRounding(t, "-5.5", "-5", 0, m)
}

func TestRoundFloor(t *testing.T) {
	m := Floor
	testRounding(t, "5.5", "5", 0, m)
	testRounding(t, "2.5", "2", 0, m)
	testRounding(t, "1.6", "1", 0, m)
	testRounding(t, "1.1", "1", 0, m)
	testRounding(t, "1.0", "1", 0, m)
	testRounding(t, "-1.0", "-1", 0, m)
	testRounding(t, "-1.1", "-2", 0, m)
	testRounding(t, "-1.6", "-2", 0, m)
	testRounding(t, "-2.5", "-3", 0, m)
	testRounding(t, "-5.5", "-6", 0, m)
}

func TestRoundHalfUp(t *testing.T) {
	m := HalfUp
	testRounding(t, "5.5", "6", 0, m)
	testRounding(t, "2.5", "3", 0, m)
	testRounding(t, "1.6", "2", 0, m)
	testRounding(t, "1.1", "1", 0, m)
	testRounding(t, "1.0", "1", 0, m)
	testRounding(t, "-1.0", "-1", 0, m)
	testRounding(t, "-1.1", "-1", 0, m)
	testRounding(t, "-1.6", "-2", 0, m)
	testRounding(t, "-2.5", "-3", 0, m)
	testRounding(t, "-5.5", "-6", 0, m)
}

func TestRoundHalfDown(t *testing.T) {
	m := HalfDown
	testRounding(t, "5.5", "5", 0, m)
	testRounding(t, "2.5", "2", 0, m)
	testRounding(t, "1.6", "2", 0, m)
	testRounding(t, "1.1", "1", 0, m)
	testRounding(t, "1.0", "1", 0, m)
	testRounding(t, "-1.0", "-1", 0, m)
	testRounding(t, "-1.1", "-1", 0, m)
	testRounding(t, "-1.6", "-2", 0, m)
	testRounding(t, "-2.5", "-2", 0, m)
	testRounding(t, "-5.5", "-5", 0, m)
}

func TestRoundHalfEven(t *testing.T) {
	m := HalfEven
	testRounding(t, "5.5", "6", 0, m)
	testRounding(t, "2.5", "2", 0, m)
	testRounding(t, "1.6", "2", 0, m)
	testRounding(t, "1.1", "1", 0, m)
	testRounding(t, "1.0", "1", 0, m)
	testRounding(t, "-1.0", "-1", 0, m)
	testRounding(t, "-1.1", "-1", 0, m)
	testRounding(t, "-1.6", "-2", 0, m)
	testRounding(t, "-2.5", "-2", 0, m)
	testRounding(t, "-5.5", "-6", 0, m)
}

func testRounding(t *testing.T, a, b string, prec int, method RoundingMode) {
	x, ok := new(big.Rat).SetString(a)
	if !ok {
		t.Fatalf("Failed to parse: %s", a)
	}
	Round(x, prec, method)
	r := x.FloatString(prec)
	if r != b {
		t.Errorf("test Round(%v, %v, %v) == %s. Got %v", a, prec, method, b, r)
	}
}

func BenchmarkRoundUp(b *testing.B) {
	benchmarkRounding(b, roundUp)
}

func BenchmarkRoundHalfEven(b *testing.B) {
	benchmarkRounding(b, HalfEven)
}

func benchmarkRounding(b *testing.B, mode RoundingMode) {
	r := rand.New(rand.NewSource(rand.Int63()))
	x := new(big.Rat)
	d := big.NewInt(1e10)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x.SetFrac(x.Num().Rand(r, d), d)
		if i%2 == 0 {
			x.Neg(x)
		}
		prec := rand.Intn(10)
		b.StartTimer()

		Round(x, prec, mode)
	}
}
