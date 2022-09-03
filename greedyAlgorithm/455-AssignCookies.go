package greedyAlgorithm

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	count, gIdx, sIdx := 0, len(g)-1, len(s)-1
	for gIdx >=0 && sIdx >= 0 {
		if g[gIdx] > s[sIdx] {
			gIdx--
		} else {
			gIdx--
			sIdx--
			count++
		}
	}

	return count
}
