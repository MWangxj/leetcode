package main

import (
	"fmt"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkBinaryGap1-4           30000000                42.6 ns/op             0 B/op          0 allocs/op
BenchmarkBinaryGapMy1-4         10000000               200 ns/op             248 B/op          3 allocs/op
PASS
ok      command-line-arguments  3.542s


 */

func BenchmarkBinaryGap1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryGap1(33556480)
	}
}

func BenchmarkBinaryGapMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryGapMy1(33556480)
	}
}

func binaryGap1(N int) int {
	cur1 := -1  // current position of 1
	prev1 := -1 // previous position of 1
	pos := -1   // current position
	dist := 0   // the longest distance between two consecutive 1's
	for N > 0 {
		digit := N % 2
		pos ++
		if digit == 1 {
			prev1 = cur1
			cur1 = pos
			if prev1 == -1 {
				dist = 0
			} else if cur1-prev1 > dist {
				dist = cur1 - prev1
			}
		}
		N /= 2
	}
	return dist
}

func binaryGapMy1(N int) int {
	max := 0
	b := fmt.Sprintf("%b", N)
	p := make([]int, 0, len(b))
	for i, v := range b {
		if v == '1' {
			p = append(p, i)
		}
	}
	for i := 0; i < len(p)-1; i++ {
		t := p[i+1] - p[i]
		if t > max {
			max = t
		}
	}
	return max
}
