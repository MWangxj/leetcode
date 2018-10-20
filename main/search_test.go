package main

import "testing"

/**

goos: darwin
goarch: amd64
BenchmarkSearch1-4      200000000                8.63 ns/op            0 B/op          0 allocs/op
BenchmarkSearchMy1-4    200000000                8.28 ns/op            0 B/op          0 allocs/op
PASS
ok      command-line-arguments  5.103s


 */

func BenchmarkSearch1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		search1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	}
}

func BenchmarkSearchMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		searchMy1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9)
	}
}

func search1(nums []int, target int) int {
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

func searchMy1(nums []int, target int) int {
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
