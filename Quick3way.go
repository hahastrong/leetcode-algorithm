package main

func QuickSort3way(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func less(i, j int) int {
	if i > j {
		return 1
	} else if i == j {
		return 0
	}
	return -1
}

func sort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	lt, gt := lo, hi
	i := lo+1
	pivot := nums[lo]

	for i <= gt {
		cmp := less(nums[i], pivot)
		if cmp < 0 {
			Exch(nums, i, lt)
			i++
			lt++
		} else if cmp > 0 {
			Exch(nums, i, gt)
			gt--
		} else {
			i++
		}
	}
	sort(nums, lo, lt-1)
	sort(nums, gt+1, hi)
}
