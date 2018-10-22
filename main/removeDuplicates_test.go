package main

import (
	"math/rand"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkRemoveDuplicatesMy1-4          2000000000               0.00 ns/op            0 B/op          0 allocs/op
BenchmarkRemoveDuplicates1-4               20000             70821 ns/op           16384 B/op          1 allocs/op
PASS
ok      command-line-arguments  2.136s

 */

func BenchmarkRemoveDuplicatesMy1(b *testing.B)  {
	removeDuplicatesMy1(getDataRemoveDuplicates())
}

func BenchmarkRemoveDuplicates1(b *testing.B)  {
	for i:=0;i<b.N;i++{
		removeDuplicates1(getDataRemoveDuplicates())
	}
}


// getDataRemoveDuplicates 数组错误   不应该无序
func getDataRemoveDuplicates() []int {
	data :=make([]int,2000)
	for i:=0;i<2000;i++ {
		data[i]=rand.Intn(1<<20)
	}
	return data
}

func removeDuplicatesMy1(nums []int) int {
	m := make(map[int]int, 0)
	for i :=0;i< len(nums);i++ {
		if _, ok := m[nums[i]]; ok {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}
		m[nums[i]] = 0
	}
	return len(m)
}

func removeDuplicates1(nums []int) int {
	if len(nums)<2{
		return len(nums)
	}

	index, previous:= 1,0

	for i:=1; i<len(nums); i++{
		if nums[i] != nums[previous]{
			nums[index]= nums[i]
			previous = i
			index++
		}
	}

	return index
}