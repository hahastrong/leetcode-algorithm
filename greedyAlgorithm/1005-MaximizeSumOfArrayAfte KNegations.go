package greedyAlgorithm

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
	})

	sum := 0
	for i:=len(nums)-1; i>=1; i-- {
		if k > 0 && nums[i] < 0 {
			sum -= nums[i]
			k--
		} else {
			sum += nums[i]
		}
	}

	if k%2 == 1 {
		sum -= nums[0]
	} else {
		sum += nums[0]
	}

	return sum
}
