package greedyAlgorithm

import "math"

func canCompleteCircuit(gas []int, cost []int) int {

	sum := 0
	for i:=0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
	}
	if sum < 0 {
		return -1
	}

	max := math.MinInt32
	idx := len(gas)
	sum = 0
	for i:=len(gas)-1; i>=0; i-- {
		sum += gas[i] - cost[i]
		if sum > max {
			max = sum
			idx = i
		}
	}

	return idx
}
