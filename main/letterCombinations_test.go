package main

import "testing"

/**

goos: darwin
goarch: amd64
BenchmarkLetterCombinations1-4              1000           1189410 ns/op         1177555 B/op      15992 allocs/op
BenchmarkLetterCombinationsMy1-4            2000            611461 ns/op          309128 B/op      12776 allocs/op
PASS
ok      command-line-arguments  2.612s


 */

func BenchmarkLetterCombinations1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations1("23456789")
	}
}

func BenchmarkLetterCombinationsMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinationsMy1("23456789")
	}
}

func letterCombinations1(digits string) []string {
	res := make([]string, 0)
	digitWordsMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	for _, digit := range digits {
		words := digitWordsMap[string(digit)]
		tmp := make([]string, 0)
		for _, word := range words {
			if len(res) > 0 {
				for _, item := range res {
					tmp = append(tmp, item+string(word))
				}
			} else {
				tmp = append(tmp, string(word))
			}
		}
		res = tmp
	}
	return res
}

func letterCombinationsMy1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	mm := make([][]string, len(digits))
	m := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	for i, v := range digits {
		if letter, ok := m[v]; ok {
			mm[i] = letter
		} else {
			return []string{}
		}
	}
	if len(digits) == 1 {
		return mm[0]
	}
	mutArray2Array1(mm, &res)
	return res
}

func mutArray2Array1(mut [][]string, res *[]string) {
	if len(mut) < 2 {
		if len(mut) == 1 {
			len1, len2 := len(*res), len(mut[0])
			newLen := len1 * len2
			temp := make([]string, newLen)
			index := 0
			for _, v0 := range *res {
				for _, v1 := range mut[0] {
					temp[index] = v0 + v1
					index++
				}
			}
			*res = append(temp)
		}
		return
	}
	len1, len2 := len(mut[0]), len(mut[1])
	newLen := len1 * len2
	temp := make([]string, newLen)
	index := 0
	for _, v0 := range mut[0] {
		for _, v1 := range mut[1] {
			temp[index] = v0 + v1
			index++
		}
	}
	if len(*res) == 0 {
		*res = append(*res, temp...)
	} else {
		newLen *= len(*res)
		index = 0
		temp1 := make([]string, newLen)
		for _, v0 := range *res {
			for _, v1 := range temp {
				temp1[index] = v0 + v1
				index++
			}
		}
		*res = append(temp1)
	}
	mut = mut[2:]
	mutArray2Array1(mut, res)
}
