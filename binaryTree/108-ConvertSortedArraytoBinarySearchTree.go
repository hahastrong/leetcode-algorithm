package binaryTree

func sortedArrayToBST(nums []int) *TreeNode {
	return constructBST(nums, 0, len(nums)-1)
}

func constructBST(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	mid := low + (high - low) / 2
	node := TreeNode{
		Val: nums[mid],
	}

	node.Left = constructBST(nums, low, mid-1)
	node.Right = constructBST(nums, mid+1, high)
	return &node
}


