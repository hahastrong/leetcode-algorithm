package binaryTree

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) > 0 {
		qLen := len(q)
		for i:=0; i < qLen; i++ {
			tmpNode := q[i]
			if i < qLen - 1 {
				tmpNode.Next = q[i+1]
			} else {
				tmpNode.Next = nil
			}
			if tmpNode.Left != nil {
				q = append(q, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				q = append(q, tmpNode.Right)
			}
		}
		q = q[qLen:]
	}
	return root
}
