package binaryTree

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxId, res := 0, 0
	BottomTraversal(root, 1, &maxId, &res)
	return res
}

func BottomTraversal(root *TreeNode, idx int, maxIdx *int, res *int) {
	if root.Left == nil && root.Right == nil {
		if idx > *maxIdx {
			*res = root.Val
			*maxIdx = idx
		}
		return
	}

	if root.Left != nil {
		BottomTraversal(root.Left, idx+1, maxIdx, res)
	}
	if root.Right != nil {
		BottomTraversal(root.Right, idx+1, maxIdx, res)
	}
}
