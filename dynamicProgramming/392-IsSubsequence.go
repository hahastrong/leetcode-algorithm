package dynamicProgramming

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	dp := make([][]int, len(t)+1)
	for i:=0; i<=len(t); i++ {
		dp[i] = make([]int, len(s)+1)
	}

	for i:=1; i<=len(t); i++ {
		for j:=1; j<=len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i][j-1], dp[i-1][j])
			}
		}
		if dp[i][len(s)] == len(s) {
			return true
		}
	}
	return false
}
