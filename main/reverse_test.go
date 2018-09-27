package main

import (
	"math"
	"strconv"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkReverse-4      300000000                5.33 ns/op            0 B/op          0 allocs/op
BenchmarkReverseMy-4    20000000                57.6 ns/op             6 B/op          2 allocs/op
PASS
ok      command-line-arguments  3.367s


 */

func reverseMy2(x int) int {
	if x < 0 {
		r := rever2(strconv.Itoa(-x))
		i, _ := strconv.Atoi(r)
		if i > 1<<31 {
			return 0
		}
		return -i
	} else {
		r := rever2(strconv.Itoa(x))
		i, _ := strconv.Atoi(r)
		if i > 1<<31 {
			return 0
		}
		return i
	}
}

func rever2(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = s[len(s)-1-i], s[i]
	}
	return string(b)
}

func reverse1(x int) int {
	result := 0
	for x != 0 {
		result = 10*result + x%10
		x /= 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse1(-321)
	}
}

func BenchmarkReverseMy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseMy2(-321)
	}
}
