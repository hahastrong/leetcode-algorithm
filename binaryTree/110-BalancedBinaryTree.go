package binaryTree

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, flag := IsBalancedHelp(root)
	return flag
}

func IsBalancedHelp(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	left, flag := IsBalancedHelp(root.Left)
	if !flag {
		return 0, false
	}
	right, flag := IsBalancedHelp(root.Right)
	if !flag {
		return 0, false
	}
	if math.Abs(float64(right - left)) > 1 {
		return 0, false
	}
	return Max(left, right) + 1, true
}
