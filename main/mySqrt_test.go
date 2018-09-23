package main

import (
	"math"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkMysqrt1-4      500000000                3.77 ns/op            0 B/op          0 allocs/op
BenchmarkMysqrtMy1-4    2000000000               0.28 ns/op            0 B/op          0 allocs/op
PASS
ok      command-line-arguments  2.876s


 */

func mySqrtMy1(x int) int {
	r:= math.Sqrt(float64(x))
	return int(r)
}

func mySqrt1(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left := 0
	right := x
	for {
		mid := (left + right) / 2
		if (mid * mid <= x) && ((mid+1) * (mid+1) > x) {
			return mid
		} else if mid * mid < x {
			left = mid
		} else {
			right = mid
		}
	}
	return 0
}

func BenchmarkMysqrt1(b *testing.B)  {
	for i:=0;i<b.N;i++{
		mySqrt1(8)
	}
}

func BenchmarkMysqrtMy1(b *testing.B)  {
	for i:=0;i<b.N;i++{
		mySqrtMy1(8)
	}
}
