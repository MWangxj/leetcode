package main

import "testing"

/**

goos: darwin
goarch: amd64
BenchmarkTwosum-4       20000000                94.3 ns/op            16 B/op          1 allocs/op
BenchmarkTwosumMy-4     100000000               23.0 ns/op            16 B/op          1 allocs/op
PASS
ok      command-line-arguments  4.320s



 */

func BenchmarkTwosum(t *testing.B) {
	for i := 0; i < t.N; i++ {
		twoSum1([]int{1, 3, 7, 8}, 9)
	}
}

func BenchmarkTwosumMy(t *testing.B) {
	for i := 0; i < t.N; i++ {
		twoSumMy1([]int{1, 3, 7, 8}, 9)
	}
}

func twoSumMy1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	hashMap := make(map[int]int,len(nums))
	for i, num := range nums {
		if index, ok := hashMap[num]; ok {
			return []int{index, i}
		} else {
			hashMap[target-num] = i
		}
	}
	return []int{}
}
