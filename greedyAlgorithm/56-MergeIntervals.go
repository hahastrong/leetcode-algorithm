package greedyAlgorithm

import "sort"

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])
	for i:=1; i<len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = Max(intervals[i][1], res[len(res)-1][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}
