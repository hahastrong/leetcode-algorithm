package dynamicProgramming

func canPartition(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}

	sum := 0
	for i:=0; i<n; i++ {
		sum += nums[i]
	}

	if sum % 2 == 1 {
		return false
	}
	m := sum/2
	dp := make([]int, m+1)

	for i:=0; i<n; i++ {
		for j:=m; j>0; j-- {
			if j < nums[i] {
				break
			}
			dp[j] = Max(dp[j], dp[j-nums[i]] + nums[i])
		}
	}
	return dp[m] == m
}

