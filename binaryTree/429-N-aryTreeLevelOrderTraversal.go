package binaryTree

//func levelOrderN(root *Node) [][]int {
//	var res [][]int
//	if root == nil {
//		return res
//	}
//
//	q := make([]*Node, 0)
//	q = append(q, root)
//	for len(q) > 0 {
//		qLen := len(q)
//		levelRes := make([]int, qLen)
//		for i:=0; i < qLen; i++ {
//			tmpNode := q[i]
//			levelRes[i] = tmpNode.Val
//			for child:=0; child < len(tmpNode.Children); child++ {
//				q = append(q, tmpNode.Children[child])
//			}
//		}
//		res = append(res, levelRes)
//		q = q[qLen:]
//	}
//	return res
//}
