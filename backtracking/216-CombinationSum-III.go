package backtracking

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	backtrackingCom3(k, n, 1, []int{}, &res)
	return res
}

func backtrackingCom3(k, remainSum, idx int, path []int, res *[][]int) {
	if k == 0  {
		if remainSum == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			*res = append(*res, temp)
		}
		return
	}

	if remainSum <= 0 {
		return
	}

	for i:=idx; i<10; i++ {
		path = append(path, i)
		backtrackingCom3(k-1, remainSum-i, i+1, path, res)
		path = path[:len(path)-1]
	}
}
