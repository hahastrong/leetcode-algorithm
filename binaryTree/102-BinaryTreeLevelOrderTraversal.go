package binaryTree

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		qLen := len(q)
		levelRes := make([]int, qLen)
		for i:=0; i<qLen; i++ {
			tmpNode := q[i]
			levelRes[i] = tmpNode.Val
			if tmpNode.Left != nil {
				q = append(q, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				q = append(q, tmpNode.Right)
			}
		}
		q = q[qLen:]
		res = append(res, levelRes)
	}

	return res
}


