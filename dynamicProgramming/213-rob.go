package dynamicProgramming

func rob2(nums []int) int {
	if len(nums) == 1 {
		return rob(nums)
	}
	robLeft := rob(nums[:len(nums)-1])
	robRight := rob(nums[1:])
	return Max(robLeft, robRight)
}
