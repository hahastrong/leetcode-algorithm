package dynamicProgramming

func findTargetSumWays(nums []int, target int) int {

	sum := 0
	for i:=0; i<len(nums); i++ {
		sum += nums[i]
	}
	if (sum + target)%2 == 1{
		return 0
	}
	m := (sum + target)/2
	dp := make([]int, m+1)
	dp[0] = 1
	for i:=0; i<len(nums); i++ {
		for j:=m; j>0; j-- {
			if j < nums[i] {
				break
			}

			dp[j] += dp[j-nums[i]]

		}
	}
	return dp[m]
}
