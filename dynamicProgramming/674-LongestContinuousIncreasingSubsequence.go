package dynamicProgramming

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 1
	dp[0] = 1
	for i:=1; i<len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if dp[i] > max {
				max = dp[i]
			}
		} else {
			dp[i] = 1
		}
	}

	return max
}
