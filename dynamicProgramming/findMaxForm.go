package dynamicProgramming

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i:=0; i<=m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i:=0; i<len(strs); i++ {
		nums0 := 0
		for j:=0; j<len(strs[i]); j++ {
			if strs[i][j] == '0' {
				nums0++
			}
		}
		nums1 := len(strs[i]) - nums0

		for j:=m; j>=nums0; j-- {
			for k := n; k>=nums1; k-- {
				dp[j][k] = Max(dp[j][k], dp[j-nums0][k-nums1]+1)
			}
		}
	}
	return dp[m][n]
}
