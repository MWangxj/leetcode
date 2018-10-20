package main

import "testing"

/**

goos: darwin
goarch: amd64
BenchmarkIsValid-4      20000000                82.1 ns/op            16 B/op          2 allocs/op
BenchmarkIsValidMy-4    10000000               203 ns/op              96 B/op          1 allocs/op
PASS
ok      command-line-arguments  3.980s


 */

func isValidMy(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	ch := make([]string, len(s))
	i := 0
	for _, c := range s {
		if c == int32([]byte("(")[0]) {
			ch[i] = ")"
			i++
		} else if c == int32([]byte("[")[0]) {
			ch[i] = "]"
			i++
		} else if c == int32([]byte("{")[0]) {
			ch[i] = "}"
			i++
		} else {
			if i == 0 {
				return false
			}
			i--
			v := int32([]byte(ch[i])[0])
			if len(ch) > 1 {
				ch = append(ch[:i], ch[i+1:]...)
			}
			if v != c {
				return false
			}
		}
	}
	return len(ch) == 0 || ch[0] == ""
}

func isValid1(s string) bool {
	var stack []rune

	// add all parans/braces/brackets to list
	for i := 0; i < len(s); i++ {
		curr := rune(s[i])
		if curr == '(' || curr == '[' || curr == '{' {
			//push
			stack = append(stack, 0)
			copy(stack[1:], stack)
			stack[0] = curr
		}

		if len(stack) == 0 {
			return false
		}

		if curr == ')' || curr == ']' || curr == '}' {
			// pop
			ch := stack[0]
			stack = stack[1:]

			if curr == ')' {
				if ch == '[' || ch == '{' {
					return false
				}
			} else if curr == ']' {
				if ch == '(' || ch == '{' {
					return false
				}
			} else if curr == '}' {
				if ch == '(' || ch == '[' {
					return false
				}
			}
		}

	}

	return len(stack) == 0
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValid1("[]{}()")
	}
}

func BenchmarkIsValidMy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidMy("[]{}()")
	}
}
