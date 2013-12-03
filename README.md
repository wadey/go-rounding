go-rounding
===========

[![Build Status](https://travis-ci.org/wadey/go-rounding.png?branch=master)](https://travis-ci.org/wadey/go-rounding)
[![Coverage Status](https://coveralls.io/repos/wadey/go-rounding/badge.png?branch=master)](https://coveralls.io/r/wadey/go-rounding?branch=master)

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

Example output on `go1.2 darwin_amd64` commit f8f259b28f

    BenchmarkFinite           2000000       817 ns/op     234 B/op       5 allocs/op
    BenchmarkFinitePrec       1000000      1775 ns/op     356 B/op       8 allocs/op
    BenchmarkRoundUp          1000000      2639 ns/op     436 B/op      10 allocs/op
    BenchmarkRoundHalfEven    1000000      2628 ns/op     434 B/op      10 allocs/op
    BenchmarkTrunc            5000000       502 ns/op      71 B/op       1 allocs/op
