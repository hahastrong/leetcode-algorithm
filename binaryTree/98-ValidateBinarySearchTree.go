package binaryTree

func isValidBST(root *TreeNode) bool {

	_, flag := IsValidBSTHelp(root, nil)
	return flag
}

func IsValidBSTHelp(cur *TreeNode, pre *TreeNode) (*TreeNode, bool) {
	if cur == nil {
		return pre, true
	}

	var flag bool
	if cur.Left != nil {
		if pre, flag = IsValidBSTHelp(cur.Left, pre); !flag {
			return pre, false
		}
	}

	if pre != nil && cur.Val <= pre.Val {
		return pre, false
	}

	pre = cur

	if cur.Right != nil {
		if pre, flag = IsValidBSTHelp(cur.Right, pre); !flag {
			return pre, false
		}
	}

	return pre, true
}
