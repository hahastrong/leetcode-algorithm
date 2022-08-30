package dynamicProgramming

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)

	dp[0] = true

	for i:=0; i<len(s); i++ {
		for j:=0; j<len(wordDict); j++ {
			if len(wordDict[j]) <= i+1 && isValid(s, i, wordDict[j]) {
				if dp[i+1-len(wordDict[j])] {
					dp[i+1] = true
				}
			}
		}
	}

	return dp[len(s)]
}

func isValid(s string, last int, target string) bool {
	idx := last - len(target) + 1
	for i:=0; i<len(target); i++ {
		if s[idx + i] != target[i] {
			return false
		}
	}
	return true
}
