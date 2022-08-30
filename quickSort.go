package main

func QuickSort(nums []int) {
	Quick(nums, 0, len(nums)-1)

}

func Quick(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivot := partition(nums, lo, hi)
	Quick(nums, lo, pivot-1)
	Quick(nums, pivot+1, hi)
}

func partition(nums []int, lo, hi int) int {
	pivot := lo
	i, j := lo+1, hi
	for i<=j {
		for Less(nums, i, pivot) && i <= j {
			i++
		}

		for Less(nums, pivot, j) && j >= i{
			j--
		}

		if i >= j {
			break
		}

		Exch(nums, i, j)
	}
	Exch(nums, pivot, j)
	return j
}
