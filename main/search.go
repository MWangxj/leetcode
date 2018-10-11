package main

import "fmt"

/**

Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.


Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1


Note:

You may assume that all elements in nums are unique.
n will be in the range [1, 10000].
The value of each element in nums will be in the range [-9999, 9999].

 */

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func searchMy(nums []int, target int) int {
	l := len(nums)
	pre, sub := 0, l-1
	if pre == sub && nums[pre] != target {
		return -1
	}
	for {
		med := (pre + sub) / 2
		if nums[med] == target {
			return med
		}
		if nums[med+1] == target {
			return med + 1
		}
		if pre == sub || pre == sub-1 {
			return -1
		}
		if nums[med] > target {
			sub = med
			continue
		}
		if nums[med] < target {
			pre = med
			continue
		}

	}

}

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6}, 7))
}
