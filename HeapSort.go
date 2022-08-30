package main

func HeapSort(nums []int) []int {
	num := append([]int{0}, nums...)
	n := len(num)-1
	for i:=n/2; i>=1; i-- {
		sink(num, i, n)
	}
	for n > 1 {
		Exch(num, 1, n)
		n -= 1
		sink(num, 1, n)
	}

	return num[1:]
}

func sink(nums []int, k, N int) {
	for 2*k <= N {
		j := 2 * k
		if j < N && Less(nums, j, j+1) {
			j++
		}
		if Less(nums, j, k) {
			break
		}
		Exch(nums, k, j)
		k = j
	}
}


