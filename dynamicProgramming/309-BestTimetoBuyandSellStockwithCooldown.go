package dynamicProgramming

func maxProfitV(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i:=0; i<len(prices); i++ {
		dp[i] = make([]int, 4)
	}

	dp[0][1] = -prices[0]

	for i:=1; i<len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][3])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0] - prices[i], dp[i-1][3] - prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return Max(dp[len(prices)-1]...)
}