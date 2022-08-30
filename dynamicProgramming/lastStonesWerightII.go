package dynamicProgramming

func lastStoneWeightII(stones []int) int {
	n := len(stones)
	if n <= 1 {
		return stones[0]
	}

	sum := 0
	for i:= 0; i<n; i++ {
		sum += stones[i]
	}

	m := sum/2

	dp := make([]int, m+1)
	for i:=0; i<n; i++ {
		for j:=m; j>0; j-- {
			if j < stones[i] {
				break
			}
			dp[j] = Max(dp[j], dp[j-stones[i]] + stones[i])
		}
	}
	return sum - dp[m] - dp[m]
}
