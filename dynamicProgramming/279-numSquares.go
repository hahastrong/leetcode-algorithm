package dynamicProgramming

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i:=1; i<=n; i++ {
		dp[i] = math.MaxInt32
	}

	for i:=1; i<=n; i++ {
		for j:=1; j*j<=i; j++ {
			dp[i] = Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

