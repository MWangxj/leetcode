package main

import (
	"math"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkArrangeCoins1-4        2000000000               0.28 ns/op            0 B/op          0 allocs/op
BenchmarkArrangeCoinsMy1-4      20000000                59.2 ns/op             0 B/op          0 allocs/op
PASS
ok      command-line-arguments  1.849s

 */

func arrangeCoinsMy1(n int) int {
	if n < 0 {
		return 0
	}
	i := int(math.Sqrt(float64(n)))
	for {
		if t := i*i + i; t > 2*n {
			i--
		} else if t := (i+1)*(i+1) + i; t < 2*n+1 {
			i++
		} else {
			return int(i)
		}
	}
	return 0
}

func arrangeCoins1(n int) int {
	total := float64(n) * 2
	num := int(math.Sqrt(total))
	if num*(num+1) <= int(total) {
		return num
	}
	return num - 1
}

func BenchmarkArrangeCoins1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrangeCoins1(10000)
	}
}

func BenchmarkArrangeCoinsMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrangeCoinsMy1(10000)
	}
}
