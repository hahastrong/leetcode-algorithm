package dynamicProgramming

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i:=2; i<=n; i++ {
		for j:=1; j<i; j++ {
			dp[i] = Max(dp[i], (i-j) * j, (i-j) * dp[j], dp[i-j]*j, dp[i-j]*dp[j])
		}
	}
	return dp[n]
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


