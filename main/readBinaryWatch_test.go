package main

import (
	"fmt"
	"math/bits"
	"strings"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkReadBinaryWatch1-4               200000             11046 ns/op            2928 B/op        119 allocs/op
BenchmarkReadBinaryWatchMy1-4              20000             78357 ns/op           14434 B/op       1570 allocs/op
PASS
ok      command-line-arguments  4.695s

 */

func BenchmarkReadBinaryWatch1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readBinaryWatch1(2)
	}
}

func BenchmarkReadBinaryWatchMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readBinaryWatchMy1(2)
	}
}

const (
	MinutesBits1 = 6
	HoursBits1   = 4
	WatchBits1   = HoursBits1 + MinutesBits1
	HoursMax1    = 12
	MinutesMax1  = 60
)

func inRange1(n, l, r uint) bool {
	return n >= l && n < r
}

/**
	不是很理解
 */
func readBinaryWatch1(num int) (answer []string) {
	var i uint
	for i = 0; i < 2<<WatchBits1; i++ {
		if bits.OnesCount(i) == num {
			hours := i >> MinutesBits1
			minutes := i & 0x3f
			if inRange1(hours, 0, HoursMax1) && inRange1(minutes, 0, MinutesMax1) {
				answer = append(answer, fmt.Sprintf("%d:%02d", hours, minutes))
			}
		}
	}
	return
}

func readBinaryWatchMy1(num int) []string {
	s := []string{}
	for i := 0; i < 12; i++ {
		nbh := strings.Count(fmt.Sprintf("%b", i), "1")
		for j := 0; j < 60; j++ {
			nbm := strings.Count(fmt.Sprintf("%b", j), "1")
			if num == nbh+nbm {
				s = append(s, fmt.Sprintf("%d:%02d", i, j))
			}

		}
	}
	return s
}
