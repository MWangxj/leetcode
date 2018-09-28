package main

import "fmt"

/**

Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:

Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
Example 2:

Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
Example 3:

Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
Example 4:

Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
Note:

1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.
Follow up:

Can you solve it in O(N) time and O(1) space?

 */

func backspaceCompare(S string, T string) bool {
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

func backspaceCompareMy(S string, T string) bool {
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

func main() {
	fmt.Println(backspaceCompare("cad#b##", "a#p#"))
}
