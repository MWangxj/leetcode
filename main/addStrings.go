package main

import (
	"fmt"
	"strconv"
)

/**

Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.

 */

func addStringsMy(num1 string, num2 string) string {
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

func addStrings(num1 string, num2 string) string {
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

func main() {
	fmt.Println(addStrings("123456", "654321"))
}
