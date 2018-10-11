package main

import "fmt"

/**

Given two strings A and B of lowercase letters, return true if and only if we can swap two letters in A so that the result equals B.



Example 1:

Input: A = "ab", B = "ba"
Output: true
Example 2:

Input: A = "ab", B = "ab"
Output: false
Example 3:

Input: A = "aa", B = "aa"
Output: true
Example 4:

Input: A = "aaaaaaabc", B = "aaaaaaacb"
Output: true
Example 5:

Input: A = "", B = "aa"
Output: false


Note:

0 <= A.length <= 20000
0 <= B.length <= 20000
A and B consist only of lowercase letters.

 */
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	c, ls := 0, make([]uint8, 0, 4)
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			c++
			ls = append(ls, []uint8{A[i], B[i]}...)
		}
		if c > 2 {
			return false
		}
	}
	if len(ls) == 0 {
		return true
	}
	if len(ls) != 4 {
		return false
	}
	if len(ls) == 4 && ls[0] == ls[3] && ls[1] == ls[2] {
		return true
	}
	return false
}

func main() {
	fmt.Println(buddyStrings("ab", "ab"))
}
