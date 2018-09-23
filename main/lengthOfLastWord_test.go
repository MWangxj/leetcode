package main

import "testing"

/**

goos: darwin
goarch: amd64
BenchmarkLengthOfLastWord-4             500000000                3.30 ns/op            0 B/op          0 allocs/op
BenchmarkLengthOfLastWordMy-4           50000000                32.1 ns/op             0 B/op          0 allocs/op
PASS
ok      command-line-arguments  3.641s


 */

func lengthOfLastWordMy1(s string) int {
	l := 0
	for i, v := range s {
		if v != 32 {
			if i > 1 {
				last := s[i-1]
				if last == ' ' {
					l = 1
					continue
				}
			}
			l++
		} else {
			continue
		}
	}
	return l
}

func lengthOfLastWord1(s string) int {
	var from, to int

	for to = len(s) - 1; (to >= 0) && (' ' == s[to]); to-- {
	}

	for from = to; (from >= 0) && (' ' != s[from]); from-- {
	}

	return to - from
}

func BenchmarkLengthOfLastWord(t *testing.B) {
	for i := 0; i < t.N; i++ {
		lengthOfLastWord1("w a n g x i a n j i n")
	}
}

func BenchmarkLengthOfLastWordMy(t *testing.B) {
	for i := 0; i < t.N; i++ {
		lengthOfLastWordMy1("w a n g x i a n j i n")
	}
}
