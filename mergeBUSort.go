package main

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func MergeBUSort(nums []int) {
	n := len(nums)
	for sz := 1; sz < n; sz = sz + sz {
		for i:=0; i<n-sz; i += sz+sz {
			Merge(nums, i, i+sz-1, min(i+sz+sz-1, n-1))
		}
	}
}
