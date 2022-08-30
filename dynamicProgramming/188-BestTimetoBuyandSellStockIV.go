package dynamicProgramming

func maxProfitIV(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i:=0; i<len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}

	for i:=1; i<2*k; i++ {
		dp[0][i] = -prices[0]
	}

	for i:=1; i<len(prices); i++ {
		for j:=1; j<=2*k; j+=2 {
			dp[i][j] = Max(dp[i-1][j], dp[i-1][j-1] - prices[i])
			dp[i][j+1] = Max(dp[i-1][j+1], dp[i-1][j] + prices[i])
		}
	}

	return Max(dp[len(prices)-1]...)
}
