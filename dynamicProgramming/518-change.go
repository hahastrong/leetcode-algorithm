package dynamicProgramming

// 为什么内层遍历从前往后，因为每个元素可以重复
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1


	for j:=0; j<len(coins); j++ {
		for i:=1; i<=amount; i++ {
			if coins[j] > i {
				continue
			}
			dp[i] += dp[i-coins[j]]
		}
	}
	return dp[amount]
}