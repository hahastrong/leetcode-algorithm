package dynamicProgramming

// 凑成目标数字的方法总数， 全排列
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i:=0; i<=target; i++ {
		for j:=0; j<len(nums); j++ {
			if nums[j] > i {
				continue
			}
			dp[i] += dp[i-nums[j]]
		}
	}
	return dp[target]
}
