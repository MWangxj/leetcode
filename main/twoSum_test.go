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

func BenchmarkTwosum10001(t *testing.B) {
	t.Log("two sum target is 10001")
	for i := 0; i < t.N; i++ {
		//twoSum1(getData(), 19999)
		twoSum1(getData(), 10001)
	}
}

func BenchmarkTwosumMy10001(t *testing.B) {
	t.Log("two sum my target is 10001")

	for i := 0; i < t.N; i++ {
		//twoSumMy1(getData(), 19999)
		twoSumMy1(getData(), 10001)
	}
}

func BenchmarkTwosum19999(t *testing.B) {
	t.Log("two sum target is 19999")
	for i := 0; i < t.N; i++ {
		twoSum1(getData(), 19999)
		//twoSum1(getData(), 10001)
	}
}

func BenchmarkTwosumMy19999(t *testing.B) {
	t.Log("two sum my target is 19999")

	for i := 0; i < t.N; i++ {
		twoSumMy1(getData(), 19999)
		//twoSumMy1(getData(), 10001)
	}
}

func getData()[]int{
	d :=make([]int,10000)
	for i:=1;i<10000;i++{
		d[i]=i
	}
	return d
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
