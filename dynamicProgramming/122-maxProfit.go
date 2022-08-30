package dynamicProgramming

// 买卖股票的最佳时机，不限制交易次数
func maxProfitII(prices []int) int {
	dp := make([][]int, len(prices))
	for i:=0; i<len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][1] = -prices[0]

	for i:=1; i<len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0] - prices[i])
	}

	return Max(dp[len(prices)-1]...)
}