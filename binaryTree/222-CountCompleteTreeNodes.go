package binaryTree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	idx := 0
	for len(q) > 0 {
		qLen := len(q)
		if qLen < 1 << idx {
			return (1 << idx) - 1 + qLen
		}
		for i:=0; i<qLen; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[qLen:]
		idx++
	}
	return (1 << idx) - 1
}


