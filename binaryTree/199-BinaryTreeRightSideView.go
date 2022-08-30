package binaryTree


func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		tmpValue := 0
		qLen := len(q)
		for i:=0; i < qLen; i++ {
			tmpNode := q[i]
			tmpValue = tmpNode.Val
			if tmpNode.Left != nil {
				q = append(q, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				q = append(q, tmpNode.Right)
			}
		}
		res = append(res, tmpValue)
		q = q[qLen:]
	}
	return res
}