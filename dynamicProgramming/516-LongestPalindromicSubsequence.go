package dynamicProgramming

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s)+1)
	for i:=0; i<=len(s); i++ {
		dp[i] = make([]int, len(s)+1)
		dp[i][i] = 1
	}

	for i:=len(s)-1; i>0; i-- {
		for j:=i+1; j<=len(s); j++ {
			if s[i-1] == s[j-1] {
				if j-i == 1 {
					dp[i][j] = 2
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[1][len(s)]
}
