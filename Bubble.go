package main

func BubbleSort(nums []int) {
	n := len(nums)
	for i:=0; i<n-1; i++ {
		for j:=n-1; j>i; j-- {
			if Less(nums, j, j-1) {
				Exch(nums, j, j-1)
			}
		}
	}
}
