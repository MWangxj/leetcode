package main

import "fmt"

/**

Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5

 */

func lengthOfLastWordMy(s string) int {
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

func lengthOfLastWord(s string) int {
	var from, to int

	for to = len(s) - 1; (to >= 0) && (' ' == s[to]); to-- {
	}

	for from = to; (from >= 0) && (' ' != s[from]); from-- {
	}

	return to - from
}

func main() {
	fmt.Println(lengthOfLastWord("Hello    "))
}
