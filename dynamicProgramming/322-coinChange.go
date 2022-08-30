package dynamicProgramming

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0]=0
	for i:=1; i<= amount; i++ {
		dp[i] = 10000+1
	}
	for i:=1; i<=amount; i++ {
		for j:=0; j<len(coins); j++ {
			if coins[j] > i {
				continue
			}
			dp[i] = Min(dp[i], dp[i-coins[j]]+1)
		}
	}
	if dp[amount] <= amount {
		return dp[amount]
	}
	return -1
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
