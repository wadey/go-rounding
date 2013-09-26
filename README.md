go-rounding
===========

[![Build Status](https://travis-ci.org/wadey/go-rounding.png?branch=master)](https://travis-ci.org/wadey/go-rounding)

Rounding and truncation methods for big.Rat to close some of the gap in
functionality between Rat and Decimal (such as Java's BigDecimal).

Documentation on GoDoc: <http://godoc.org/github.com/wadey/go-rounding>.

Why go-rounding?
----------------

There are a few other Decimal implementations for Go:

- [dec](http://godoc.org/code.google.com/p/godec/dec)
- [inf](http://godoc.org/speter.net/go/exp/math/dec/inf)

So why does go-rounding exist? big.Rat is a superset of the basic needs for
a Decimal representation, the only major thing missing is some nice rounding
methods. go-rounding is less than 150 lines of code, so it is easy to review
and understand its implementation. If you need more features, you should use
one of the above packages.

Benchmark
---------

To run:

    go test -v -benchmem -bench .

Example output on `go1.1.2 darwin/amd64` commit cde8d46

    BenchmarkFinite           2000000           796 ns/op         238 B/op          5 allocs/op
    BenchmarkFinitePrec       1000000          1373 ns/op         368 B/op          8 allocs/op
    BenchmarkRoundUp           500000          3160 ns/op         439 B/op         10 allocs/op
    BenchmarkRoundHalfEven     500000          3170 ns/op         438 B/op         10 allocs/op
    BenchmarkTrunc            5000000           406 ns/op          74 B/op          1 allocs/op
