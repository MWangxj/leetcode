package main

import (
	"math/rand"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkTwosum10001-4     	    5000	    374406 ns/op	  401456 B/op	       4 allocs/op
--- BENCH: BenchmarkTwosum10001-4
	twoSum_test.go:42: two sum target is 10001
	twoSum_test.go:42: two sum target is 10001
	twoSum_test.go:42: two sum target is 10001
BenchmarkTwosumMy10001-4   	    2000	    832418 ns/op	   81936 B/op	       2 allocs/op
--- BENCH: BenchmarkTwosumMy10001-4
	twoSum_test.go:50: two sum my target is 10001
	twoSum_test.go:50: two sum my target is 10001
	twoSum_test.go:50: two sum my target is 10001
BenchmarkTwosum19999-4     	    5000	    356778 ns/op	  401456 B/op	       4 allocs/op
--- BENCH: BenchmarkTwosum19999-4
	twoSum_test.go:59: two sum target is 19999
	twoSum_test.go:59: two sum target is 19999
	twoSum_test.go:59: two sum target is 19999
BenchmarkTwosumMy19999-4   	    2000	    569115 ns/op	   81936 B/op	       2 allocs/op
--- BENCH: BenchmarkTwosumMy19999-4
	twoSum_test.go:67: two sum my target is 19999
	twoSum_test.go:67: two sum my target is 19999
	twoSum_test.go:67: two sum my target is 19999
PASS
ok  	command-line-arguments	6.687s




 */

func BenchmarkTwosum10001(b *testing.B) {
	b.Log("two sum target is 10001")
	for i := 0; i < b.N; i++ {
		//twoSum1(getData(), 19999)
		twoSum1(getData(), 10001)
	}
}

func BenchmarkTwosumMy10001(b *testing.B) {
	b.Log("two sum my target is 10001")

	for i := 0; i < b.N; i++ {
		//twoSumMy1(getData(), 19999)
		twoSumMy1(getData(), 10001)
	}
}

func BenchmarkTwosum19999(b *testing.B) {
	b.Log("two sum target is 19999")
	for i := 0; i < b.N; i++ {
		twoSum1(getData(), 19999)
		//twoSum1(getData(), 10001)
	}
}

func BenchmarkTwosumMy19999(b *testing.B) {
	b.Log("two sum my target is 19999")

	for i := 0; i < b.N; i++ {
		twoSumMy1(getData(), 19999)
		//twoSumMy1(getData(), 10001)
	}
}

func getData() []int {
	d := make([]int, 10000)
	for i := 1; i < 10000; i++ {
		d[i] = rand.Intn(100000)
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
	hashMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, ok := hashMap[num]; ok {
			return []int{index, i}
		} else {
			hashMap[target-num] = i
		}
	}
	return []int{}
}
