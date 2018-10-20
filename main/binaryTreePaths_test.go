package main

import (
	"strconv"
	"strings"
	"testing"
)

/**

goos: darwin
goarch: amd64
BenchmarkBinaryTreePaths1-4              5000000               374 ns/op             112 B/op          8 allocs/op
BenchmarkBinaryTreePathsMy1-4            3000000               485 ns/op             208 B/op          8 allocs/op
PASS
ok      command-line-arguments  4.212s

 */

func BenchmarkBinaryTreePaths1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryTreePaths1(&TreeNode1{Val: 1, Left: &TreeNode1{Val: 2}, Right: &TreeNode1{Val: 6, Right: &TreeNode1{Val: 11}}})
	}
}

func BenchmarkBinaryTreePathsMy1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryTreePathsMy1(&TreeNode1{Val: 1, Left: &TreeNode1{Val: 2}, Right: &TreeNode1{Val: 6, Right: &TreeNode1{Val: 11}}})
	}
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func binaryTreePaths1(root *TreeNode1) []string {
	if root == nil {
		return []string{}
	}
	result := binaryTreePath1(root, "")
	return result
}

func binaryTreePath1(root *TreeNode1, path string) []string {
	rootVal := strconv.Itoa(root.Val)
	if path == "" {
		path = rootVal
	} else {
		path += "->" + rootVal
	}
	if root.Left == nil && root.Right == nil {
		// Leaf node.
		return []string{path}
	}
	result := make([]string, 0)
	if root.Left != nil {
		result = append(result, binaryTreePath1(root.Left, path)...)
	}
	if root.Right != nil {
		result = append(result, binaryTreePath1(root.Right, path)...)
	}
	return result
}

func binaryTreePathsMy1(root *TreeNode1) []string {
	if root == nil {
		return nil
	}
	ans := []string{}
	tmp := []string{}
	helper1(root, &ans, tmp)
	return ans
}
func helper1(root *TreeNode1, ans *[]string, tmp []string) {
	tmp = append(tmp, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, strings.Join(tmp, "->"))
		return
	}
	if root.Left != nil {
		helper1(root.Left, ans, tmp)
	}
	if root.Right != nil {
		helper1(root.Right, ans, tmp)
	}
}
