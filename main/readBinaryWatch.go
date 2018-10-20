package main

import (
	"fmt"
	"math/bits"
	"strings"
)

/**

A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.


For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

Example:

Input: n = 1
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
Note:

The order of output does not matter.
The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".


8 4 2 1
32 16 8 4 2 1

 */

const (
	MinutesBits = 6
	HoursBits   = 4
	WatchBits   = HoursBits + MinutesBits
	HoursMax    = 12
	MinutesMax  = 60
)

func inRange(n, l, r uint) bool {
	return n >= l && n < r
}

/**
	不是很理解
 */
func readBinaryWatch(num int) (answer []string) {
	var i uint
	for i = 0; i < 2<<WatchBits; i++ {
		if bits.OnesCount(i) == num {
			hours := i >> MinutesBits
			minutes := i & 0x3f
			if inRange(hours, 0, HoursMax) && inRange(minutes, 0, MinutesMax) {
				answer = append(answer, fmt.Sprintf("%d:%02d", hours, minutes))
			}
		}
	}
	return
}

func readBinaryWatchMy(num int) []string {
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

func main() {
	fmt.Println(readBinaryWatch(8))
}
