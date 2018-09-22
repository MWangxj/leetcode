package main

import "testing"

/**

192:main xianjinwang$ go test twoSum_test.go -bench=. -benchmem
goos: darwin
goarch: amd64
BenchmarkTwosum10001-4              5000            271090 ns/op          401514 B/op          6 allocs/op
--- BENCH: BenchmarkTwosum10001-4
        twoSum_test.go:19: two sum target is 10001
        twoSum_test.go:19: two sum target is 10001
        twoSum_test.go:19: two sum target is 10001
BenchmarkTwosumMy10001-4           50000             28983 ns/op           81936 B/op          2 allocs/op
--- BENCH: BenchmarkTwosumMy10001-4
        twoSum_test.go:27: two sum my target is 10001
        twoSum_test.go:27: two sum my target is 10001
        twoSum_test.go:27: two sum my target is 10001
        twoSum_test.go:27: two sum my target is 10001
BenchmarkTwosum19999-4              3000            528345 ns/op          404178 B/op         12 allocs/op
--- BENCH: BenchmarkTwosum19999-4
        twoSum_test.go:36: two sum target is 19999
        twoSum_test.go:36: two sum target is 19999
        twoSum_test.go:36: two sum target is 19999
BenchmarkTwosumMy19999-4              50          28169564 ns/op           81941 B/op          1 allocs/op
--- BENCH: BenchmarkTwosumMy19999-4
        twoSum_test.go:44: two sum my target is 19999
        twoSum_test.go:44: two sum my target is 19999
PASS
ok      command-line-arguments  6.235s




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
