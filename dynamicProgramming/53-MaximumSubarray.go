package dynamicProgramming

// DP 需要注意数据的的初始化工作
// 如果遇到和小鱼0， 则需要将当前数置为 大于等于 0 的数
func maxSubArray(nums []int) int {

	dp := make([]int, len(nums))
	max := nums[0]
	dp[0] = nums[0]

	for i:=1; i<len(nums); i++ {
		if dp[i-1] + nums[i] > 0 {
			dp[i] = dp[i-1] + nums[i]
			max = Max(dp[i], max)
		} else {
			dp[i] = Max(0, nums[i])
			max = Max(nums[i], max)
		}
	}
	return max
}