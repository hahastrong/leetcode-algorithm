package binaryTree

import "math"

func getMinimumDifference(root *TreeNode) int {
	var pre, min int = math.MinInt32, math.MaxInt32
	GetMinimumDifferenceHelp(root, &pre, &min)
	return min
}

func GetMinimumDifferenceHelp(root *TreeNode, pre, min *int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		GetMinimumDifferenceHelp(root.Left, pre, min)
	}
	if *pre >= 0 && (root.Val - *pre) < *min {
		*min = root.Val - *pre
	}

	*pre = root.Val

	if root.Right != nil {
		GetMinimumDifferenceHelp(root.Right, pre, min)
	}
}
