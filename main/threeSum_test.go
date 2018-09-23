package main

import (
	"sort"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkThreeSum1-4                   1        3408523217 ns/op        2016750480 B/op 12497554 allocs/op
BenchmarkThreeSumMy1-4                 1        2215891398 ns/op        2016750576 B/op 12497555 allocs/op
PASS
ok      command-line-arguments  5.817s
192:main xianjinwang$ go test threeSum_test.go -bench=. -benchmem
goos: darwin
goarch: amd64
BenchmarkThreeSum1-4               30000             49203 ns/op          122032 B/op       1239 allocs/op
BenchmarkThreeSumMy1-4             30000             51227 ns/op          122032 B/op       1239 allocs/op
PASS
ok      command-line-arguments  4.055s


 */

func threeSumMy1(nums []int) [][]int {
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

func threeSum1(nums []int) [][]int {
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

func getSumData() []int {
	d := make([]int, 100)
	for i := 0; i < 100; i++ {
		d[i] = 50 - i
	}
	return d
}

func BenchmarkThreeSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSum1(getSumData())
	}
}

func BenchmarkThreeSumMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSumMy1(getSumData())
	}
}
