package greedyAlgorithm

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	start := points[0][0]
	end := points[0][1]
	count := 0

	for i:=1; i<len(points); i++ {
		if points[i][0] <= end {
			start = Max(start, points[i][0])
			end = Min(end, points[i][1])
		} else {
			count++
			start = points[i][0]
			end = points[i][1]
		}
	}
	count++
	return count
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}


