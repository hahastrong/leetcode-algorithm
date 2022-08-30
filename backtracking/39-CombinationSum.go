package backtracking

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	backtrackingSum(candidates, target, 0, []int{}, &res)
	return res
}

func backtrackingSum(candidates []int, target int, idx int, path []int, res *[][]int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}

	if target <= 0 {
		return
	}

	for i:=idx; i<len(candidates); i++ {
		path = append(path, candidates[i])
		backtrackingSum(candidates, target-candidates[i], i, path, res)
		path = path[:len(path)-1]
	}
}
