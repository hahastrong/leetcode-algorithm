package greedyAlgorithm

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxIdx := 0
	for maxIdx < len(nums) - 1 {
		if nums[maxIdx] == 0 {
			return false
		}
		max := 0
		tmpIdx := 0
		for i:=1; i<=nums[maxIdx] && i+maxIdx < len(nums); i++ {
			if  maxIdx+i+nums[maxIdx+i] >= max {
				max = maxIdx+i+nums[maxIdx+i]
				if max >= len(nums) - 1 {
					return true
				}
				tmpIdx = maxIdx + i
			}
		}
		maxIdx = tmpIdx
	}
	return false
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
