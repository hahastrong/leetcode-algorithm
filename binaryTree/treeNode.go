package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//type Node struct {
//	Val int
//	Children []*Node
//}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Max(nums... int) int {
	max := nums[0]
	for i:=1; i<len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func Min(nums... int) int {
	min := nums[0]
	for i:=1; i<len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
