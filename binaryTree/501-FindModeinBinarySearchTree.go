package binaryTree


var count, maxCount int = 0, 1
var pre *TreeNode
var res = make([]int, 0)

func findMode(root *TreeNode) []int {
	findModeTraversal(root)
	return res
}

func findModeTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		findModeTraversal(root.Left)
	}

	if pre != nil && pre.Val == root.Val {
		count++
	} else {
		count = 1
	}

	pre = root

	if count == maxCount {
		res = append(res, root.Val)
	}

	if count > maxCount {
		res = []int{root.Val}
		maxCount = count
	}

	if root.Right != nil {
		findModeTraversal(root.Right)
	}
}
