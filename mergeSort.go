package main

func MergeSort(nums []int) {
	Sort(nums, 0, len(nums)-1)
}

func Sort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi - lo)/2
	Sort(nums, lo, mid)
	Sort(nums, mid+1, hi)
	Merge(nums, lo, mid, hi)
}

func Merge(nums []int, lo, mid, hi int) {
	i, j := lo, mid+1
	aux := make([]int, len(nums))
	k := lo
	for k <= hi {
		if i > mid {
			aux[k] = nums[j]
			k++
			j++
		} else if j > hi {
			aux[k] = nums[i]
			k++
			i++
		} else if Less(nums, i, j) {
			aux[k] = nums[i]
			k++
			i++
		} else {
			aux[k] = nums[j]
			k++
			j++
		}
	}
	for i=lo; i<=hi; i++ {
		nums[i] = aux[i]
	}
}