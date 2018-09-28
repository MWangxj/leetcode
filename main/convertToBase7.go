package main

import (
	"fmt"
	"math"
	"strconv"
)

/**

Given an integer, return its base 7 string representation.

Example 1:

Input: 100
Output: "202"
Example 2:

Input: -7
Output: "-10"

 */

func convertToBase7(num int) string {
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

func convertToBase7My(num int) string {
	sig := 1
	if num < 0 {
		sig = -1
	}
	num = int(math.Abs(float64(num)))
	s := ""
	s = convert7(s, num)
	if sig > 0 {
		return s
	}
	return "-" + s
}

func convert7(s string, num int) string {
	if num < 7 {
		return strconv.Itoa(num) + s
	}
	return convert7(strconv.Itoa(num%7)+s, num/7)
}

func main() {
	fmt.Println(convertToBase7(-7))
}
