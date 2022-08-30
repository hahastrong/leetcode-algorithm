package dynamicProgramming

// 打家劫舍
func rob(nums []int) int {
	dp := make([][]int, 2)
	for i:=0; i<2; i++ {
		dp[i] = make([]int, len(nums)+1)
	}

	for i:=1; i<=len(nums); i++ {
		dp[1][i] = Max(dp[0][i-1]+nums[i-1], dp[1][i-1])
		dp[0][i] = Max(dp[0][i-1], dp[1][i-1])
	}
	return Max(dp[0][len(nums)], dp[1][len(nums)])
}
