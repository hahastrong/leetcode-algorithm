package binaryTree

func inorderTraversal(root *TreeNode) []int {
	var res []int
	InorderTraversal(root, &res)
	return res
}

func InorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	InorderTraversal(root.Left, res)
	*res = append(*res, root.Val)
	InorderTraversal(root.Right, res)

}
