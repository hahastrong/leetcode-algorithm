package main

func Less(nums []int, i, j int) bool {
	return nums[i] < nums[j]
}

func Exch(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}