package binaryTree

func invertTree(root *TreeNode) *TreeNode {
	helpInvertTree(root)
	return root
}

func helpInvertTree(root *TreeNode) {
	if root == nil {
		return
	}

	root.Right,root.Left = root.Left, root.Right
	helpInvertTree(root.Right)
	helpInvertTree(root.Left)
}
