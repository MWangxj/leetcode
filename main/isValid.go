package main

import "fmt"

/**

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true

 */

func isValid(s string) bool {
	if len(s)%2!=0 {
		return false
	}
	ch := make([]string, len(s))
	i :=0
	for _, c := range s {
		if c == int32([]byte("(")[0]) {
			ch[i]=")"
			i++
		} else if c == int32([]byte("[")[0]) {
			ch[i]="]"
			i++
		}else if c == int32([]byte("{")[0]) {
			ch[i]="}"
			i++
		}else {
			if i==0 {
				return false
			}
			i--
			v := int32([]byte(ch[i])[0])
			if len(ch)>1 {
				ch = append(ch[:i],ch[i+1:]...)
			}
			if  v !=c {
				return false
			}
		}
	}
	return len(ch)==0||ch[0]==""
}

func main() {
	fmt.Println(isValid("(("))
}
