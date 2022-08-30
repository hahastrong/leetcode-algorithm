package main

func InsertSort(nums []int) {
	n := len(nums)
	for i:=1; i<n; i++ {
		for j:=i; j>0 && Less(nums, j, j-1); j-- {
			Exch(nums, j, j-1)
		}
	}
}
