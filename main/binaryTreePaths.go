package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**

Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3

 */

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	result := binaryTreePath(root, "")
	return result
}

func binaryTreePath(root *TreeNode, path string) []string {
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
		result = append(result, binaryTreePath(root.Left, path)...)
	}
	if root.Right != nil {
		result = append(result, binaryTreePath(root.Right, path)...)
	}
	return result
}

func binaryTreePathsMy(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	ans := []string{}
	tmp := []string{}
	helper(root, &ans, tmp)
	return ans
}
func helper(root *TreeNode, ans *[]string, tmp []string) {
	tmp = append(tmp, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, strings.Join(tmp, "->"))
		return
	}
	if root.Left != nil {
		helper(root.Left, ans, tmp)
	}
	if root.Right != nil {
		helper(root.Right, ans, tmp)
	}
}

func main() {
	fmt.Println(binaryTreePaths(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 11}}}))
}
