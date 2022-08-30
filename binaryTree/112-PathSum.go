package binaryTree

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return PathSumTraversal(root, targetSum)
}

func PathSumTraversal(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	return PathSumTraversal(root.Left, targetSum - root.Val) || PathSumTraversal(root.Right, targetSum - root.Val)
}
