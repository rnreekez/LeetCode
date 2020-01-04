package main

import (
	"fmt"
	"github.com/rnreekez/LeetCode/kit"
)

// TreeNode for use by leetcode
type TreeNode = kit.TreeNode

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {
	// [3,9,20,null,null,15,7]
	// depth = 3
	tree := kit.SliceToTree([]int{3, 9, 20, kit.NULL, kit.NULL, 15, 7})
	fmt.Printf("%+v\n", maxDepth(tree))
}
