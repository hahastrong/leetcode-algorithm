package dynamicProgramming

func maxProfitVI(prices []int, fee int) int {
	if len(prices) <= 1 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i:=0; i<len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][1] = -prices[0] - fee

	for i:=1; i<len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
	}
	return Max(dp[len(prices)-1]...)
}
