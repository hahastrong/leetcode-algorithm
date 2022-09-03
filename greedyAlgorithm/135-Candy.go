package greedyAlgorithm

func candy(ratings []int) int {
	nums := make([]int, len(ratings))
	nums[0] = 1
	for i:=1; i<len(nums); i++ {
		if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		} else {
			nums[i] = 1
		}
	}
	sum := nums[len(nums)-1]
	for i:=len(nums)-2; i>=0; i-- {

		if ratings[i+1] < ratings[i] && nums[i+1] >= nums[i] {
			nums[i] = nums[i+1] + 1
		}
		sum += nums[i]
	}

	return sum
}
