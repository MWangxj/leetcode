package main

import "fmt"

/**

Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

Example:

Input: 38
Output: 2
Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
             Since 2 has only one digit, return it.
Follow up:
Could you do it without any loop/recursion in O(1) runtime?

 */

func addDigitsMy(num int) int {
	if num < 10 {
		return num
	}
	return addDigitsMy(num/10 + num%10)
}

func addDigits(num int) int {
	return 1 + (num-1)%9
}

func main() {
	fmt.Println(addDigits(38))
}
