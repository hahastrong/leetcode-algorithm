package dynamicProgramming

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 0
	dp[0] = 1
	for i:=1; i<len(nums); i++ {
		dp[i] = 1
		for j:=0; j<i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j] + 1)
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}
	return max
}
