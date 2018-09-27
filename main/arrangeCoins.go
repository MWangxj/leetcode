package main

import (
	"fmt"
	"math"
)

/**

You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.

Given n, find the total number of full staircase rows that can be formed.

n is a non-negative integer and fits within the range of a 32-bit signed integer.

Example 1:

n = 5

The coins can form the following rows:
¤
¤ ¤
¤ ¤

Because the 3rd row is incomplete, we return 2.
Example 2:

n = 8

The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤

Because the 4th row is incomplete, we return 3.

 */

func arrangeCoinsMy(n int) int {
	if n < 0 {
		return 0
	}
	i := int(math.Sqrt(float64(n)))
	for {
		if t := i*i + i; t > 2*n {
			i--
		} else if t := (i+1)*(i+1) + i; t < 2*n+1 {
			i++
		} else {
			return int(i)
		}
	}
	return 0
}

func arrangeCoins(n int) int {
	total := float64(n) * 2
	num := int(math.Sqrt(total))
	if num*(num+1) <= int(total) {
		return num
	}
	return num - 1
}

func main() {
	fmt.Println(arrangeCoins(6))
}
