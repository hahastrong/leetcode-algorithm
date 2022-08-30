package main

import (
	"testing"
)

func findKthLargest(nums []int, k int) int {
	idx := 0
	for i:=1; i<len(nums); i++ {
		if idx < k-1 {
			idx++
			for j := i; j > 0 ;j-- {
				if nums[i] > nums[i-1] {
					nums[i], nums[i-1] = nums[i-1], nums[i]
				}
			}

		} else if nums[i] > nums[idx] {
			nums[i], nums[idx] = nums[idx], nums[i]
			for j:=idx; j>0; j--{
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums[idx]
}

type testData struct {
	nums []int
	k int
	res int
}

func TestKthElement(t *testing.T) {
	nums := []testData{
		{
			nums: []int{3,2,3,1,2,4,5,5,6},
			k: 4,
			res: 4,
		},
		{
			nums: []int{3,2,1,5,6,4},
			k: 2,
			res: 5,
		},
		{
			nums: []int{-1,2,0},
			k: 2,
			res: 0,
		},
	}

	for idx, data := range nums {
		res := findKthLargest(data.nums, data.k)
		if res != data.res {
			t.Logf("failed to test: [%v] actual: %v, expect: %v\n", idx, res, data.res)
		}
	}
}
