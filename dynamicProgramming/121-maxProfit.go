package dynamicProgramming

import "math"

func maxProfit(prices []int) int {
	min, result := math.MaxInt32, 0

	for i:=0; i<len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		result = Max(result, prices[i] - min)
	}
	return result
}