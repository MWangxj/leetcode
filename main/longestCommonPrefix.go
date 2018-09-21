package main

import (
	`fmt`
)

/**

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.

 */

func longestCommonPrefix(strs []string) string {
	if strs==nil ||len(strs)==0{
		return ""
	}
	if len(strs)==1 {
		return strs[0]
	}
	first := strs[0]
	return longest(first,strs[1:])

}

func longest(first string,strs[]string) string {
	if strs==nil {
		return first
	}
	if len(strs)==1 {
		return common(first,strs[0])
	}
	return longest(common(first,strs[0]),strs[1:])
}

func common(str1,str2 string) string {
	s := ""
	l :=0
	if len(str1)<len(str2) {
		l=len(str1)
	}else {
		l = len(str2)
	}
	for i:=0;i<l;i++{
		if str1[i]==str2[i] {
			s +=string(str1[i])
		}else {
			break
		}
	}
	return s
}

func main()  {
	strs := []string{"wangxianjin","wangxian","wangxi"}
	fmt.Println(longestCommonPrefix(strs))
}
