package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	length := len(s)
	result, i, prev, sign := 0, 0, 0, "+"
	for i < len(s) {
		num := 0
		for i < length && unicode.IsDigit(rune(s[i])) {
			if n, err := strconv.Atoi(string(s[i])); err == nil {
				num = num*10 + n
				i++
			}
		}
		if sign == "+" {
			result += prev
			prev = num
		} else if sign == "-" {
			result += prev
			prev = -num
		} else if sign == "*" {
			prev *= num
		} else if sign == "/" {
			prev /= num
		}

		if i < length {
			sign = string(s[i])
			i++
		}
	}
	return result + prev
}

func main() {
	fmt.Println(calculate("8+5*2*4"))
}
