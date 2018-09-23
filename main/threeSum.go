package main

import (
	"sort"
)

/**

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

 */

func threeSumMy(nums []int) [][]int {
	var results [][]int
	if len(nums) < 3 {
		return results
	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {

		if nums[i] > 0 {
			return results
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			b, c := nums[j], nums[k]
			if c < 0 {
				break
			}

			sum := nums[i] + b + c

			if sum == 0 {
				results = append(results, []int{nums[i], b, c})
			}

			if sum >= 0 {
				for j < k && c == nums[k] && nums[k] >= 0 {
					k--
				}
			}
			if sum <= 0 && nums[k] >= 0 {
				for j < k && b == nums[j] {
					j++
				}
			}
		}
	}

	return results
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	var res [][]int
	sort.Ints(nums)
	for i, first := range nums[:n-2] {
		if i > 0 && first == nums[i-1] {
			continue
		}
		sec, last := i+1, n-1;
		for sec < last {
			sum := first + nums[sec] + nums[last];
			if sum == 0 {
				res = append(res, []int{first, nums[sec], nums[last]})
				sec, last = sec+1, last-1
				for sec < last && nums[sec] == nums[sec-1] {
					sec++
				}

				for sec < last && nums[last] == nums[last+1] {
					last--
				}
			} else if sum < 0 {
				sec++
			} else {
				last--
			}
		}
	}
	return res
}

func main() {

}
