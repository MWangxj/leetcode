package main

import (
	"strconv"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkAddStrings1-4     	10000000	       131 ns/op	      32 B/op	       1 allocs/op
BenchmarkAddStringsMy1-4   	 1000000	      1379 ns/op	     672 B/op	      31 allocs/op
PASS
ok  	command-line-arguments	2.850s

 */

func BenchmarkAddStrings1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addStrings1("3781616348164871678654310657", "17683561364580165046713056105634")
	}
}

func BenchmarkAddStringsMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addStringsMy1("3781616348164871678654310657", "17683561364580165046713056105634")
	}
}

func addStringsMy1(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	var c uint8
	var s string
	for {
		if i >= 0 || j >= 0 {
			c /= 10
			if i >= 0 {
				c += num1[i] - '0'
				i--
			}
			if j >= 0 {
				c += num2[j] - '0'
				j--
			}
			s = strconv.Itoa(int(c%10)) + s
		} else {
			break
		}
	}
	if c > 9 {
		s = "1" + s
	}
	return s
}

func addStrings1(num1 string, num2 string) string {
	var a, b []byte
	if len(num1) > len(num2) {
		a = []byte(num1)
		b = []byte(num2)
	} else {
		b = []byte(num1)
		a = []byte(num2)
	}
	var l int
	if len(a) > len(b) {
		l = len(a)
	} else {
		l = len(b)
	}
	la := len(a) - 1
	lb := len(b) - 1
	var bl byte = 0
	for i := 0; i < l; i++ {
		if lb-i >= 0 {
			a[la-i] = a[la-i] + b[lb-i] + bl - 48*2
		} else {
			a[la-i] = a[la-i] + bl - 48
		}

		bl = a[la-i] / 10
		a[la-i] = (a[la-i] % 10) + 48
	}

	if bl != 0 {
		return "1" + string(a)
	}
	return string(a)
}
