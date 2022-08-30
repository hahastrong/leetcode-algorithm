package backtracking

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	backtrackingSum2(candidates, target, 0, []int{}, &res)
	return res
}

func backtrackingSum2(candidates []int, target int, idx int, path []int, res *[][]int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}

	if target <= 0 {
		return
	}

	flag := make([]bool, len(candidates))
	for i:=idx; i<len(candidates); i++ {
		flag[i] = true
		if i > 0 && candidates[i-1] == candidates[i] && flag[i-1] {
			continue
		}
		path = append(path, candidates[i])
		backtrackingSum2(candidates, target-candidates[i], i+1, path, res)
		path = path[:len(path)-1]
	}
}