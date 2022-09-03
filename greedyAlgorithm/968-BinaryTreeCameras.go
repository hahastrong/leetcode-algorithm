package greedyAlgorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	res := 0
	v := backtracking(root, &res)
	if v == 0 {
		res++
	}
	return res
}

func backtracking(root *TreeNode, res *int) int {
	if root == nil {
		return 2
	}

	left := backtracking(root.Left, res)
	right := backtracking(root.Right, res)

	if left == 2 && right == 2 {
		return 0
	}

	if left == 0 || right == 0 {
		*res++
		return 1
	}

	if left == 1 || right == 1 {
		return 2
	}

	return 2
}
