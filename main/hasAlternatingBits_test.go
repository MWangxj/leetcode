package main

import (
	"fmt"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkHasAlternatingBits1-4          1000000000               2.53 ns/op            0 B/op          0 allocs/op
BenchmarkHasAlternatingBitsMy1-4        20000000                99.3 ns/op            24 B/op          2 allocs/op
PASS
ok      command-line-arguments  4.884s

 */

func BenchmarkHasAlternatingBits1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hasAlternatingBits1(12345)
	}
}

func BenchmarkHasAlternatingBitsMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hasAlternatingBitsMy1(12345)
	}
}

func hasAlternatingBits1(n int) bool {
	prev := n & 1
	n = n >> 1
	for n > 0 {
		t := n & 1
		if t == prev {
			return false
		}
		prev = t
		n = n >> 1
	}
	return true
}

func hasAlternatingBitsMy1(n int) bool {
	b := fmt.Sprintf("%b", n)
	if len(b) < 2 {
		return false
	}
	for i := 0; i < len(b)-1; i++ {
		if b[i] == b[i+1] {
			return false
		}
	}
	return true
}
