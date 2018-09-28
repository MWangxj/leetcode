package main

import "testing"

/**

goos: darwin
goarch: amd64
BenchmarkBackspaceCompare1-4            100000000               12.6 ns/op             0 B/op          0 allocs/op
BenchmarkBackspaceCompareMy1-4            500000              2200 ns/op             928 B/op         51 allocs/op
PASS
ok      command-line-arguments  2.406s


 */

func BenchmarkBackspaceCompare1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backspaceCompare1("sdgfasgfgqugofgeyegqufg#######asdfasf", "dyaadsuigfyuegquygasdhgfaosgd#####sdgf")
	}
}

func BenchmarkBackspaceCompareMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backspaceCompareMy1("sdgfasgfgqugofgeyegqufg#######asdfasf", "dyaadsuigfyuegquygasdhgfaosgd#####sdgf")
	}
}

func backspaceCompare1(S string, T string) bool {
	slen, tlen := len(S), len(T)
	if slen == 0 && tlen == 0 {
		return true
	}
	scur, tcur := slen-1, tlen-1
	sbs, tbs := 0, 0

	for {
		for scur >= 0 && (S[scur] == '#' || sbs > 0) {
			if S[scur] == '#' {
				sbs++
			} else {
				sbs--
			}
			scur--
		}
		for tcur >= 0 && (T[tcur] == '#' || tbs > 0) {
			if T[tcur] == '#' {
				tbs++
			} else {
				tbs--
			}
			tcur--
		}

		if (scur == 0 && tcur == 0) || (scur < 0 && tcur < 0) {
			return true
		} else if (scur >= 0 && tcur < 0) || (scur < 0 && tcur >= 0) || S[scur] != T[tcur] {
			return false
		}
		scur--
		tcur--
	}
}

func backspaceCompareMy1(S string, T string) bool {
	s := ""
	istep := 0
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '#' {
			istep++
			continue
		}
		if istep > 0 {
			istep--
			continue
		}
		s = string(S[i]) + s
	}

	t := ""
	lstep := 0
	for l := len(T) - 1; l >= 0; l-- {
		if T[l] == '#' {
			lstep++
			continue
		}
		if lstep > 0 {
			lstep--
			continue
		}
		t = string(T[l]) + t
	}
	return s == t
}
