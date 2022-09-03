package greedyAlgorithm

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})


	end := intervals[0][1]

	count := 0
	for i:=1; i<len(intervals); i++ {
		if intervals[i][0] >= end {

			end = intervals[i][1]
		} else {
			count++
		}
	}
	return count

}
