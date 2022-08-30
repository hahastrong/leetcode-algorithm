package binaryTree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return helpMinDepth(root)
}

func helpMinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helpMinDepth(root.Left)
	right := helpMinDepth(root.Right)

	if left == 0 {
		return right + 1
	}
	if right == 0 {
		return left + 1
	}
	return Min(left, right) + 1
}
