package binaryTree


func preorderTraversal(root *TreeNode) []int {
	var res []int
	traversal(root, &res)
	return res
}

func traversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	traversal(root.Left, res)
	traversal(root.Right, res)
}
