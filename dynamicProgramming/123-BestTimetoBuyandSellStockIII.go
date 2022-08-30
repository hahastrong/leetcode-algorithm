package dynamicProgramming

func maxProfitIII(prices []int) int {
	dp := make([][]int, len(prices))
	for i:=0; i<len(prices); i++ {
		dp[i] = make([]int, 5)
	}

	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]

	for i:=1; i<len(prices); i++ {
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0] - prices[i])
		dp[i][2] = Max(dp[i-1][2], dp[i-1][1] + prices[i])
		dp[i][3] = Max(dp[i-1][3], dp[i-1][2] - prices[i])
		dp[i][4] = Max(dp[i-1][4], dp[i-1][3] + prices[i])
	}

	return Max(dp[len(prices)-1]...)
}
