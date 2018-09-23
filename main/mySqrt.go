package main

import (
	"fmt"
	"math"
)

func mySqrtMy(x int) int {
	r:= math.Sqrt(float64(x))
	return int(r)
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left := 0
	right := x
	for {
		mid := (left + right) / 2
		if (mid * mid <= x) && ((mid+1) * (mid+1) > x) {
			return mid
		} else if mid * mid < x {
			left = mid
		} else {
			right = mid
		}
	}
	return 0
}

func main()  {
	fmt.Println(mySqrtMy(8))
}
