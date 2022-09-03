package greedyAlgorithm

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	maxIdx := 0
	res := 0
	for maxIdx < len(nums) - 1 {
		max := 0
		tmpIdx := 0
		res++
		for i:=1; i<=nums[maxIdx] && i+maxIdx < len(nums); i++ {
			if maxIdx + i >= len(nums)-1 {
				return res
			}
			if  maxIdx+i+nums[maxIdx+i] >= max {
				max = maxIdx+i+nums[maxIdx+i]
				tmpIdx = maxIdx + i
			}
		}
		maxIdx = tmpIdx
	}
	return res
}
