package binaryTree

func maxDepth(root *TreeNode) int {
	return helpMaxDepth(root)
}

func helpMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helpMaxDepth(root.Left)
	right := helpMaxDepth(root.Right)
	return Max(left, right) + 1
}
