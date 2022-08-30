package binaryTree


func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convertBSTHelp(root, &sum)
	return root
}

func convertBSTHelp(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	convertBSTHelp(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	convertBSTHelp(root.Left, sum)
}