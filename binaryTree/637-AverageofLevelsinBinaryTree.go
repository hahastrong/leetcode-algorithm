package binaryTree

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64

	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		qLen := len(q)
		var tmpValue float64
		for i:=0; i < qLen; i++ {
			tmpNode := q[i]
			tmpValue += float64(tmpNode.Val)
			if tmpNode.Left != nil {
				q = append(q, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				q = append(q, tmpNode.Right)
			}
		}
		res = append(res, tmpValue/float64(qLen))
		q = q[qLen:]
	}
	return res
}
