package binaryTree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaxTreeHelp(nums, 0, len(nums)-1)
}

func constructMaxTreeHelp(nums []int, start, end int) *TreeNode {
	if start == end {
		return &TreeNode{Val: nums[start]}
	}

	max := start
	for i:=start+1; i <= end; i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}

	node := TreeNode{
		Val: nums[max],
	}

	if max > start {
		node.Left = constructMaxTreeHelp(nums, start, max-1)
	}
	if max < end {
		node.Right = constructMaxTreeHelp(nums, max+1, end)
	}
	return &node
}
