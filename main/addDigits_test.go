package main

import "testing"

/**

goos: darwin
goarch: amd64
BenchmarkAddDigits1-4           2000000000               0.28 ns/op            0 B/op          0 allocs/op
BenchmarkAddDigitsMy1-4         50000000                27.9 ns/op             0 B/op          0 allocs/op
PASS
ok      command-line-arguments  2.020s


 */

func addDigitsMy1(num int) int {
	if num < 10 {
		return num
	}
	return addDigitsMy1(num/10 + num%10)
}

func addDigits1(num int) int {
	return 1 + (num-1)%9
}

func BenchmarkAddDigits1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDigits1(123456789)
	}
}

func BenchmarkAddDigitsMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDigitsMy1(123456789)
	}
}
