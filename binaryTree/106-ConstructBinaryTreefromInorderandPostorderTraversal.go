package binaryTree

func buildTree(inorder []int, postorder []int) *TreeNode {
	return constructIP(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func constructIP(inorder []int, istart, iend int,
					postorder []int, pstart, pend int) *TreeNode {
	if istart == iend {
		return &TreeNode{
			Val: inorder[istart],
		}
	}

	idx := 0
	for i:=istart; i<= iend; i++ {
		if inorder[i] == postorder[pend] {
			idx = i
			break
		}
	}

	node := TreeNode{
		Val: inorder[idx],
	}

	if idx > istart {
		node.Left = constructIP(inorder, istart, idx-1, postorder, pstart, pstart + (idx - istart) - 1)
	}
	if idx < iend {
		node.Right = constructIP(inorder, idx+1, iend, postorder, pstart + (idx - istart), pend-1)
	}
	return &node
}
