package binaryTree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	return help(root, p, q)
}

func help(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}

	if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val {
		return lowestCommonAncestor(root.Right,p, q)
	}
	return nil
}
