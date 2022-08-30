package binaryTree

func postorderTraversal(root *TreeNode) []int {
	var res []int
	postTraversal(root, &res)
	return res
}

func postTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	postTraversal(root.Left, res)
	postTraversal(root.Right, res)
	*res = append(*res, root.Val)
}
