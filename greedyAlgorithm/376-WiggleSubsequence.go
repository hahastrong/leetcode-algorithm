package greedyAlgorithm

func wiggleMaxLength(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return 1
		}
		return 2
	}

	count := 2

	last := nums[1] - nums[0]
	if last == 0 {
		count = 1
	}

	for i:=2; i<len(nums); i++ {
		cur := nums[i] - nums[i-1]
		if (last > 0 && cur < 0 ) || (last < 0 && cur > 0) {
			count++
			last = cur
		}
	}
	return count
}