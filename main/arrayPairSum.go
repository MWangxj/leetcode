package main

import "fmt"

/**

Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

Example 1:

Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).
Note:

n is a positive integer, which is in the range of [1, 10000].
All the integers in the array will be in the range of [-10000, 10000].


 */

func arrayPairSumMy(nums []int) int {
	sort(nums)
	result := 0
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}
func sort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func min(A int, B int) int {
	if A > B {
		return B
	}
	return A
}

func max(A int, B int) int {
	if A > B {
		return A
	}
	return B
}

func arrayPairSum(nums []int) int {
	result := 0
	flag := true
	check_map := make(map[int]int)
	minV := -10000
	maxV := 10000
	for _, v := range nums {
		cnt, ok := check_map[v]
		if ok {
			check_map[v] = cnt + 1
		} else {
			check_map[v] = 1
		}
		minV = min(minV, v)
		maxV = max(maxV, v)
	}
	//fmt.Println(check_map)
	for i := minV; i <= maxV; i++ {
		cnt, ok := check_map[i]
		if !ok {
			continue
		}
		for j := cnt; j > 0; j-- {
			if flag {
				result += i
			}
			flag = !flag
		}
	}
	return result

}

func main() {
	a := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(a))
}
