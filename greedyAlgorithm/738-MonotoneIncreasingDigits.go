package greedyAlgorithm

func monotoneIncreasingDigits(n int) int {
	last, all, count, num, cur := n%10, 0, 0, n/10, 0
	for num > 0 {
		cur = num %10
		all++
		if cur > last {
			cur--
			count = all
		}
		last = cur
		num /= 10
	}
	for i:=0; i<count; i++ {
		n /= 10
	}
	if count > 0 {
		n--
	}
	for count>0 {
		n = n * 10 + 9
		count--
	}
	return n

}
