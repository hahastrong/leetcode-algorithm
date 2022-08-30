package main

func ShellSort(nums []int) {
	h := 1
	n := len(nums)
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i<n; i++ {
			for j := i; j >= h && Less(nums, j, j-h); j-- {
				Exch(nums, j, j-h)
			}
		}
		h /= 3
	}
}
