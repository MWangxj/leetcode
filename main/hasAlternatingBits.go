package main

import "fmt"

/**


Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

Example 1:

Input: 5
Output: True
Explanation:
The binary representation of 5 is: 101
Example 2:

Input: 7
Output: False
Explanation:
The binary representation of 7 is: 111.
Example 3:

Input: 11
Output: False
Explanation:
The binary representation of 11 is: 1011.
Example 4:

Input: 10
Output: True
Explanation:
The binary representation of 10 is: 1010.

 */
func hasAlternatingBits(n int) bool {
	prev := n & 1
	n = n >> 1
	for n > 0 {
		t := n & 1
		if t == prev {
			return false
		}
		prev = t
		n = n >> 1
	}
	return true
}

func hasAlternatingBitsMy(n int) bool {
	b := fmt.Sprintf("%b", n)
	if len(b) < 2 {
		return true
	}
	for i := 0; i < len(b)-1; i++ {
		if b[i] == b[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(11))
}
