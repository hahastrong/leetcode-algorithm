package binaryTree

import (
	"fmt"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var (
		res []string
	)
	BinaryTreePathsHelp(root, &res, []string{})
	return res
}


func BinaryTreePathsHelp(root *TreeNode, res *[]string, path []string) {
	path = append(path, fmt.Sprintf("%d", root.Val))

	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(path, "->"))
	}

	if root.Left != nil {
		BinaryTreePathsHelp(root.Left, res, path)
	}
	if root.Right != nil {
		BinaryTreePathsHelp(root.Right, res, path)
	}
	path = path[:len(path)-1]
}

