package dynamicProgramming

type TreeNode struct {
	 Val int
     Left *TreeNode
     Right *TreeNode
}

// 打家劫舍 III 通过遍历整个二叉树寻找上一个节点与这一个节点可获得最大利益的组合
func rob3(root *TreeNode) int {
	res := robOneNode(root)
	return Max(res[0], res[1])
}

func robOneNode(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := robOneNode(root.Left)
	right := robOneNode(root.Right)

	res := make([]int, 2)
	res[0] = Max(left[0]+right[0], left[1]+right[1])
	res[1] = left[0]+right[0]+root.Val
	return res
}

