package dynamicProgramming

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i:=0; i<=len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	for i:=1; i<=len(word1); i++ {
		for j:=1; j<=len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return len(word1) + len(word2) - 2 * dp[len(word1)][len(word2)]
}
