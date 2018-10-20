package main

import (
	"math"
	"strconv"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkConvertToBase71-4     	20000000	        92.0 ns/op	      40 B/op	       3 allocs/op
BenchmarkConvertToBase7My1-4   	 5000000	       253 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	command-line-arguments	3.464s

*/

func convertToBase71(num int) string {
	if num == 0 {
		return "0"
	}
	arr := []byte{}
	neg := false
	if num < 0 {
		neg = true
		num *= -1
	}
	for num != 0 {
		arr = append(arr, byte(num%7)+'0')
		num /= 7
	}
	if neg {
		arr = append(arr, '-')
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}

func convertToBase7My1(num int) string {
	sig := 1
	if num < 0 {
		sig = -1
	}
	num = int(math.Abs(float64(num)))
	s := ""
	s = convert71(s, num)
	if sig > 0 {
		return s
	}
	return "-" + s
}

func convert71(s string, num int) string {
	if num < 7 {
		return strconv.Itoa(num) + s
	}
	return convert71(strconv.Itoa(num%7)+s, num/7)
}

func BenchmarkConvertToBase71(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertToBase71(123456789)
	}
}

func BenchmarkConvertToBase7My1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertToBase7My1(123456789)
	}
}
