package main

import (
	`fmt`
)

/**

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].


 */

func twoSumMy(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		if index, ok := hashMap[num]; ok {
			return []int{index, i}
		} else {
			hashMap[target-num] = i
		}
	}
	return []int{}
}


func main(){
	fmt.Println(twoSum([]int{1,5,7,8},9))
}




