package binaryTree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return IsSymmetricHelp(root.Left, root.Right)
}

func IsSymmetricHelp(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return IsSymmetricHelp(left.Left, right.Right) && IsSymmetricHelp(left.Right, right.Left)
}
