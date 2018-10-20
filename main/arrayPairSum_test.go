package main

import "testing"

/**

goos: darwin
goarch: amd64
BenchmarkArrayPairSumMy-4       100000000               17.4 ns/op             0 B/op          0 allocs/op
BenchmarkArrayPairSum-4            10000            165427 ns/op               0 B/op          0 allocs/op
PASS
ok      command-line-arguments  3.435s



 */

func BenchmarkArrayPairSumMy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayPairSumMy1([]int{1, 4, 3, 2})
	}
}

func BenchmarkArrayPairSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayPairSum1([]int{1, 4, 3, 2})
	}
}

func arrayPairSumMy1(nums []int) int {
	sort1(nums)
	result := 0
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}
func sort1(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func min1(A int, B int) int {
	if A > B {
		return B
	}
	return A
}

func max1(A int, B int) int {
	if A > B {
		return A
	}
	return B
}

func arrayPairSum1(nums []int) int {
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
		minV = min1(minV, v)
		maxV = max1(maxV, v)
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
