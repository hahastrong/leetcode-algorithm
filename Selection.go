package main

func SelectionSort(nums []int) {
	n := len(nums)
	for i:=0; i<n-1; i++ {
		min := i
		for j:=i+1; j<n; j++ {
			if Less(nums, j, min) {
				min = j
			}
		}
		Exch(nums, min, i)
	}
}
