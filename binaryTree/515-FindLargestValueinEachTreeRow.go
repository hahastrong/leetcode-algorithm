package binaryTree

import "math"

func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q =  append(q, root)

	for len(q) > 0 {
		qLen := len(q)
		max := math.MinInt32
		for i:=0; i < qLen; i++ {
			tmpNode := q[i]
			if tmpNode.Val > max {
				max = tmpNode.Val
			}
			if tmpNode.Left != nil {
				q = append(q, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				q = append(q, tmpNode.Right)
			}
		}
		res = append(res, max)
		q = q[qLen:]
	}
	return res
}
